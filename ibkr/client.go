package ibkr

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/go-playground/validator/v10"
)

const ProdBaseUrl = "https://api.ibkr.com"
const GatewayBaseUrl = "https://localhost:5000"
const DefaultUserAgent = "schmidthole/ibkr-webapi-go/1"

const (
	methodGet    = "GET"
	methodPost   = "POST"
	methodDelete = "DELETE"
)

var IbkrGlobalRateLimit = 50

type clientResponse struct {
	statusCode int
	bytes      []byte
}

type IbkrWebClient struct {
	VerboseLogging bool
	BaseUrl        string
	client         *http.Client
	oauth          OAuthContext
	validator      *validator.Validate
}

func NewIbkrWebClient(baseUrl string, authContext OAuthContext) *IbkrWebClient {
	client := http.Client{Timeout: 15 * time.Second}

	// skip cert auth check if talking to gateway
	if authContext == nil {
		skipVerifyTransport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
		client.Transport = skipVerifyTransport
	}

	return &IbkrWebClient{
		VerboseLogging: false,
		BaseUrl:        baseUrl,
		client:         &client,
		oauth:          authContext,
		validator:      validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (c *IbkrWebClient) DoRequest(
	method string,
	path string,
	queryParams map[string]string,
	body interface{},
) (*clientResponse, error) {
	base, err := url.Parse(c.BaseUrl)
	if err != nil {
		return nil, err
	}
	requestUrl := base.ResolveReference(&url.URL{Path: fmt.Sprintf("/v1/api%s", path)})

	if queryParams != nil {
		params := url.Values{}
		for key, value := range queryParams {
			params.Add(key, value)
		}

		requestUrl.RawQuery = params.Encode()
	}

	var requestBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		requestBody = bytes.NewBuffer(jsonBody)
	}

	request, err := http.NewRequest(method, requestUrl.String(), requestBody)
	if err != nil {
		return nil, err
	}

	if c.oauth != nil {
		authHeader, err := c.oauth.GetOAuthHeader(method, requestUrl.String())
		if err != nil {
			return nil, err
		}
		request.Header.Set("Authorization", authHeader)
	}

	request.Header.Set("User-Agent", DefaultUserAgent)

	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	logRequest(request, c.VerboseLogging)

	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	logResponse(response, c.VerboseLogging)

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &clientResponse{statusCode: response.StatusCode, bytes: bodyBytes}, nil
}

func (c *IbkrWebClient) Get(path string, queryParams map[string]string) (*clientResponse, error) {
	return c.DoRequest(methodGet, path, queryParams, nil)
}

func (c *IbkrWebClient) Post(path string, queryParams map[string]string, body interface{}) (*clientResponse, error) {
	return c.DoRequest(methodPost, path, queryParams, body)
}

func (c *IbkrWebClient) Delete(path string, queryParams map[string]string) (*clientResponse, error) {
	return c.DoRequest(methodDelete, path, queryParams, nil)
}

func (c *IbkrWebClient) Authenticate() error {
	return c.oauth.GenerateLiveSessionToken(c.client, c.BaseUrl)
}

func (c *IbkrWebClient) ParseJsonResponse(response *clientResponse, v interface{}) error {
	err := json.Unmarshal(response.bytes, v)
	if err != nil {
		return err
	}

	value := reflect.ValueOf(v)
	if value.Kind() == reflect.Ptr && value.Elem().Kind() == reflect.Slice {
		sliceValue := value.Elem()

		for i := 0; i < sliceValue.Len(); i++ {
			elem := sliceValue.Index(i).Addr().Interface()
			err = c.validator.Struct(elem)
			if err != nil {
				logValidationErrors(err)
				return err
			}
		}
	} else {
		err = c.validator.Struct(v)
		if err != nil {
			logValidationErrors(err)
			return err
		}
	}

	return nil
}

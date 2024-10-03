package ibkr

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testBody struct {
	Message string `json:"message"`
	Num     int    `json:"num"`
}

type testResponse struct {
	Message string `json:"message"`
}

func TestNewIbkrWebClient(t *testing.T) {
	client := NewIbkrWebClient("mockurl", &MockOAuthContext{})

	assert.Equal(t, "mockurl", client.BaseUrl)
}

func TestIbkrWebClient_Get(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)

		assert.NotEmpty(t, r.Header.Get("User-Agent"))
		assert.NotEmpty(t, r.Header.Get("Authorization"))

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer mockServer.Close()

	client := NewIbkrWebClient(mockServer.URL, &MockOAuthContext{})

	rsp, err := client.Get("/test-endpoint", nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rsp.statusCode)

	var rspStruct testResponse
	err = client.ParseJsonResponse(rsp, &rspStruct)
	assert.NoError(t, err)

	assert.Equal(t, "success", rspStruct.Message)
}

func TestIbkrWebClient_GetBadRsp(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)

		assert.NotEmpty(t, r.Header.Get("User-Agent"))
		assert.NotEmpty(t, r.Header.Get("Authorization"))

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "failed"}`))
	}))
	defer mockServer.Close()

	client := NewIbkrWebClient(mockServer.URL, &MockOAuthContext{})

	rsp, err := client.Get("/test-endpoint", nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rsp.statusCode)

	var rspStruct testResponse
	err = client.ParseJsonResponse(rsp, &rspStruct)
	assert.NoError(t, err)

	assert.Equal(t, "failed", rspStruct.Message)
}

func TestIbkrWebClient_GetWithQueryParams(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)

		assert.NotEmpty(t, r.Header.Get("User-Agent"))
		assert.NotEmpty(t, r.Header.Get("Authorization"))

		params := r.URL.Query()
		assert.Equal(t, "yes", params.Get("test"))

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer mockServer.Close()

	client := NewIbkrWebClient(mockServer.URL, &MockOAuthContext{})

	params := map[string]string{
		"test": "yes",
	}

	rsp, err := client.Get("/test-endpoint", params)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rsp.statusCode)

	var rspStruct testResponse
	err = client.ParseJsonResponse(rsp, &rspStruct)
	assert.NoError(t, err)

	assert.Equal(t, "success", rspStruct.Message)
}

func TestIbkrWebClient_Post(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)

		assert.NotEmpty(t, r.Header.Get("User-Agent"))
		assert.NotEmpty(t, r.Header.Get("Authorization"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		bodyBytes, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		var reqBody testBody
		err = json.Unmarshal(bodyBytes, &reqBody)
		assert.NoError(t, err)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer mockServer.Close()

	client := NewIbkrWebClient(mockServer.URL, &MockOAuthContext{})

	rsp, err := client.Post(
		"/test-endpoint",
		nil,
		testBody{Message: "hello", Num: 1},
	)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rsp.statusCode)

	var rspStruct testResponse
	err = client.ParseJsonResponse(rsp, &rspStruct)
	assert.NoError(t, err)

	assert.Equal(t, "success", rspStruct.Message)
}

func TestIbkrWebClient_Delete(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)

		assert.NotEmpty(t, r.Header.Get("User-Agent"))
		assert.NotEmpty(t, r.Header.Get("Authorization"))

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer mockServer.Close()

	client := NewIbkrWebClient(mockServer.URL, &MockOAuthContext{})

	rsp, err := client.Delete("/test-endpoint", nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rsp.statusCode)

	var rspStruct testResponse
	err = client.ParseJsonResponse(rsp, &rspStruct)
	assert.NoError(t, err)

	assert.Equal(t, "success", rspStruct.Message)
}

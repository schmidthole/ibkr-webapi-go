package ibkr

import (
	"fmt"
	"net/http"
)

type AuthStatus struct {
	Authenticated bool       `json:"authenticated" validation:"required"`
	Competing     bool       `json:"competing" validation:"required"`
	Connected     bool       `json:"connected" validation:"required"`
	Message       string     `json:"message" validation:"required"`
	MAC           string     `json:"MAC" validation:"required"`
	ServerInfo    ServerInfo `json:"serverInfo" validation:"required"`
}

type ServerInfo struct {
	ServerName    string `json:"serverName" validation:"required"`
	ServerVersion string `json:"serverVersion" validation:"required"`
}

/******************************************************************************
* logout
******************************************************************************/

type LogoutResponse struct {
	Status bool `json:"status" validation:"required"`
}

func (c *IbkrWebClient) Logout() error {
	response, err := c.Post("/logout", nil, nil)
	if err != nil {
		return err
	}

	if response.statusCode != http.StatusOK {
		return fmt.Errorf("bad logout statusCode: %v", response.statusCode)
	}

	var responseStruct LogoutResponse
	err = c.ParseJsonResponse(response, &responseStruct)
	if err != nil {
		return err
	}

	if !responseStruct.Status {
		return fmt.Errorf("ibkr logout response bad status")
	}

	return nil
}

/******************************************************************************
* brokerage session init
******************************************************************************/

type InitializeBrokerageSessionRequest struct {
	Publish bool `json:"publish" validation:"required"`
	Compete bool `json:"compete" validation:"required"`
}

func (c *IbkrWebClient) InitializeBrokerSession() (*AuthStatus, error) {
	requestBody := InitializeBrokerageSessionRequest{
		Publish: true,
		Compete: true,
	}

	response, err := c.Post("/iserver/auth/ssodh/init", nil, requestBody)
	if err != nil {
		return nil, err
	}

	if response.statusCode != http.StatusOK {
		return nil, fmt.Errorf("bad logout statusCode: %v", response.statusCode)
	}

	var responseStruct AuthStatus
	err = c.ParseJsonResponse(response, &responseStruct)
	if err != nil {
		return nil, err
	}

	return &responseStruct, nil
}

/******************************************************************************
* auth status
******************************************************************************/

func (c *IbkrWebClient) AuthStatus() (*AuthStatus, error) {
	response, err := c.Post("/iserver/auth/status", nil, nil)
	if err != nil {
		return nil, err
	}

	if response.statusCode != http.StatusOK {
		return nil, fmt.Errorf("bad logout statusCode: %v", response.statusCode)
	}

	var responseStruct AuthStatus
	err = c.ParseJsonResponse(response, &responseStruct)
	if err != nil {
		return nil, err
	}

	return &responseStruct, nil
}

/******************************************************************************
* server ping
******************************************************************************/

type TickleResponse struct {
	Session    string         `json:"session" validation:"required"`
	SSOExpires int32          `json:"ssoExpires" validation:"required"`
	Collission bool           `json:"collission" validation:"required"`
	UserID     int32          `json:"userId" validation:"required"`
	HMDS       HMDSDetails    `json:"hmds" validation:"required"`
	IServer    IServerDetails `json:"iserver" validation:"required"`
}

type HMDSDetails struct {
	Error string `json:"error"`
}

type IServerDetails struct {
	AuthStatus AuthStatus `json:"authStatus"`
}

func (c *IbkrWebClient) Tickle() (*TickleResponse, error) {
	response, err := c.Post("/tickle", nil, nil)
	if err != nil {
		return nil, err
	}

	if response.statusCode != http.StatusOK {
		return nil, fmt.Errorf("bad logout statusCode: %v", response.statusCode)
	}

	var responseStruct TickleResponse
	err = c.ParseJsonResponse(response, &responseStruct)
	if err != nil {
		return nil, err
	}

	return &responseStruct, nil
}

package ibkr

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testLogoutResponse = `{
  "status":true
}`

var testInitializeBrokerageSessionResponse = `{
  "authenticated": true,
  "competing": false,
  "connected": true,
  "message": "",
  "MAC": "98:F2:B3:23:BF:A0",
  "serverInfo": {
    "serverName": "JifN19053",
    "serverVersion": "Build 10.25.0p, Dec 5, 2023 5:48:12 PM"
  }
}`

var testAuthStatusResponse = `{
  "authenticated": true,
  "competing": false,
  "connected": true,
  "message": "",
  "MAC": "12:B:B3:23:BF:A0",
  "serverInfo": {
    "serverName": "JifN19053",
    "serverVersion": "Build 10.25.0p, Dec 5, 2023 5:48:12 PM"
  },
  "fail": ""
}`

var testTickleResponse = `{
  "session": "bb665d0f55b6289d70bc7380089fc96f",
  "ssoExpires": 460311,
  "collission": false,
  "userId": 123456789,
  "hmds": {
    "error": "no bridge"
  },
  "iserver": {
    "authStatus": {
      "authenticated": true,
      "competing": false,
      "connected": true,
      "message": "",
      "MAC": "98:F2:B3:23:BF:A0",
      "serverInfo": {
        "serverName": "JifN19053",
        "serverVersion": "Build 10.25.0p, Dec 5, 2023 5:48:12 PM"
      }
    }
  }
}`

func TestIbkrWebClient_Logout(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, testLogoutResponse)
	}))
	defer mockServer.Close()

	client := NewIbkrWebClient(mockServer.URL, &MockOAuthContext{})
	err := client.Logout()

	assert.NoError(t, err)
}

func TestIbkrWebClient_InitializeBrokerageSession(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bodyBytes, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		var reqBody PlaceOrderRequest
		err = json.Unmarshal(bodyBytes, &reqBody)
		assert.NoError(t, err)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, testInitializeBrokerageSessionResponse)
	}))
	defer mockServer.Close()

	client := NewIbkrWebClient(mockServer.URL, &MockOAuthContext{})
	rsp, err := client.InitializeBrokerSession()

	assert.NoError(t, err)
	assert.NotNil(t, rsp)
}

func TestIbkrWebClient_AuthStatus(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, testAuthStatusResponse)
	}))
	defer mockServer.Close()

	client := NewIbkrWebClient(mockServer.URL, &MockOAuthContext{})
	rsp, err := client.AuthStatus()

	assert.NoError(t, err)
	assert.NotNil(t, rsp)
}

func TestIbkrWebClient_Tickle(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, testTickleResponse)
	}))
	defer mockServer.Close()

	client := NewIbkrWebClient(mockServer.URL, &MockOAuthContext{})
	rsp, err := client.Tickle()

	assert.NoError(t, err)
	assert.NotNil(t, rsp)
}

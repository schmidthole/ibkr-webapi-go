package ibkr

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testAcctId = "1234"

var testSwitchAccountResponse = `
{
	"set": true, 
	"acctId": "1234"
}`

var testSwitchAccountResponseError = `
{
	"set": false, 
	"acctId": "1234"
}`

func TestIbkrWebClient_SwitchAccount(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bodyBytes, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		var reqBody SwitchAccountRequest
		err = json.Unmarshal(bodyBytes, &reqBody)
		assert.NoError(t, err)

		assert.Equal(t, testAcctId, reqBody.AccountID)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, testSwitchAccountResponse)
	}))
	defer mockServer.Close()

	client := NewIbkrWebClient(mockServer.URL, &MockOAuthContext{})
	err := client.SwitchAccount(testAcctId)

	assert.NoError(t, err)
}

func TestIbkrWebClient_SwitchAccountError(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bodyBytes, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		var reqBody SwitchAccountRequest
		err = json.Unmarshal(bodyBytes, &reqBody)
		assert.NoError(t, err)

		assert.Equal(t, testAcctId, reqBody.AccountID)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, testSwitchAccountResponseError)
	}))
	defer mockServer.Close()

	client := NewIbkrWebClient(mockServer.URL, &MockOAuthContext{})
	err := client.SwitchAccount(testAcctId)

	assert.Error(t, err)
}

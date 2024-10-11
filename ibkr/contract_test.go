package ibkr

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testSearchContractBySymbolResponse = `
[
  {
    "conid": "43645865",
    "companyHeader": "IBKR INTERACTIVE BROKERS GRO-CL A (NASDAQ) ",
    "companyName": "INTERACTIVE BROKERS GRO-CL A (NASDAQ)",
    "symbol": "IBKR",
    "description": null,
    "restricted": null,
    "fop": null,
    "opt": null,
    "war": null,
    "sections": [],
    "secType": "STK"
  }
]`

func TestIbkrWebClient_SearchContractBySymbol(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, testSearchContractBySymbolResponse)
	}))
	defer mockServer.Close()

	client := NewIbkrWebClient(mockServer.URL, &MockOAuthContext{})
	rsp, err := client.SearchContractBySymbol("TEST")

	assert.NoError(t, err)
	assert.NotNil(t, rsp)
}

package ibkr

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testConid = 1234

var testMarketDataHistoryResponse = `{
  "startTime": "20231211-04:00:00",
  "data": [
    {
      "t": 1702285200000,
      "o": 195.01,
      "c": 194.8,
      "h": 195.01,
      "l": 194.8,
      "v": 1723.0
    }
  ],
  "points": 14,
  "mktDataDelay": 0
}`

var testMarketDataSnapshotResponse = `
[
  {
    "_updated": 1702334859712,
    "conidEx": "265598",
    "conid": 265598,
    "server_id": "q1",
    "55": "TEST",
    "31": "0.00",
    "70": "0.00",
    "71": "0.00",
    "73": "0.00",
    "80": "0.00",
    "7295": "0.00",
    "7296": "0.00",
    "7635": "0.00",
    "7741": "0.00"
  }
]`

func TestIbkrWebClient_MarketDataHistory(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, testMarketDataHistoryResponse)
	}))
	defer mockServer.Close()

	client := NewIbkrWebClient(mockServer.URL, &MockOAuthContext{})
	history, err := client.MarketDataHistory(testConid, "1y", "1d")

	assert.NotNil(t, history)
	assert.NoError(t, err)
}

func TestIbkrWebClient_MarketDataSnapshot(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, testMarketDataSnapshotResponse)
	}))
	defer mockServer.Close()

	client := NewIbkrWebClient(mockServer.URL, &MockOAuthContext{})
	history, err := client.MarketDataSnapshot([]int{testConid})

	assert.NotNil(t, history)
	assert.NoError(t, err)
}

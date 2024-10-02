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
  "startTimeVal": 1702285200000,
  "endTime": "20231211-17:51:20",
  "endTimeVal": 1702335080000,
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

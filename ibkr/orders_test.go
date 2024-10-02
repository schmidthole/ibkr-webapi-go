package ibkr

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testPlaceOrderResponsePlain = `[
  {
    "order_id": "1234567890",
    "order_status": "Submitted",
    "encrypt_message": "1"
  }
]`

var testPlaceOrderResponseMessage = `[
  {
    "id": "07a13a5a-4a48-44a5-bb25-5ab37b79186c",
    "message": [
      "The following order \"BUY 5 AAPL NASDAQ.NMS @ 150.0\" price exceeds \nthe Percentage constraint of 3%.\nAre you sure you want to submit this order?"
    ],
    "isSuppressed": false,
    "messageIds": [
      "o163"
    ]
  }
]`

var testPlaceOrderResponseReject = `{
  "error":"We cannot accept an order at the limit price you selected. Please submit your order using a limit price that is closer to the current market price of 197.79.  Alternatively, you can convert your order to an Algorithmic Order (IBALGO)."
}`

var testCancelOrderResponse = `{
    "msg": "Request was submitted",
    "order_id": 123456789,
    "conid": 265598,
    "account": "U1234567"
}`

var testGetLiveOrdersResponse = `{
  "orders": [
    {
      "acct": "U1234567",
      "conidex": "265598",
      "conid": 265598,
      "account": "U1234567",
      "orderId": 1234568790,
      "cashCcy": "USD",
      "sizeAndFills": "5",
      "orderDesc": "Sold 5 Market, GTC",
      "description1": "AAPL",
      "ticker": "AAPL",
      "secType": "STK",
      "listingExchange": "NASDAQ.NMS",
      "remainingQuantity": 0.0,
      "filledQuantity": 5.0,
      "totalSize": 5.0,
      "companyName": "APPLE INC",
      "status": "Filled",
      "order_ccp_status": "Filled",
      "avgPrice": "192.26",
      "origOrderType": "MARKET",
      "supportsTaxOpt": "1",
      "lastExecutionTime": "231211180049",
      "orderType": "Market",
      "bgColor": "#FFFFFF",
      "fgColor": "#000000",
      "order_ref": "Order123",
      "timeInForce": "GTC",
      "lastExecutionTime_r": 1702317649000,
      "side": "SELL"
    }
  ],
  "snapshot": true
}`

func TestIbkrWebClient_PlaceOrderPlain(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bodyBytes, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		var reqBody PlaceOrderRequest
		err = json.Unmarshal(bodyBytes, &reqBody)
		assert.NoError(t, err)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, testPlaceOrderResponsePlain)
	}))
	defer mockServer.Close()

	client := NewIbkrWebClient(mockServer.URL, &MockOAuthContext{})
	rsp, err := client.PlaceOrder("1234", Order{})

	assert.NotNil(t, rsp)
	assert.NoError(t, err)
}

func TestIbkrWebClient_PlaceOrderMessage(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bodyBytes, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		var reqBody PlaceOrderRequest
		err = json.Unmarshal(bodyBytes, &reqBody)
		assert.NoError(t, err)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, testPlaceOrderResponseMessage)
	}))
	defer mockServer.Close()

	client := NewIbkrWebClient(mockServer.URL, &MockOAuthContext{})
	rsp, err := client.PlaceOrder("1234", Order{})

	assert.NotNil(t, rsp)
	assert.NoError(t, err)
}

func TestIbkrWebClient_PlaceOrderReject(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bodyBytes, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		var reqBody PlaceOrderRequest
		err = json.Unmarshal(bodyBytes, &reqBody)
		assert.NoError(t, err)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, testPlaceOrderResponseReject)
	}))
	defer mockServer.Close()

	client := NewIbkrWebClient(mockServer.URL, &MockOAuthContext{})
	rsp, err := client.PlaceOrder("1234", Order{})

	assert.Nil(t, rsp)
	assert.Error(t, err)
}

func TestIbkrWebClient_CancelOrder(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, testCancelOrderResponse)
	}))
	defer mockServer.Close()

	client := NewIbkrWebClient(mockServer.URL, &MockOAuthContext{})
	rsp, err := client.CancelOrder("1234", "12345")

	assert.NotNil(t, rsp)
	assert.NoError(t, err)
}

func TestIbkrWebClient_GetLiveOrders(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, testGetLiveOrdersResponse)
	}))
	defer mockServer.Close()

	client := NewIbkrWebClient(mockServer.URL, &MockOAuthContext{})
	rsp, err := client.GetLiveOrders()

	assert.NotNil(t, rsp)
	assert.NoError(t, err)
}

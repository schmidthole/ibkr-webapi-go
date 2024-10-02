package ibkr

import (
	"fmt"
	"net/http"
)

/******************************************************************************
* place order
******************************************************************************/

type Order struct {
	AccountId   string  `json:"acctId"`
	ConID       int32   `json:"conid"`
	OrderType   string  `json:"orderType"`
	Side        string  `json:"side"`
	TimeInForce string  `json:"tif"`
	Quantity    float64 `json:"quantity"`
}

type PlaceOrderRequest struct {
	Orders []Order `json:"orders"`
}

type PlaceOrderResponsePlain struct {
	OrderID     string `json:"order_id" validation:"required"`
	OrderStatus string `json:"order_status" validation:"required"`
}

type PlaceOrderResponseMessage struct {
	ID           string   `json:"id" validation:"required"`
	Message      string   `json:"message" validation:"required"`
	IsSuppressed bool     `json:"isSupressed" validation:"required"`
	MessageIDs   []string `json:"messageIds" validation:"required"`
}

type PlaceOrderRejectResponse struct {
	Error string `json:"error" validation:"required"`
}

type PlaceOrderResponse struct {
	ID         string
	Status     string
	Message    string
	MessageIDs []string
}

func (c *IbkrWebClient) PlaceOrder(accountId string, order Order) (*PlaceOrderResponse, error) {
	requestBody := PlaceOrderRequest{Orders: []Order{order}}

	response, err := c.Post(fmt.Sprintf("/iserver/account/%s/orders", accountId), nil, requestBody)
	if err != nil {
		return nil, err
	}

	if response.statusCode != http.StatusOK {
		return nil, fmt.Errorf("place order bad statusCode: %v", response.statusCode)
	}

	var plainResponse []PlaceOrderResponsePlain
	err = c.ParseJsonResponse(response, &plainResponse)
	if err == nil {
		return &PlaceOrderResponse{
			ID:     plainResponse[0].OrderID,
			Status: plainResponse[0].OrderStatus,
		}, nil
	}

	var messageResponse []PlaceOrderResponseMessage
	err = c.ParseJsonResponse(response, &messageResponse)
	if err == nil {
		return &PlaceOrderResponse{
			ID:         messageResponse[0].ID,
			Message:    messageResponse[0].Message,
			MessageIDs: messageResponse[0].MessageIDs,
		}, nil
	}

	var rejectResponse PlaceOrderRejectResponse
	err = c.ParseJsonResponse(response, &rejectResponse)
	if err == nil {
		return nil, fmt.Errorf("place order rejected: %s", rejectResponse.Error)
	}

	return nil, fmt.Errorf("could not parse any possible response for place order")
}

/******************************************************************************
* cancel order
******************************************************************************/

type CancelOrderResponse struct {
	Message string `json:"msg" validation:"required"`
	OrderID int    `json:"order_id" validation:"required"`
	ConID   int    `json:"conid" validation:"required"`
	Account string `json:"account" validation:"required"`
}

type CancelOrderErrorResponse struct {
	Error string `json:"error" validation:"required"`
}

func (c *IbkrWebClient) CancelOrder(accountId string, orderId string) (*CancelOrderResponse, error) {
	response, err := c.Delete(fmt.Sprintf("/iserver/account/%s/order/%s", accountId, orderId), nil)
	if err != nil {
		return nil, err
	}

	if response.statusCode != http.StatusOK {
		return nil, fmt.Errorf("cancel order bad statusCode: %v", response.statusCode)
	}

	var responseStruct CancelOrderResponse
	err = c.ParseJsonResponse(response, &responseStruct)
	if err != nil {
		return nil, err
	}

	return &responseStruct, nil
}

/******************************************************************************
* live orders
******************************************************************************/

type LiveOrdersResponse struct {
	Orders []OrderStatus `json:"orders" validation:"required"`
}

type OrderStatus struct {
	Account           string  `json:"acct" validation:"required"`
	ConID             int32   `json:"conid" validation:"required"`
	OrderID           int32   `json:"orderId" validation:"required"`
	Ticker            string  `json:"ticker" validation:"required"`
	RemainingQuantity float64 `json:"remainingQuantity" validation:"required"`
	FilledQuantity    float64 `json:"filledQuantity" validation:"required"`
	Status            string  `json:"status" validation:"required"`
	OrderType         string  `json:"orderType" validation:"required"`
	Side              string  `json:"side" validation:"required"`
	TimeInForce       string  `json:"timeInForce" validation:"required"`
}

func (c *IbkrWebClient) GetLiveOrders() (*LiveOrdersResponse, error) {
	response, err := c.Get("/iserver/account/orders", nil)
	if err != nil {
		return nil, err
	}

	if response.statusCode != http.StatusOK {
		return nil, fmt.Errorf("get live orders bad statusCode: %v", response.statusCode)
	}

	var responseStruct LiveOrdersResponse
	err = c.ParseJsonResponse(response, &responseStruct)
	if err != nil {
		return nil, err
	}

	return &responseStruct, nil
}

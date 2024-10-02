package ibkr

import (
	"fmt"
	"net/http"
	"strconv"
)

/******************************************************************************
* market data history
******************************************************************************/

type MarketDataHistoryResponse struct {
	StartTime       string    `json:"startTime" validation:"required"`
	StartTimeVal    int       `json:"startTimeVal" validation:"required"`
	EndTime         string    `json:"endTime" validation:"required"`
	EndTimeVal      int       `json:"endTimeVal" validation:"required"`
	Data            []OHLCBar `json:"data" validation:"required"`
	Points          int       `json:"points" validation:"required"`
	MarketDataDelay int       `json:"mktDataDelay" validation:"required"`
}

type OHLCBar struct {
	T int     `json:"t" validation:"required"`
	O float64 `json:"o" validation:"required"`
	C float64 `json:"c" validation:"required"`
	H float64 `json:"h" validation:"required"`
	L float64 `json:"l" validation:"required"`
	V float64 `json:"v" validation:"required"`
}

func (c *IbkrWebClient) MarketDataHistory(
	conId int,
	period string,
	barType string,
) (*MarketDataHistoryResponse, error) {
	params := map[string]string{
		"conid":  strconv.Itoa(int(conId)),
		"period": period,
		"bar":    barType,
	}

	response, err := c.Get("/hmds/history", params)
	if err != nil {
		return nil, err
	}

	if response.statusCode != http.StatusOK {
		return nil, fmt.Errorf("bad market data history statusCode: %v", response.statusCode)
	}

	var responseStruct MarketDataHistoryResponse
	err = c.ParseJsonResponse(response, &responseStruct)
	if err != nil {
		return nil, err
	}

	return &responseStruct, nil
}

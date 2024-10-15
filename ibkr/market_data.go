package ibkr

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

/******************************************************************************
* market data history
******************************************************************************/

type MarketDataHistoryResponse struct {
	StartTime       string    `json:"startTime" validation:"required"`
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

const (
	MarketDataPeriodSeconds = "S"
	MarketDataPeriodDay     = "d"
	MarketDataPeriodWeek    = "w"
	MarketDataPeriodMonth   = "m"
	MarketDataPeriodYear    = "y"
)

const (
	MarketDataBarSeconds = "secs"
	MarketDataBarMinutes = "mins"
	MarketDataBarHours   = "hrs"
	MarketDataBarDay     = "d"
	MarketDataBarWeek    = "w"
	MarketDataBarMonth   = "m"
)

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

	response, err := c.Get("/iserver/marketdata/history", params)
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

/******************************************************************************
* market data snapshot
******************************************************************************/

const (
	snapshotFieldSymbol      = "55"
	snapshotFieldLastPrice   = "31"
	snapshotFieldHigh        = "70"
	snapshotFieldLow         = "71"
	snapshotFieldMarketValue = "73"
	snapshotFieldPnLPercent  = "80"
	snapshotFieldOpen        = "7295"
	snapshotFieldClose       = "7296"
	snapshotFieldMark        = "7635"
	snapshotFieldPriorClose  = "7741"
)

var marketDataSnapshotFields = []string{
	snapshotFieldSymbol,
	snapshotFieldLastPrice,
	snapshotFieldHigh,
	snapshotFieldLow,
	snapshotFieldMarketValue,
	snapshotFieldPnLPercent,
	snapshotFieldOpen,
	snapshotFieldClose,
	snapshotFieldMark,
	snapshotFieldPriorClose,
}

type MarketDataSnapshotResponse struct {
	ConID       int    `json:"conid" validation:"required"`
	Symbol      string `json:"55" validation:"required"`
	LastPrice   string `json:"31" validation:"required"`
	High        string `json:"70" validation:"required"`
	Low         string `json:"71" validation:"required"`
	MarketValue string `json:"73" validation:"required"`
	PnLPercent  string `json:"80" validation:"required"`
	Open        string `json:"7295" validation:"required"`
	Close       string `json:"7296" validation:"required"`
	Mark        string `json:"7635" validation:"required"`
	PriorClose  string `json:"7741" validation:"required"`
}

type MarketDataSnapshot struct {
	ConID         int
	Symbol        string
	TradingHalted bool
	TradingActive bool
	LastPrice     float64
	High          float64
	Low           float64
	MarketValue   float64
	PnLPercent    float64
	Open          float64
	Close         float64
	Mark          float64
	PriorClose    float64
}

func (c *IbkrWebClient) MarketDataSnapshot(
	conIds []int,
) ([]MarketDataSnapshot, error) {
	conIdParam := ""
	for i, conid := range conIds {
		if i == 0 {
			conIdParam = conIdParam + strconv.Itoa(conid)
		} else {
			conIdParam = conIdParam + "," + strconv.Itoa(conid)
		}
	}

	fieldsParam := ""
	for i, field := range marketDataSnapshotFields {
		if i == 0 {
			fieldsParam = fieldsParam + field
		} else {
			fieldsParam = fieldsParam + "," + field
		}
	}

	params := map[string]string{
		"conids": conIdParam,
		"fields": fieldsParam,
	}

	response, err := c.Get("/iserver/marketdata/snapshot", params)
	if err != nil {
		return nil, err
	}

	if response.statusCode != http.StatusOK {
		return nil, fmt.Errorf("bad market data history statusCode: %v", response.statusCode)
	}

	var responseStruct []MarketDataSnapshotResponse
	err = c.ParseJsonResponse(response, &responseStruct)
	if err != nil {
		return nil, err
	}

	snapshots := []MarketDataSnapshot{}
	for _, raw := range responseStruct {

		lastPriceString := raw.LastPrice
		var lastPricePrefix string
		if strings.HasPrefix(lastPriceString, "C") || strings.HasPrefix(lastPriceString, "H") {
			lastPricePrefix = lastPriceString[:1]
			lastPriceString = lastPriceString[1:]
		}

		tradingActive := true
		tradingHalted := false
		if lastPricePrefix == "C" {
			tradingActive = false
		} else if lastPricePrefix == "H" {
			tradingActive = false
			tradingHalted = true
		}

		lastPriceFloat, err := strconv.ParseFloat(lastPriceString, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing last price for conid %v, found: %v", raw.ConID, raw.LastPrice)
		}

		highFloat, err := strconv.ParseFloat(raw.High, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing high for conid %v, found: %v", raw.ConID, raw.High)
		}

		lowFloat, err := strconv.ParseFloat(raw.Low, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing low for conid %v, found: %v", raw.ConID, raw.Low)
		}

		marketValueFloat, err := strconv.ParseFloat(raw.MarketValue, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing market value for conid %v, found: %v", raw.ConID, raw.MarketValue)
		}

		pnlPctFloat, err := strconv.ParseFloat(raw.PnLPercent, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing pnl percent for conid %v, found: %v", raw.ConID, raw.PnLPercent)
		}

		openFloat, err := strconv.ParseFloat(raw.Open, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing open for conid %v, found: %v", raw.ConID, raw.Open)
		}

		closeFloat, err := strconv.ParseFloat(raw.Close, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing close for conid %v, found: %v", raw.ConID, raw.Close)
		}

		markFloat, err := strconv.ParseFloat(raw.Mark, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing mark for conid %v, found: %v", raw.ConID, raw.Mark)
		}

		priorCloseFloat, err := strconv.ParseFloat(raw.PriorClose, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing prior close for conid %v, found: %v", raw.ConID, raw.PriorClose)
		}

		snapshot := MarketDataSnapshot{
			ConID:         raw.ConID,
			Symbol:        raw.Symbol,
			TradingActive: tradingActive,
			TradingHalted: tradingHalted,
			LastPrice:     lastPriceFloat,
			High:          highFloat,
			Low:           lowFloat,
			MarketValue:   marketValueFloat,
			PnLPercent:    pnlPctFloat,
			Open:          openFloat,
			Close:         closeFloat,
			Mark:          markFloat,
			PriorClose:    priorCloseFloat,
		}

		snapshots = append(snapshots, snapshot)
	}

	return snapshots, nil
}

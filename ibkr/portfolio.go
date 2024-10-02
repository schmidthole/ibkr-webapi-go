package ibkr

import (
	"fmt"
	"net/http"
)

/******************************************************************************
* get subaccounts
******************************************************************************/

type PortfolioSubaccount struct {
	ID              string `json:"id" validation:"required"`
	Currency        string `json:"currency" validation:"required"`
	Type            string `json:"type" validation:"required"`
	BusinessType    string `json:"businessType" validation:"required"`
	IBEntity        string `json:"ibEntity" validation:"required"`
	ClearingStatus  string `json:"clearingStatus" validation:"required"`
	NoClientTrading bool   `json:"noClientTrading" validation:"required"`
}

func (c *IbkrWebClient) GetPortfolioSubaccounts() ([]PortfolioSubaccount, error) {
	response, err := c.Get("/portfolio/subaccounts", nil)
	if err != nil {
		return nil, err
	}

	if response.statusCode != http.StatusOK {
		return nil, fmt.Errorf("get portfolio subaccounts bad statusCode: %v", response.statusCode)
	}

	var responseStruct []PortfolioSubaccount
	err = c.ParseJsonResponse(response, &responseStruct)
	if err != nil {
		return nil, err
	}

	return responseStruct, nil
}

/******************************************************************************
* account ledger
******************************************************************************/

type PortfolioAccountLedger struct {
	Base AccountLedger `json:"BASE" validation:"required"`
}

type AccountLedger struct {
	SettledCash         float64 `json:"settledCash" validation:"required"`
	CashBalance         float64 `json:"cashBalance" validation:"required"`
	NetLiquidationValue float64 `json:"netLiquidationValue" validation:"required"`
	UnrealizedPnL       float64 `json:"unrealizedpnl" validation:"required"`
	RealizedPnL         float64 `json:"realizedpnl" validation:"required"`
	Funds               float64 `json:"funds" validation:"required"`
}

func (c *IbkrWebClient) GetPortfolioAccountLedger(acctId string) (*PortfolioAccountLedger, error) {
	response, err := c.Get(fmt.Sprintf("/portfolio/%s/ledger", acctId), nil)
	if err != nil {
		return nil, err
	}

	if response.statusCode != http.StatusOK {
		return nil, fmt.Errorf("bad portfolio account ledger statusCode: %v", response.statusCode)
	}

	var responseStruct PortfolioAccountLedger
	err = c.ParseJsonResponse(response, &responseStruct)
	if err != nil {
		return nil, err
	}

	return &responseStruct, nil
}

/******************************************************************************
* positions
******************************************************************************/

type Position struct {
	AccountID     string  `json:"acctId" validation:"required"`
	ConID         int32   `json:"conid" validation:"required"`
	ContractDesc  string  `json:"contractDesc" validation:"required"`
	Position      float64 `json:"position" validation:"required"`
	MarketPrice   float64 `json:"mktPrice" validation:"required"`
	MarketValue   float64 `json:"mktValue" validation:"required"`
	AveragePrice  float64 `json:"avgPrice" validation:"required"`
	AverageCost   float64 `json:"avgCost" validation:"required"`
	RealizedPnL   float64 `json:"realizedPnl" validation:"required"`
	UnrealizedPnL float64 `json:"unrealizedPnl" validation:"required"`
	Name          string  `json:"name" validation:"required"`
	Ticker        string  `json:"ticker" validation:"required"`
	PageSize      int32   `json:"pageSize" validation:"required"`
}

func (c *IbkrWebClient) GetPositions(acctId string, page int32) ([]Position, error) {
	response, err := c.Get(fmt.Sprintf("/portfolio/%s/positions/%d", acctId, page), nil)
	if err != nil {
		return nil, err
	}

	if response.statusCode != http.StatusOK {
		return nil, fmt.Errorf("bad get positions statusCode: %v", response.statusCode)
	}

	var responseStruct []Position
	err = c.ParseJsonResponse(response, &responseStruct)
	if err != nil {
		return nil, err
	}

	return responseStruct, nil
}

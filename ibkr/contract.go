package ibkr

import (
	"fmt"
	"net/http"
)

/******************************************************************************
* search contract by symbol
******************************************************************************/

type SearchContractBySymbolResponse struct {
	ConID       string `json:"conid" validate:"required"`
	CompanyName string `json:"companyName" validate:"required"`
	Symbol      string `json:"symbol" validate:"required"`
}

func (c *IbkrWebClient) SearchContractBySymbol(symbol string) ([]SearchContractBySymbolResponse, error) {
	params := map[string]string{
		"symbol":  symbol,
		"name":    "false",
		"secType": "STK",
	}

	response, err := c.Get("/iserver/secdef/search", params)
	if err != nil {
		return nil, err
	}

	if response.statusCode != http.StatusOK {
		return nil, fmt.Errorf("search contract by symbol bad statusCode: %v", response.statusCode)
	}

	var responseStruct []SearchContractBySymbolResponse
	err = c.ParseJsonResponse(response, &responseStruct)
	if err != nil {
		return nil, err
	}

	return responseStruct, nil
}

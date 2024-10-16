package ibkr

import (
	"fmt"
	"net/http"
)

/******************************************************************************
* get accounts
******************************************************************************/

type GetAccountsResponse struct {
	Accounts []string `json:"accounts" validate:"required"`
}

func (c *IbkrWebClient) GetAccounts() error {
	response, err := c.Get("/iserver/accounts", nil)
	if err != nil {
		return err
	}

	if response.statusCode != http.StatusOK {
		return fmt.Errorf("bad get accounts responseCode: %v", response.statusCode)
	}

	var responseStruct GetAccountsResponse
	err = c.ParseJsonResponse(response, &responseStruct)
	if err != nil {
		return err
	}

	return nil
}

/******************************************************************************
* switch account
******************************************************************************/

type SwitchAccountRequest struct {
	AccountID string `json:"acctId"`
}

type SwitchAccountResponse struct {
	Set       bool   `json:"set" validate:"required"`
	AccountID string `json:"acctId" validate:"required"`
}

func (c *IbkrWebClient) SwitchAccount(accountId string) error {
	requestBody := SwitchAccountRequest{
		AccountID: accountId,
	}

	response, err := c.Post("/iserver/account", nil, requestBody)
	if err != nil {
		return err
	}

	if response.statusCode != http.StatusOK {
		return fmt.Errorf("bad switch account responseCode: %v", response.statusCode)
	}

	var responseStruct SwitchAccountResponse
	err = c.ParseJsonResponse(response, &responseStruct)
	if err != nil {
		return err
	}

	if !responseStruct.Set || (responseStruct.AccountID != accountId) {
		return fmt.Errorf("ibkr error setting account")
	}

	return nil
}

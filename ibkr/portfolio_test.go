package ibkr

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testPortfolioSubaccountsResponse = `[
  {
    "id": "U1234567",
    "PrepaidCrypto-Z": false,
    "PrepaidCrypto-P": false,
    "brokerageAccess": false,
    "accountId": "U1234567",
    "accountVan": "U1234567",
    "accountTitle": "",
    "displayName": "U1234567",
    "accountAlias": null,
    "accountStatus": 1644814800000,
    "currency": "USD",
    "type": "DEMO",
    "tradingType": "PMRGN",
    "businessType": "IB_PROSERVE",
    "ibEntity": "IBLLC-US",
    "faclient": false,
    "clearingStatus": "O",
    "covestor": false,
    "noClientTrading": false,
    "trackVirtualFXPortfolio": true,
    "parent": {
      "mmc": [],
      "accountId": "",
      "isMParent": false,
      "isMChild": false,
      "isMultiplex": false
    },
    "desc": "U1234567"
  }
]`

var testPortfolioAccountLedger = `{
  "USD": {
    "commoditymarketvalue": 0.0,
    "futuremarketvalue": -1051.0,
    "settledcash": 214716688.0,
    "exchangerate": 1,
    "sessionid": 1,
    "cashbalance": 214716688.0,
    "corporatebondsmarketvalue": 0.0,
    "warrantsmarketvalue": 0.0,
    "netliquidationvalue": 215335840.0,
    "interest": 305569.94,
    "unrealizedpnl": 39695.82,
    "stockmarketvalue": 314123.88,
    "moneyfunds": 0.0,
    "currency": "USD",
    "realizedpnl": 0.0,
    "funds": 0.0,
    "acctcode": "U1234567",
    "issueroptionsmarketvalue": 0.0,
    "key": "LedgerList",
    "timestamp": 1702582321,
    "severity": 0,
    "stockoptionmarketvalue": -2.88,
    "futuresonlypnl": -1051.0,
    "tbondsmarketvalue": 0.0,
    "futureoptionmarketvalue": 0.0,
    "cashbalancefxsegment": 0.0,
    "secondkey": "USD",
    "tbillsmarketvalue": 0.0,
    "endofbundle": 1,
    "dividends": 0.0
  },
  "BASE": {
    "commoditymarketvalue": 0.0,
    "futuremarketvalue": -1051.0,
    "settledcash": 215100080.0,
    "exchangerate": 1,
    "sessionid": 1,
    "cashbalance": 215100080.0,
    "corporatebondsmarketvalue": 0.0,
    "warrantsmarketvalue": 0.0,
    "netliquidationvalue": 215721776.0,
    "interest": 305866.88,
    "unrealizedpnl": 39907.37,
    "stockmarketvalue": 316365.38,
    "moneyfunds": 0.0,
    "currency": "BASE",
    "realizedpnl": 0.0,
    "funds": 0.0,
    "acctcode": "U1234567",
    "issueroptionsmarketvalue": 0.0,
    "key": "LedgerList",
    "timestamp": 1702582321,
    "severity": 0,
    "stockoptionmarketvalue": -2.88,
    "futuresonlypnl": -1051.0,
    "tbondsmarketvalue": 0.0,
    "futureoptionmarketvalue": 0.0,
    "cashbalancefxsegment": 0.0,
    "secondkey": "BASE",
    "tbillsmarketvalue": 0.0,
    "dividends": 0.0
  }
}`

var testPortfolioPositions = `[
  {
    "acctId": "U1234567",
    "conid": 756733,
    "contractDesc": "SPY",
    "position": 5.0,
    "mktPrice": 471.16000365,
    "mktValue": 2355.8,
    "currency": "USD",
    "avgCost": 434.93,
    "avgPrice": 434.93,
    "realizedPnl": 0.0,
    "unrealizedPnl": 181.15,
    "exchs": null,
    "expiry": null,
    "putOrCall": null,
    "multiplier": null,
    "strike": 0.0,
    "exerciseStyle": null,
    "conExchMap": [],
    "assetClass": "STK",
    "undConid": 0,
    "model": ""
  },
  {
    "acctId": "U1234567",
    "conid": 76792991,
    "contractDesc": "TSLA",
    "position": 7.0,
    "mktPrice": 250.73399355,
    "mktValue": 1755.14,
    "currency": "USD",
    "avgCost": 221.67142855,
    "avgPrice": 221.67142855,
    "realizedPnl": 0.0,
    "unrealizedPnl": 203.44,
    "exchs": null,
    "expiry": null,
    "putOrCall": null,
    "multiplier": null,
    "strike": 0.0,
    "exerciseStyle": null,
    "conExchMap": [],
    "assetClass": "STK",
    "undConid": 0,
    "model": ""
  },
  {
    "acctId": "U1234567",
    "conid": 107113386,
    "contractDesc": "META",
    "position": 11.0,
    "mktPrice": 333.1199951,
    "mktValue": 3664.32,
    "currency": "USD",
    "avgCost": 306.6909091,
    "avgPrice": 306.6909091,
    "realizedPnl": 0.0,
    "unrealizedPnl": 290.72,
    "exchs": null,
    "expiry": null,
    "putOrCall": null,
    "multiplier": null,
    "strike": 0.0,
    "exerciseStyle": null,
    "conExchMap": [],
    "assetClass": "STK",
    "undConid": 0,
    "model": ""
  }
]`

func TestIbkrWebClient_PortfolioGetSubaccounts(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, testPortfolioSubaccountsResponse)
	}))
	defer mockServer.Close()

	client := NewIbkrWebClient(mockServer.URL, &MockOAuthContext{})
	rsp, err := client.GetPortfolioSubaccounts()

	assert.NotNil(t, rsp)
	assert.NoError(t, err)
}

func TestIbkrWebClient_PortfolioGetAccountLedger(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, testPortfolioAccountLedger)
	}))
	defer mockServer.Close()

	client := NewIbkrWebClient(mockServer.URL, &MockOAuthContext{})
	rsp, err := client.GetPortfolioAccountLedger("1234")

	assert.NotNil(t, rsp)
	assert.NoError(t, err)
}

func TestIbkrWebClient_PortfolioPositions(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, testPortfolioPositions)
	}))
	defer mockServer.Close()

	client := NewIbkrWebClient(mockServer.URL, &MockOAuthContext{})
	rsp, err := client.GetPositions("1234", 0)

	assert.NotNil(t, rsp)
	assert.NoError(t, err)
}

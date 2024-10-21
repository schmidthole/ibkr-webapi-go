package main

import (
	"flag"
	"log"

	"github.com/schmidthole/ibkr-webapi-go/ibkr"
)

func initClient(useOAuth bool) *ibkr.IbkrWebClient {
	if useOAuth {
		oauth, err := ibkr.NewIbkrOAuthContextFromFile("./cmd/credentials.yml")
		if err != nil {
			log.Printf("error initializing oauth context: %v", err)
			panic(err)
		}
		log.Printf("oauth context initialized")

		return ibkr.NewIbkrWebClient(ibkr.ProdBaseUrl, oauth)
	} else {
		return ibkr.NewIbkrWebClient(ibkr.GatewayBaseUrl, nil)
	}
}

func main() {
	log.Println("ibkr web client test")

	oauthFlag := flag.Bool("oauth", false, "authenticate directly to ibkr with oauth, otherwise gateway will be used")
	flag.Parse()

	client := initClient(*oauthFlag)
	client.VerboseLogging = true
	log.Printf("client initialized, baseurl: %v", client.BaseUrl)

	if *oauthFlag {
		log.Println("starting oauth")
		err := client.Authenticate()
		if err != nil {
			panic(err)
		}

		_, err = client.InitializeBrokerSession()
		if err != nil {
			panic(err)
		}
	}

	_, err := client.Tickle()
	if err != nil {
		panic(err)
	}

	_, err = client.GetAccounts()
	if err != nil {
		panic(err)
	}

	err = client.SwitchAccount("DUA717516")
	if err != nil {
		panic(err)
	}

	_, err = client.SearchContractBySymbol("TQQQ")
	if err != nil {
		panic(err)
	}

	err = client.SuppressMessages()
	if err != nil {
		panic(err)
	}

	_, err = client.GetPortfolioAccountLedger("DUA717516")
	if err != nil {
		panic(err)
	}

	// _, err = client.MarketDataHistory(72539702, "1y", "1d")
	// if err != nil {
	// 	panic(err)
	// }

	_, err = client.MarketDataSnapshot([]int{72539702})
	if err != nil {
		panic(err)
	}

}

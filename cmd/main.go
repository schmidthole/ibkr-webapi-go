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
	log.Printf("client initialized, baseurl: %v", client.BaseUrl)

	if *oauthFlag {
		log.Println("starting oauth")
		err := client.Authenticate()
		if err != nil {
			panic(err)
		}
	}

	_, err := client.Tickle()
	if err != nil {
		panic(err)
	}
}

package main

import (
	"log"

	"github.com/schmidthole/ibkr-webapi-go/ibkr"
)

func main() {
	log.Println("ibkr web client test")

	oauth, err := ibkr.NewIbkrOAuthContextFromFile("./cmd/credentials.yml")
	if err != nil {
		log.Printf("error initializing oauth context: %v", err)
		panic(err)
	}
	log.Printf("oauth context initialized")

	client := ibkr.NewIbkrWebClient(ibkr.ProdBaseUrl, oauth)
	log.Printf("client initialized, baseurl: %v", client.BaseUrl)

	log.Println("starting oauth")
	err = client.Authenticate()
	if err != nil {
		panic(err)
	}
}

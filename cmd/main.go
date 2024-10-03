package main

import (
	"github.com/schmidthole/ibkr-webapi-go/ibkr"
	"log"
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
}

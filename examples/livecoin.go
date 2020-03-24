package main

import (
	"fmt"

	"github.com/Tomiyou/go-livecoin"
)

const (
	API_KEY    = ""
	API_SECRET = ""
)

func main() {
	// livecoin client
	livecoin := livecoin.New(API_KEY, API_SECRET)

	// GetBalances
	balances, _ := livecoin.GetBalances()
	fmt.Println(len(balances))

	for i, _ := range balances {
		if balances[i].Currency == "BTC" {
			fmt.Println(balances[i].Value)
		}
	}

}

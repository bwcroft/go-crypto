package main

import (
	"fmt"
	"github.com/bwcroft/go-crypto/api/coinGecko"
	"github.com/bwcroft/go-crypto/handlers"
)

func printCoins() {
	coins, err := coinGecko.GetCoins()

	if err != nil {
		println("Unable to fetch coins")
	}

	for _, coin := range coins {
		fmt.Printf("Name: %s, Symbol: %s\n", coin.Name, coin.Symbol)
	}
}

func main() {
	handlers.PrintCoins()
}

package handlers

import (
	"fmt"
	"github.com/bwcroft/go-crypto/api/coinGecko"
)

var validIDs = []string{"bitcoin", "ethereum", "litecoin", "dogecoin", "solana", "ripple", "helium"}

func PrintCoins() {
	matched := 0
	coins, err := coinGecko.GetCoins()

	if err != nil {
		fmt.Println("Unable to get Coin Gecko Coins")
	}

	for _, coin := range coins {
		if len(validIDs) == 0 {
			break
		}

		for index, vs := range validIDs {
			if coin.ID == vs {
				matched++
				validIDs = append(validIDs[:index], validIDs[index+1:]...)
				fmt.Printf("ID: %-10s Name: %-10s Symbol: %-10s\n", coin.ID, coin.Name, coin.Symbol)
				break
			}
		}
	}
}

package coinGecko

import (
	"encoding/json"
	"fmt"
)

type Coin struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

func GetCoins() ([]Coin, error) {
	response, err := ClientCoinGecko.Get("/coins/list")
	if err != nil {
		return nil, err
	}

	var coins []Coin
	err = json.NewDecoder(response.Body).Decode(&coins)
	if err != nil {
		return nil, err
	}

	return coins, nil
}

func GetCoin(id string) (*Coin, error) {
	url := fmt.Sprintf("/coins/%v", id)

	response, err := ClientCoinGecko.Get(url)
	if err != nil {
		return nil, err
	}

	var coin *Coin
	err = json.NewDecoder(response.Body).Decode(&coin)
	if err != nil {
		return nil, err
	}

	return coin, nil
}

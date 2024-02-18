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

type CoinFull struct {
	Coin
	WebSlug       string     `json:"web_slug"`
	AssetPlatform string     `json:"asset_platform"`
	Algorithm     string     `json:"hashing_algorithm"`
	MarketData    MarketData `json:"market_data"`
}

type MarketData struct {
	CurrentPrice CurrentPrice `json:"current_price"`
}

type CurrentPrice struct {
	USD uint32 `json:"usd"`
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

func GetCoin(id string) (*CoinFull, error) {
	url := fmt.Sprintf("/coins/%v", id)

	response, err := ClientCoinGecko.Get(url)
	if err != nil {
		return nil, err
	}

	var coin *CoinFull
	err = json.NewDecoder(response.Body).Decode(&coin)
	if err != nil {
		return nil, err
	}

	return coin, nil
}

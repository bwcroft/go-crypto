package coinGecko

func Ping() bool {
	response, err := ClientCoinGecko.Get("/ping")
	if err != nil {
		return false
	}
	return response.StatusCode == 200
}

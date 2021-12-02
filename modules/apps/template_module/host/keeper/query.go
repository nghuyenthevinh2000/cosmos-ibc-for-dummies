package keeper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CoinRes struct {
	Price float64 `json:"price"`
}

// get price of one coin
// get price over ibc packet
func (q Keeper) GetCoinPrice(coin string) (float64, error) {
	url := "https://api-osmosis.imperator.co/tokens/v1/price/" + coin
	coinRes := new(CoinRes)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("fail to get coin price")
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	//decode resquest to json
	json.Unmarshal(bodyBytes, &coinRes)

	//return price of a coin

	return coinRes.Price, nil
}

// get price of both coins
func (q Keeper) GetPairPrice(pair [2]string) ([2]float64, error) {
	var pairPrice [2]float64

	for i := 0; i < 2; i++ {
		price, err := q.GetCoinPrice(pair[i])
		if err != nil {
			return [2]float64{}, err
		}

		pairPrice[i] = price
	}

	return pairPrice, nil
}

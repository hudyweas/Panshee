package coingeckoapi

import (
	"fmt"
	"strings"

	"github.com/hudyweas/panshee/services/wallet_service/external/httpmethods"
)

const (
	getTokenPrices = `https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=%v&order=market_cap_desc&per_page=250&page=1&sparkline=false&price_change_percentage=24h`
)

func GetTokenPriceFromTokenList(tokens []string) (f []TokenPrice, err error) {
	tokenString := ""
	for i, token := range tokens {
		token = strings.ToLower(token)
		tokenString += strings.ReplaceAll(token, " ", "-")

		if i != len(token) - 1{
			tokenString += "%2C"
		}
	}

	req := fmt.Sprintf(getTokenPrices, tokenString)

	price := []TokenPrice{}
	if err = httpmethods.GetHttpJsonAndDecode(req, &price); err != nil {
		return nil, err
	}

	return price, nil
}

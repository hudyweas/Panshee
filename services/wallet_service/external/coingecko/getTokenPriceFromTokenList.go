package coingeckoapi

import (
	"fmt"
	"strings"

	"github.com/hudyweas/panshee/services/wallet_service/external/httpmethods"
)

const (
	getTokenID     = "https://api.coingecko.com/api/v3/coins/list?include_platform=false"
	getTokenPrices = `https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=%v&order=market_cap_desc&per_page=250&page=1&sparkline=false&price_change_percentage=24h`
)

func GetTokenIDs() ([]TokenID, map[string]int, error) {
	ids := []TokenID{}
	if resp, err := httpmethods.GetHttpJsonAndDecode(getTokenID, &ids); err != nil {
		if (resp.StatusCode == 429) {
			return nil, nil, fmt.Errorf("Please try again in a minute. Too many response to coingecko api")
		}
		return nil, nil, err
	}

	tokenIDmap := make(map[string]int)

	for i, ti := range ids {
		tokenIDmap[ti.Symbol] = i + 1
	}

	return ids, tokenIDmap, nil
}

func GetTokenPriceFromTokenSymbolList(tokenSymbols []string) ([]TokenPrice, error) {
	tokenIDs, tokenMap, err := GetTokenIDs()
	if err != nil {
		return nil, err
	}

	tokenIdString := ""
	for i, token := range tokenSymbols {
		token = strings.ToLower(token)
		id := tokenMap[token]
		if id == 0 {
			continue
		}

		tokenIdString += strings.ReplaceAll(tokenIDs[id-1].ID, " ", "-")

		if i != len(token)-1 {
			tokenIdString += "%2C"
		}
	}

	req := fmt.Sprintf(getTokenPrices, tokenIdString)

	price := []TokenPrice{}
	if resp, err := httpmethods.GetHttpJsonAndDecode(req, &price); err != nil {
		if (resp.StatusCode == 429) {
			return nil, fmt.Errorf("Please try again in a minute. Too many response to coingecko api")
		}
		return nil, err
	}

	return price, nil
}

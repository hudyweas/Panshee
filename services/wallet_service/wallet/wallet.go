package wallet

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	bscscanapi "github.com/hudyweas/panshee/services/wallet_service/external/bscscan"
	coingeckoapi "github.com/hudyweas/panshee/services/wallet_service/external/coingecko"
)

type Wallet struct {
	Address  string
	Currency []Currency
}

type Currency struct {
	TokenName 	  				string
	TokenSymbol					string
	Amount						float64
	PriceUSD 					float64
	PriceChangePercentage24h 	float64
}

func NewWalletFromBep20List(hash string, bep20list []bscscanapi.Bep20Trans) (w Wallet, err error) {
	w.Address = hash
	err = w.UpdateFromBep20ListTransaction(bep20list)

	return
}

func (w *Wallet) UpdateFromBep20ListTransaction(bep20list []bscscanapi.Bep20Trans) error{
	currencyMap := make(map[string]int)

	for _, trans := range bep20list{
		id, ok := currencyMap[trans.TokenName]
		if !ok {
			id = len(w.Currency)
			currencyMap[trans.TokenName] = id

			c := new(Currency)

			c.TokenName = trans.TokenName
			c.TokenSymbol = trans.TokenSymbol
			c.Amount = 0

			w.Currency = append(w.Currency, *c)
		}

		transFloat, err := strconv.ParseFloat(trans.Value, 64)
		if err != nil{
			return fmt.Errorf("error during converting result to float")
		}
		tokenDecimal, err := strconv.Atoi(trans.TokenDecimal)
		if err != nil{
			return fmt.Errorf("error during converting result to float")
		}

		transFloat /= math.Pow10(tokenDecimal)

		if strings.EqualFold(trans.To, w.Address) {
			w.Currency[id].Amount += transFloat
		}else{
			w.Currency[id].Amount -= transFloat
		}
	}

	return nil
}

func (w *Wallet) SetUpTokenPrices() error {
	var tokens []string
	currencyMap := make(map[string]int)

	for i, c := range w.Currency {
		tokens = append(tokens, c.TokenName)
		currencyMap[c.TokenName] = i
	}

	prices, err := coingeckoapi.GetTokenPriceFromTokenList(tokens)
	if err != nil {
		return err
	}

	for _, tokenPrice := range prices {
		id := currencyMap[tokenPrice.Name]
		if(w.Currency[id].TokenName == tokenPrice.Name){
			w.Currency[id].PriceUSD = tokenPrice.CurrentPrice
			w.Currency[id].PriceChangePercentage24h = tokenPrice.PriceChangePercentage24h
		}
	}

	return nil
}

func (w *Wallet) ClearEmptyAndScamCurrency(){
	noEmptyCurrency := *new([]Currency)

	for _, c := range w.Currency {
		if c.Amount != 0.0 && c.PriceUSD != 0.0 {
			noEmptyCurrency = append(noEmptyCurrency, c)
		}
	}

	w.Currency = noEmptyCurrency
}

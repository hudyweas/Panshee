package coingeckoapi

type TokensPriceResponse struct {
	Prices []TokenPrice `json:""`
}

type TokenPrice struct {
	Id                                 string  `json:"id"`
	Symbol                             string  `json:"symbol"`
	Name                               string  `json:"name"`
	Image                              string  `json:"image"`
	CurrentPrice                       float64 `json:"current_price"`
	MarketCap                          float64 `json:"market_cap"`
	MarketCapRank                      float64 `json:"market_cap_rank"`
	FullyDilutedValuation              float64 `json:"fully_diluted_valuation"`
	TotalVolume                        float64 `json:"total_volume"`
	High24h                            float64 `json:"high_24h"`
	Low24h                             float64 `json:"low_24h"`
	PriceChange24h                     float64 `json:"price_change_24h"`
	PriceChangePercentage24h           float64 `json:"price_change_percentage_24h"`
	MarketCapChange24h                 float64 `json:"market_cap_change_24h"`
	MarketCapChangePercentage24h       float64 `json:"market_cap_change_percentage_24h"`
	CirculatingSupply                  float64 `json:"circulating_supply"`
	TotalSupply                        float64 `json:"total_supply"`
	MaxSupply                          float64 `json:"max_supply"`
	Ath                                float64 `json:"ath"`
	AthChangePercentage                float64 `json:"ath_change_percentage"`
	AthDate                            string  `json:"ath_date"`
	Atl                                float64 `json:"atl"`
	AtlChangePercentage                float64 `json:"atl_change_percentage"`
	AtlDate                            string  `json:"atl_date"`
	Roi                                string  `json:"roi"`
	LastUpdated                        string  `json:"last_updated"`
	PriceChangePercentage24hInCurrency float64 `json:"price_change_percentage_24h_in_currency"`
}

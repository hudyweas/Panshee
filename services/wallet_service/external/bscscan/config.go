package bscscanapi

import "net/http"

type bscscanapi struct {
	key    string
	client *http.Client
}

type BscAPI interface {
	GetBnbBalanceFromAddress(walletAddress string) (f float64, err error)
	GetBep20ListFromAddress(walletAddress string) (bep20 []Bep20Trans, err error)
}

func CreateBscScanAPI(key string) bscscanapi {
	api := bscscanapi{
		key:    key,
		client: &http.Client{},
	}

	return api
}

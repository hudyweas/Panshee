package bscscanapi

import (
	"fmt"
	"strconv"

	"github.com/hudyweas/panshee/services/wallet_service/external/httpmethods"
)

const (
	getBnbBalace = `https://api.bscscan.com/api?module=account&action=balance&address=%v&apikey=%v`
)

func (bsc bscscanapi) GetBnbBalanceFromAddress(walletAddress string) (f float64, err error) {
	req := fmt.Sprintf(getBnbBalace, walletAddress, bsc.key)

	bnb := BnbBalance{}
	if _, err = httpmethods.GetHttpJsonAndDecode(req, &bnb); err != nil{
		return
	}

	return strconv.ParseFloat(bnb.Amount, 64)
}

package bscscanapi

import (
	"fmt"

	"github.com/hudyweas/panshee/services/wallet_service/external/httpmethods"
)

const (
	getBep20List = `https://api.bscscan.com/api?module=account&action=tokentx&address=%v&startblock=0&endblock=99999999&page=1&offset=1000&sort=desc&apikey=%v`
)

func (bsc bscscanapi) GetBep20ListFromAddress(walletAddress string) (bep20 []Bep20Trans, err error) {
	req := fmt.Sprintf(getBep20List, walletAddress, bsc.key)

	bep20response := Bep20TransListResponse{}
	if err = httpmethods.GetHttpJsonAndDecode(req, &bep20response); err != nil{
		return
	}

	bep20 = bep20response.Result

	return
}

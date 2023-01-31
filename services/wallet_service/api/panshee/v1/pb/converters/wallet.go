package converters

import (
	"github.com/hudyweas/panshee/services/wallet_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/wallet_service/wallet"
)

func PbWalletfromWallet(w wallet.Wallet) *pb.Wallet {
	pbw := &pb.Wallet{
		Address: w.Address,
		Tokens:  []*pb.Currency{},
	}

	pbc := []*pb.Currency{}

	for _, c := range w.Currency {
		pbc = append(pbc, PbCurrencyFromCurrency(c))
	}

	pbw.Tokens = pbc

	return pbw
}

func PbCurrencyFromCurrency(c wallet.Currency) (pbc *pb.Currency) {
	pbc = &pb.Currency{
		TokenName:                c.TokenName,
		TokenSymbol:              c.TokenSymbol,
		Amount:                   float32(c.Amount),
		PriceUSD:                 float32(c.PriceUSD),
		PriceChangePercentage24H: float32(c.PriceChangePercentage24h),
	}

	return
}

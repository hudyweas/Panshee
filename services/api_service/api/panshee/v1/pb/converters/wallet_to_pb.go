package converters

import (
	walletPb "github.com/hudyweas/panshee/services/api_service/services/wallet_service/pb"

	"github.com/hudyweas/panshee/services/api_service/api/panshee/v1/pb"
)

func ConvertWalletToPbWallet(wallet *walletPb.Wallet) *pb.WalletDTO {
	pbWallet := pb.WalletDTO{
		Address: wallet.Address,
	}

	for _, c := range wallet.Tokens {
		pbWallet.Tokens = append(pbWallet.Tokens, &pb.CurrencyDTO{
			TokenName:                c.TokenName,
			TokenSymbol:              c.TokenSymbol,
			Amount:                   c.Amount,
			PriceUSD:                 c.PriceUSD,
			PriceChangePercentage24H: c.PriceChangePercentage24H,
		})
	}

	return &pbWallet
}

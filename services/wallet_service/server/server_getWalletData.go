package server

import (
	"context"
	"fmt"

	"github.com/hudyweas/panshee/services/wallet_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/wallet_service/api/panshee/v1/pb/converters"
	val "github.com/hudyweas/panshee/services/wallet_service/internal/validators"
	"github.com/hudyweas/panshee/services/wallet_service/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetWalletData(ctx context.Context, req *pb.GetWalletDataRequest) (*pb.Wallet, error) {
	if validationErrors := validateGetWalletRequest(req); len(validationErrors) > 0 {
		return nil, status.Errorf(codes.InvalidArgument, validationErrors)
	}

	//TODO: implement bnb
	bnbBalance, err := s.bsc.GetBnbBalanceFromAddress(req.GetWalletAddress())
	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	bnbCurrency := wallet.Currency{
		TokenName:   "BNB",
		TokenSymbol: "BNB",
		Amount:      bnbBalance,
	}

	bep20list, err := s.bsc.GetBep20ListFromAddress(req.GetWalletAddress())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	if len(bep20list) == 0{
		return nil, status.Errorf(codes.InvalidArgument, "no matching wallet for provided address: %s", req.GetWalletAddress())
	}

	w, err := wallet.NewWalletFromBep20List(req.GetWalletAddress(), bep20list)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	w.Currency = append(w.Currency, bnbCurrency)

	err = w.SetUpTokenPrices()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	w.ClearEmptyAndScamCurrency()

	res := converters.PbWalletfromWallet(w)

	return res, nil
}

func validateGetWalletRequest(req *pb.GetWalletDataRequest) (errors string) {
	if err := val.ValidateWalletAddress(req.GetWalletAddress()); err != nil {
		errors += err.Error() + "\n"
	}

	return
}

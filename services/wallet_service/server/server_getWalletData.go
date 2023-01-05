package server

import (
	"context"
	"fmt"

	"github.com/hudyweas/panshee/services/wallet_service/api/panshee/v1/pb"
	val "github.com/hudyweas/panshee/services/wallet_service/internal/validators"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetWalletData(ctx context.Context, req *pb.GetWalletDataRequest) (*pb.GetWalletDataResponse, error){
	if validationErrors := validateGetWalletRequest(req); len(validationErrors) > 0{
		return nil, status.Errorf(codes.InvalidArgument, validationErrors)
	}

	_, err := s.bsc.GetBnbBalanceFromAddress(req.GetWalletAddress())
	if err != nil{
		fmt.Println(err)
	}
	_, err = s.bsc.GetBep20ListFromAddress(req.GetWalletAddress())
	if err != nil{
		fmt.Println(err)
	}
	res := &pb.GetWalletDataResponse{}

	return res, nil
}

func validateGetWalletRequest(req *pb.GetWalletDataRequest) (errors string){

	if err := val.ValidateWalletAddress(req.GetWalletAddress()); err != nil{
		errors += err.Error() + "\n"
	}

	return
}

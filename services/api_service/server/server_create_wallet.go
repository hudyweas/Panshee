package server

import (
	"context"

	"github.com/google/uuid"
	"github.com/hudyweas/panshee/services/api_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/api_service/database"
	"github.com/hudyweas/panshee/services/api_service/internal/contextData"
	val "github.com/hudyweas/panshee/services/api_service/internal/validators"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateWallet(ctx context.Context, req *pb.CreateWalletRequest) (*pb.WalletDTO, error) {
	//checking request validation
	if validationErrors := validateCreateWalletRequest(req); len(validationErrors) > 0{
		return nil, status.Errorf(codes.InvalidArgument, validationErrors)
	}

	user_id, err := uuid.Parse(contextData.GetAccessPayload(ctx).GetUserID())
	if err != nil{
		return nil, status.Errorf(codes.Internal, "error with user id")
	}

	createdWallet, err := s.db.WalletCreate(database.Wallet{
		WalletAddress: req.Wallet.GetAddress(),
		UserID:        user_id,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user %s:", err)
	}

	res := pb.WalletDTO{ Address: createdWallet.WalletAddress}

	return &res, nil
}

func validateCreateWalletRequest(req *pb.CreateWalletRequest) (errors string){
	//validating email from request
	if err := val.ValidateWalletAddress(req.Wallet.GetAddress()); err != nil {
		errors += "invalid wallet address\n"
	}

	return
}

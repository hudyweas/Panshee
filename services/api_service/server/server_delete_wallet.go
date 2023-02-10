package server

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/hudyweas/panshee/services/api_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/api_service/internal/contextData"
	val "github.com/hudyweas/panshee/services/api_service/internal/validators"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteWallet(ctx context.Context, req *pb.DeleteWalletRequest) (*pb.DeleteWalletResponse ,error) {
	//checking request validation
	validationErrors, user_id, walletAddress := validateDeleteWalletRequest(req);
	if len(validationErrors) > 0 {
		return nil, status.Errorf(codes.InvalidArgument, validationErrors)
	}

	payloadId, err := uuid.Parse(contextData.GetAccessPayload(ctx).UserID)
	if err != nil{
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if user_id != payloadId {
		return nil, status.Errorf(codes.InvalidArgument, "user: %s don't have access to delete wallet: %s", user_id, walletAddress)
	}

	err = s.db.DeleteWalletByIdAndWalletAddress(user_id, walletAddress)
	if err != nil{
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	res := &pb.DeleteWalletResponse{
		Success: true,
	}

	return res, nil
}

func validateDeleteWalletRequest(req *pb.DeleteWalletRequest) (errors string, user_id uuid.UUID, walletAddress string){
	parentArray := strings.Split(req.Parent, "/")
	user_id, err := uuid.Parse(parentArray[1])
	if err != nil{
		errors += "wrong user id\n"
	}

	walletAddress = parentArray[3]
	err = val.ValidateWalletAddress(walletAddress)
	if err != nil{
		errors += "wrong wallet address\n"
	}

	return
}

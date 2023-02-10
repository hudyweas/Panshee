package server

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	e "github.com/hudyweas/panshee/services/api_service/errors"
	walletPb "github.com/hudyweas/panshee/services/api_service/services/wallet_service/pb"

	"github.com/hudyweas/panshee/services/api_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/api_service/api/panshee/v1/pb/converters"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//TODO:error and wrong input checking

func (s *Server) GetWallets(ctx context.Context, req *pb.GetWalletDTODataRequest) (*pb.GetWalletDTODataResponse, error) {
	//checking request validation
	if validationErrors := validateGetWalletDataRequest(req); len(validationErrors) > 0{
		return nil, status.Errorf(codes.InvalidArgument, validationErrors)
	}

	userID, walletAddresses := getUserIdAndWalletAdressesFromParent(req.GetParent())
	userUUID, err := uuid.Parse(userID)
	if err == nil{
		_, err = s.db.UserSelectById(userUUID)
		if err != nil {
			if err.Error() == e.ErrNoDataInDatabase {
				return nil, status.Errorf(codes.NotFound, "no user with this id: %s in database", userID)
			}

			if len(walletAddresses) == 0 {//it means that for sure we wanted wallets from this username
				return nil, err
			}
		}else{
			dbWalletAddress, err := s.db.WalletSelectByID(userUUID)
			fmt.Println(len(*dbWalletAddress))
			if err != nil{
				if err.Error() == e.ErrNoDataInDatabase {
					return nil, status.Errorf(codes.NotFound, "no wallet addresses for this user: %s in database", userID)
				}
			}else{
				walletAddresses = []string{}
				for _, address := range *dbWalletAddress{
					walletAddresses = append(walletAddresses, address.WalletAddress)
				}
			}
		}
	}


	if len(walletAddresses) == 0 || walletAddresses[0] == "" {
		return nil, status.Errorf(codes.InvalidArgument, "no wallets provided")
	}

	wallets := []*pb.WalletDTO{}
	for _, walletAddress := range walletAddresses {
		wallet, err := s.services.WalletService.GetWalletData(ctx, &walletPb.GetWalletDataRequest{
			WalletAddress: walletAddress,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error during getting wallet data: %s", err)
		}else{
			wallets = append(wallets, converters.ConvertWalletToPbWallet(wallet))
		}
	}

	res := &pb.GetWalletDTODataResponse{
		Wallets: wallets,
	}

	return res, nil
}

//parent=user/*/wallets/*
//wallets are separeted with comma (",")
func getUserIdAndWalletAdressesFromParent(parent string) (string, []string){
	parentArray := strings.Split(parent, "/")
	walletsNotSeparated := parentArray[len(parentArray) - 1]
	walletAddresses := strings.Split(walletsNotSeparated, ",")

	return parentArray[1], walletAddresses
}

func validateGetWalletDataRequest(req *pb.GetWalletDTODataRequest) (errors string){
	return
}

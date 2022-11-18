package server

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/hudyweas/panshee/services/api_service/api/panshee/v1/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	authPb "github.com/hudyweas/panshee/services/api_service/services/auth_service/pb"
)

func (s *Server) RefreshAccessToken(ctx context.Context, req *pb.RefreshAccessTokenDTORequest) (*pb.RefreshAccessTokenDTOResponse, error) {
	//checking request validation
	if validationErrors := validateRefreshAccessToken(req); len(validationErrors) > 0{
		return nil, status.Errorf(codes.InvalidArgument, validationErrors)
	}

	reqUserID := strings.Split(req.GetParent(), "/")[1]

	authServiceRes, err := s.services.AuthService.RefreshAccessToken(context.Background(), &authPb.RefreshAccessTokenRequest{
		UserId:       reqUserID,
		RefreshToken: &authPb.Token{
			Value:          req.RefreshToken.GetValue(),
		},
	})
	if err != nil {
		return nil, err
	}

	res := &pb.RefreshAccessTokenDTOResponse{
		AccessToken: &pb.TokenDTO{
			Value:          authServiceRes.AccessToken.GetValue(),
			ExpirationTime: authServiceRes.AccessToken.GetExpirationTime(),
		},
	}

	return res, nil
}

func validateRefreshAccessToken(req *pb.RefreshAccessTokenDTORequest) (errors string){
	//validating email from request
	parent := strings.Split(req.GetParent(), "/")
	if(len(parent) < 2){
		errors += "invalid parent\n"
	}

	if _,err := uuid.Parse(parent[1]); err != nil{
		errors += "invalid user id\n"
	}

	return
}

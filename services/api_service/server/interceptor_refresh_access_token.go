package server

import (
	"context"
	"fmt"

	"github.com/hudyweas/panshee/services/api_service/api/panshee/v1/pb"
)

func (s *Server) RefreshAccessToken(ctx context.Context, req *pb.RefreshAccessTokenDTORequest) (*pb.RefreshAccessTokenDTOResponse, error) {
	validateRefreshAccessToken(req)

	res := &pb.RefreshAccessTokenDTOResponse{
		AccessToken: &pb.TokenDTO{},
	}

	return res, nil
}

func validateRefreshAccessToken(req *pb.RefreshAccessTokenDTORequest) (errors string){
	//validating email from request
	fmt.Print(req.GetParent())

	return
}

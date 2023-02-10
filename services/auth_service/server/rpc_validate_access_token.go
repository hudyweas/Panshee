package server

import (
	"context"
	"fmt"

	"github.com/hudyweas/panshee/services/auth_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/auth_service/api/panshee/v1/pb/converters"
)

func (s *Server) ValidateAccessToken(ctx context.Context, req *pb.ValidateAccessTokenRequest) (*pb.ValidateAccessTokenResponse, error) {
	accessTokenPayload, err := s.accessTokenGenerator.VerifyToken(req.AccessToken.GetValue())
	if err != nil {
		return nil, fmt.Errorf("invalid token")
	}

	res := &pb.ValidateAccessTokenResponse{
		AccessTokenPayload: converters.PbPayloadFromTokenPayload(accessTokenPayload),
	}
	return res, nil
}

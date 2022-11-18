package server

import (
	"context"
	"fmt"

	"github.com/hudyweas/panshee/services/auth_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/auth_service/api/panshee/v1/pb/converters"
	e "github.com/hudyweas/panshee/services/auth_service/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) RefreshAccessToken(ctx context.Context, req *pb.RefreshAccessTokenRequest) (*pb.RefreshAccessTokenResponse, error) {
	refreshTokenPayload, err := s.refreshTokenGenerator.VerifyToken(req.RefreshToken.GetValue())
	if err != nil{
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	fmt.Println(refreshTokenPayload)

	if refreshTokenPayload.UserID.String() != req.GetUserId() {
		return nil, status.Errorf(codes.PermissionDenied, "token owns to another user")
	}

	session, err := s.db.SessionSelectByID(refreshTokenPayload.ID)
	if err != nil {
		if err.Error() == e.ErrNoDataInDatabase {
			return nil, status.Errorf(codes.Unauthenticated, "no session for this token: %s", refreshTokenPayload.ID)
		}
		return nil, err
	}

	if session.IsBlocked {
		return nil, status.Errorf(codes.Unauthenticated, "session is blocked")
	}

	accessToken, accessTokenPayload, err := s.accessTokenGenerator.CreateToken(refreshTokenPayload.UserID, s.config.ACCESS_TOKEN_DURATION)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "creating access token failed: %s", err)
	}

	res := &pb.RefreshAccessTokenResponse{
		AccessToken: converters.PbTokenFromValueAndPayload(accessToken, accessTokenPayload),
	}

	return res, nil
}

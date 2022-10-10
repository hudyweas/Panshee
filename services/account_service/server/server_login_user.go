package server

import (
	"context"

	"github.com/hudyweas/panshee/services/account_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/account_service/database"
	e "github.com/hudyweas/panshee/services/account_service/errors"
	"github.com/hudyweas/panshee/services/account_service/util"
	val "github.com/hudyweas/panshee/services/account_service/validators"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	md := util.GetMetadatas(ctx)

	//validating user email
	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		return nil, status.Errorf(codes.NotFound, "invalid email")
	}

	//getting user password from database
	user, err := s.db.UserSelectByEmail(req.GetEmail())
	if err != nil {
		if err.Error() == e.ErrNoDataInDatabase {
			return nil, status.Errorf(codes.NotFound, "no user with this email: %s in database", req.GetEmail())
		}

		return nil, err
	}

	//checking password validation
	if err := val.ValidateHash(user.Password, req.GetPassword()); err != nil {
		return nil, status.Errorf(codes.NotFound, "invalid password")
	}

	//get access token
	accessToken, accessTokenPayload, err := s.tokenGenerator.CreateToken(req.GetEmail(), s.config.ACCESS_TOKEN_DURATION)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "creating access token failed: %s", err)
	}

	//get refresh token for current session
	refreshToken, refreshTokenPayload, err := s.tokenGenerator.CreateToken(req.GetEmail(), s.config.REFRESH_TOKEN_DURATION)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "creating refresh token failed: %s", err)
	}

	//adding new session to database
	session, err := s.db.SessionCreate(database.Session{
		SessionID:    refreshTokenPayload.ID,
		Email:        user.Email,
		RefreshToken: refreshToken,
		ClientIp:     md.ClientIp,
		UserAgent:    md.UserAgent,
		IsBlocked:    false,
		ExpiresAt:    refreshTokenPayload.ExpiredAt,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "creating session failed: %s", err)
	}

	//preparing response
	rsp := &pb.LoginUserResponse{
		SessionId:             session.SessionID.String(),
		AccessToken:           accessToken,
		AccessTokenExpirationTime:  timestamppb.New(accessTokenPayload.ExpiredAt),
		RefreshToken:          refreshToken,
		RefreshTokenExpirationTime: timestamppb.New(refreshTokenPayload.ExpiredAt),
		User:                  convertUser(*user),
	}

	return rsp, nil
}

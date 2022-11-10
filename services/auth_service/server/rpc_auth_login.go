package server

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hudyweas/panshee/services/auth_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/auth_service/api/panshee/v1/pb/converters"
	"github.com/hudyweas/panshee/services/auth_service/database"
	e "github.com/hudyweas/panshee/services/auth_service/errors"
	val "github.com/hudyweas/panshee/services/auth_service/internal/validators"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AuthLogin(ctx context.Context, req *pb.AuthLoginRequest) (*pb.AuthLoginResponse, error) {
	if validationErrors := validateAuthLoginRequest(req); len(validationErrors) > 0{
		return nil, status.Errorf(codes.InvalidArgument, validationErrors)
	}

	userId := uuid.MustParse(req.GetUserId())

	//getting password from database
	userPasswordFromDB, err := s.db.UserPasswordSelectByUserID(userId)
	if err != nil {
		if err.Error() == e.ErrNoDataInDatabase {
			return nil, status.Errorf(codes.NotFound, "no password in database for this user: %s", userId)
		}
		return nil, err
	}

	fmt.Print(userPasswordFromDB.Password)

	//checking password validation with hash from database
	if err := val.ValidateHash(userPasswordFromDB.Password, req.GetPassword()); err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "incorrect password")
	}

	//creating access token
	accessToken, accessTokenPayload, err := s.accessTokenGenerator.CreateToken(userId, s.config.ACCESS_TOKEN_DURATION)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "creating access token failed: %s", err)
	}

	//creating refresh token
	refreshToken, refreshTokenPayload, err := s.refreshTokenGenerator.CreateToken(userId, s.config.REFRESH_TOKEN_DURATION)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "creating refresh token failed: %s", err)
	}

	//creating session
	session, err := s.db.SessionCreate(database.Session{
		ID:           refreshTokenPayload.ID,
		UserID:       userId,
		RefreshToken: refreshToken,
		ExpiresAt:    refreshTokenPayload.ExpiredAt,
		ClientIp:     req.GetClientIp(),
		UserAgent:    req.GetUserAgent(),
		IsBlocked:    false,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "creating session failed: %s", err)
	}

	//creating response
	res := &pb.AuthLoginResponse{
		SessionId:    session.ID.String(),
		AccessToken:  converters.PbTokenFromValueAndPayload(accessToken, accessTokenPayload),
		RefreshToken:  converters.PbTokenFromValueAndPayload(refreshToken, refreshTokenPayload),
	}

	return res, nil
}

func validateAuthLoginRequest(req *pb.AuthLoginRequest) (errors string){
	if _ ,err := uuid.Parse(req.GetUserId()); err != nil{
		errors += err.Error() + "\n"
	}

	if err := val.ValidatePasswordLength(req.GetPassword()); err != nil{
		errors += err.Error() + "\n"
	}

	return
}

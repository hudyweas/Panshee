package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hudyweas/panshee/services/api_service/api/panshee/v1/pb"
	e "github.com/hudyweas/panshee/services/api_service/errors"

	"github.com/hudyweas/panshee/services/api_service/internal/contextData"
	val "github.com/hudyweas/panshee/services/api_service/internal/validators"
	authPb "github.com/hudyweas/panshee/services/api_service/services/auth_service/pb"
)

func (s *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	user, err := s.db.UserSelectByEmail(req.GetEmail())
	if err != nil {
		if err.Error() == e.ErrNoDataInDatabase {
			return nil, status.Errorf(codes.NotFound, "no user with this email: %s in database", req.GetEmail())
		}

		return nil, err
	}

	loginResponse, err := s.services.AuthService.AuthLogin(context.Background(), &authPb.AuthLoginRequest{
		UserId:    user.ID.String(),
		Password:  req.GetPassword(),
		UserAgent: contextData.GetUserAgent(ctx),
		ClientIp:  contextData.GetClientIP(ctx),
	})
	if err != nil{
		return nil, status.Errorf(codes.Unauthenticated, "failed to get authentication: %s", err)
	}

	rsp := &pb.LoginUserResponse{
		SessionId:    loginResponse.SessionId,
		AccessToken:  &pb.TokenDTO{
			Value:          loginResponse.AccessToken.Value,
			ExpirationTime: loginResponse.AccessToken.GetExpirationTime(),
		},
		RefreshToken: &pb.TokenDTO{
			Value:          loginResponse.RefreshToken.Value,
			ExpirationTime: loginResponse.RefreshToken.GetExpirationTime(),
		},
		User:         &pb.User{},
	}

	return rsp, nil
}

func validateLoginUserRequest(req *pb.LoginUserRequest) (errors string){
	//validating email from request
	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		errors += "invalid email" + "\n"
	}

	//validating password from request
	if err := val.ValidatePassword(req.GetPassword()); err != nil{
		errors += err.Error() + "\n"
	}

	return
}

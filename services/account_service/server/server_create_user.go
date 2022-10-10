package server

import (
	"context"

	"github.com/hudyweas/panshee/services/account_service/api/panshee/v1/pb"
	e "github.com/hudyweas/panshee/services/account_service/errors"
	"github.com/hudyweas/panshee/services/account_service/util"
	val "github.com/hudyweas/panshee/services/account_service/validators"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	//checking if email is valid
	if err := val.ValidateEmail(req.User.GetEmail()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid email", err)
	}

	//checking if user is already in database
	_, err := s.db.UserSelectByEmail(req.User.GetEmail())
	if err != nil {
		if err.Error() == e.ErrNoDataInDatabase {
			return nil, status.Errorf(codes.AlreadyExists, "user with this email: %s, already exists", req.User.GetEmail())
		}

		return nil, err
	}

	//preparing password
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	//sending data to database
	createdUser, err := s.db.UserCreate(req.User.GetEmail(), hashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user %s:", err)
	}

	//creating response
	rsp := convertUser(*createdUser)

	return rsp, nil
}

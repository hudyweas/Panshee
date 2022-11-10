package server

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hudyweas/panshee/services/auth_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/auth_service/database"
	"github.com/hudyweas/panshee/services/auth_service/internal/util"
	val "github.com/hudyweas/panshee/services/auth_service/internal/validators"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUserPassword(ctx context.Context, req *pb.CreateUserPasswordRequest) (*pb.Empty, error) {
	if validationErrors := validateCreateUserPasswordRequest(req); len(validationErrors) > 0{
		return nil, status.Errorf(codes.InvalidArgument, validationErrors)
	}

	//checking if user password is already in database
	if is, err := s.db.IsUserPasswordInDatabaseByUserID(uuid.MustParse(req.GetUserId())); is && err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "user with id: '%s' , already exists in database", req.GetUserId())
	}else if err != nil{
		return nil, err
	}

	fmt.Print(req.GetPassword())

	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	_, err = s.db.UserPasswordCreate(database.UserPassword{
		ID:       uuid.New(),
		UserID:   uuid.MustParse(req.GetUserId()),
		Password: hashedPassword,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user %s:", err)
	}

	return &pb.Empty{}, nil
}

func validateCreateUserPasswordRequest(req *pb.CreateUserPasswordRequest) (errors string){
	//validating password from request
	if err := val.ValidatePasswordLength(req.GetPassword()); err != nil{
		errors += err.Error() + "\n"
	}

	if _ ,err := uuid.Parse(req.GetUserId()); err != nil{
		errors += err.Error() + "\n"
	}

	return
}

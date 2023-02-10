package server

import (
	"context"

	authPb "github.com/hudyweas/panshee/services/api_service/services/auth_service/pb"
	emailPb "github.com/hudyweas/panshee/services/api_service/services/email_service/pb"

	"github.com/hudyweas/panshee/services/api_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/api_service/api/panshee/v1/pb/converters"

	val "github.com/hudyweas/panshee/services/api_service/internal/validators"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	//checking request validation
	if validationErrors := validateCreateUserRequest(req); len(validationErrors) > 0{
		return nil, status.Errorf(codes.InvalidArgument, validationErrors)
	}

	//checking if user is already in database
	_, err := s.db.UserSelectByEmail(req.User.GetEmail())
	if err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "user with email: '%s' , already exists", req.User.GetEmail())
	}

	//sending data to database
	createdUser, err := s.db.UserCreate(req.User.GetEmail())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user %s:", err)
	}

	//sending password to auth service
	_, err = s.services.AuthService.CreateUserPassword(context.Background(), &authPb.CreateUserPasswordRequest{
		UserId:   createdUser.ID.String(),
		Password: req.Password,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user password %s:", err)
	}

	//sending email
	_, err = s.services.EmailService.SendEmail(context.Background(), &emailPb.SendEmailRequest{
		Email: &emailPb.Email{
			To:      createdUser.Email,
			Subject: "Zalozenie konta",
			Message: "Brawo",
		},
	})
	if err != nil {
		return nil, err
	}

	//creating response
	rsp := converters.ConvertDbUserToPbUser(*createdUser)

	return rsp, nil
}

func validateCreateUserRequest(req *pb.CreateUserRequest) (errors string){
	//validating email from request
	if err := val.ValidateEmail(req.User.GetEmail()); err != nil {
		errors += "invalid email" + "\n"
	}

	//validating password from request
	if err := val.ValidatePassword(req.GetPassword()); err != nil{
		errors += err.Error() + "\n"
	}

	return
}

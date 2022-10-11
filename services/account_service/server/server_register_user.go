package server

import (
	"context"
	"fmt"

	"github.com/hudyweas/panshee/services/account_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/account_service/util"
	val "github.com/hudyweas/panshee/services/account_service/validators"
	emailPb "github.com/hudyweas/panshee/services/email_service/api/panshee/v1/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.User, error) {
	//checking request validation
	if validationErrors := validateCreateUserRequest(req); len(validationErrors) > 0{
		return nil, status.Errorf(codes.InvalidArgument, validationErrors)
	}

	//checking if user is already in database
	_, err := s.db.UserSelectByEmail(req.GetEmail())
	if err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "user with email: '%s' , already exists", req.GetEmail())
	}

	//preparing password
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	//sending data to database
	createdUser, err := s.db.UserCreate(req.GetEmail(), hashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user %s:", err)
	}

	//sending email
	_, err = s.services.EmailService.SendEmail(context.Background(), &emailPb.SendEmailRequest{
		Email: &emailPb.Email{
			To:      "hudyweas@gmail.com",
			Subject: "Zalozenie konta",
			Message: "Brawo",
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	//creating response
	rsp := convertUser(*createdUser)

	return rsp, nil
}

func validateCreateUserRequest(req *pb.RegisterUserRequest) (errors string){
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

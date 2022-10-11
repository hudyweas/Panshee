package server

import (
	"context"
	"strconv"

	"github.com/hudyweas/panshee/services/account_service/api/panshee/v1/pb"
	e "github.com/hudyweas/panshee/services/account_service/errors"
	"github.com/hudyweas/panshee/services/account_service/token"
	"github.com/hudyweas/panshee/services/account_service/util"
	val "github.com/hudyweas/panshee/services/account_service/validators"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error){
	//validating request data
	if validationErrors := validateChangePasswordRequest(req); len(validationErrors) > 0{
		return nil, status.Errorf(codes.InvalidArgument, validationErrors)
	}

	//getting user info from context
	payload := ctx.Value("accessPayload").(*token.Payload)

	//checking if user wants to check his own password
	user, err := s.db.UserSelectByEmail(payload.Email)
	if err != nil {
		if err.Error() == e.ErrNoDataInDatabase {
			return nil, status.Errorf(codes.NotFound, "no user with this email: %s in database", req.GetId())
		}

		return nil, err
	}

	if strconv.FormatInt(int64(user.ID), 10) != req.GetId() {
		return nil, status.Errorf(codes.Unauthenticated, "you can't change other users password")
	}

	//preparing new password to update
	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error during hashing password: %s", err)
	}

	//updating password
	user, err = s.db.UserPasswordUpdateByID(user.ID, hashedPassword)
	if err != nil{
		return nil, status.Errorf(codes.Internal, "coudn't update password: %s", err)
	}

	//creating response
	res := &pb.ChangePasswordResponse{
		User: convertUser(*user),
	}

	return res, nil
}

func validateChangePasswordRequest(req *pb.ChangePasswordRequest) (errors string){
	//validating password from request
	if err := val.ValidatePassword(req.GetPassword()); err != nil{
		errors += err.Error() + "\n"
	}

	return
}

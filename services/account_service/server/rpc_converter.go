package server

import (
	"github.com/hudyweas/panshee/services/account_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/account_service/database"
)

func convertUser(user database.User) *pb.User {
	return &pb.User{
		Id:     int32(user.ID),
		Email:  user.Email,
		Active: user.Active,
	}
}

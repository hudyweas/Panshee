package server

import (
	"github.com/hudyweas/panshee/services/api_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/api_service/database"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertDbUserToPbUser(user database.User) *pb.User {
	return &pb.User{
		Id:         user.ID.String(),
		Email:      user.Email,
		Active:     user.Active,
		CreateTime: timestamppb.New(user.CreatedAt),
		UpdateTime: timestamppb.New(user.UpdatedAt),
		DeleteTime: timestamppb.New(user.DeletedAt),
	}
}

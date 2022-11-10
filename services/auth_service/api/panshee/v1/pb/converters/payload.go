package converters

import (
	"github.com/hudyweas/panshee/services/auth_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/auth_service/token"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func PbPayloadFromTokenPayload(payload *token.Payload) *pb.Payload{
	return &pb.Payload{
		ID:        payload.ID.String(),
		UserID:    payload.UserID.String(),
		CreatedAt: timestamppb.New(payload.CreatedAt),
		ExpiredAt: timestamppb.New(payload.ExpiredAt),
	}
}

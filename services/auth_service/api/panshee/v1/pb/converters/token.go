package converters

import (
	"github.com/hudyweas/panshee/services/auth_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/auth_service/token"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func PbTokenFromValueAndPayload(value string, payload *token.Payload) *pb.Token{
	return &pb.Token{
		Value:          value,
		ExpirationTime: timestamppb.New(payload.ExpiredAt),
	}
}

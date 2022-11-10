package contextData

import (
	"context"

	"github.com/hudyweas/panshee/services/api_service/services/auth_service/pb"
)

type Metadata struct {
	ClientIp  string
	UserAgent string
}

func GetClientIP(ctx context.Context) string {
	return ctx.Value("clientIP").(string)
}

func SetClientIP(ctx context.Context, clientIP string) context.Context {
	return context.WithValue(ctx, "clientIP", clientIP)
}

func GetUserAgent(ctx context.Context) string {
	return ctx.Value("userAgent").(string)
}

func SetUserAgent(ctx context.Context, userAgent string) context.Context {
	return context.WithValue(ctx, "userAgent", userAgent)
}

func GetAccessPayload(ctx context.Context) *pb.Payload{
	return ctx.Value("accessPayload").(*pb.Payload)
}

func SetAccessPayload(ctx context.Context, payload *pb.Payload) context.Context {
	return context.WithValue(ctx, "accessPayload", payload)
}

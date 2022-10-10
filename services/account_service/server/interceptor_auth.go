package server

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

//authHeaderScheme
//auth: bearer <token>

func (s *Server) AuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		//checking if path needs to auth
		path := strings.Replace(info.FullMethod, "/pb.PansheeUserRestService/", "", 1)
		if path == "LoginUser"{
			return handler(ctx, req)
		}

		//getting header
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, fmt.Errorf("missing metadata")
		}

		authHeader := md.Get("auth")

		//checking if header is not empty
		if len(authHeader) == 0 {
			return nil, fmt.Errorf("no auth header")
		}

		if len(authHeader) < 2 {
			return nil, fmt.Errorf("wrong auth header format")
		}

		authType := authHeader[0]
		token := authHeader[1]

		if authType != "bearer" {
			return nil, fmt.Errorf("invalid auth type")
		}

		payload, err := s.tokenGenerator.VerifyToken(token)
		if err != nil {
			return nil, fmt.Errorf("invalid token")
		}

		ctx = context.WithValue(ctx, "accessPayload", payload)

		return handler(ctx, req)
	}
}

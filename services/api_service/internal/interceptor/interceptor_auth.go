package interceptor

import (
	"context"
	"fmt"
	"strings"

	"github.com/hudyweas/panshee/services/api_service/internal/contextData"
	"github.com/hudyweas/panshee/services/api_service/services"
	"github.com/hudyweas/panshee/services/api_service/services/auth_service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

//authHeaderScheme
//auth: bearer <token>

func AuthUnaryServerInterceptor(services *services.Services) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		//checking if path needs to auth
		path := strings.Replace(info.FullMethod, "/api.panshee.v1.proto.ApiService/", "", 1)
		if path == "LoginUser" ||
			path == "RefreshAccessToken" ||
			path == "CreateUser" {
			return handler(ctx, req)
		}

		//getting header
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, fmt.Errorf("missing metadata")
		}

		authorization := md.Get("authorization")

		if(len(authorization) == 0 ){
			return nil, fmt.Errorf("no auth header")
		}

		authHeader := strings.Fields(md.Get("authorization")[0])

		//checking if header is not empty
		if len(authHeader) == 0 {
			return nil, fmt.Errorf("no auth header")
		}

		if len(authHeader) < 2 {
			return nil, fmt.Errorf("invalid auth header format")
		}

		authType := authHeader[0]
		tokenString := authHeader[1]

		if authType != "bearer" {
			return nil, fmt.Errorf("invalid auth type")
		}

		validateAccessTokenResponse, err := services.AuthService.ValidateAccessToken(context.Background(), &pb.ValidateAccessTokenRequest{
			AccessToken: &pb.Token{
				Value: tokenString,
			},
		})
		if err != nil {
			return nil, err
		}

		payload := validateAccessTokenResponse.GetAccessTokenPayload()

		ctx = contextData.SetAccessPayload(ctx, payload)

		return handler(ctx, req)
	}
}

package server

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	//http Request
	grpcGatewayUserAgentHeader = "grpcfateway-user-agent"
	xForwardedForHeader        = "x-forwarded-for"

	//grpc request
	userAgentHeader = "user-agent"
)

func (s *Server) MetadataUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		//from grpc request
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			//getting user agent
			if userAgents := md.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
				ctx = context.WithValue(ctx, "userAgent", userAgents[0])
			}

			//getting client ip
			if clientIPs := md.Get(xForwardedForHeader); len(clientIPs) > 0 {
				ctx = context.WithValue(ctx, "clientIP", clientIPs[0])
			}
		}

		//From http request
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			//getting user agent
			if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
				ctx = context.WithValue(ctx, "userAgent", userAgents[0])
			}
		}

		//getting client ip
		if p, ok := peer.FromContext(ctx); ok {
			ctx = context.WithValue(ctx, "clientIP", p.Addr.String())
		}

		h, err := handler(ctx, req)
		return h, err
	}
}

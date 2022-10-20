package server

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GrpcLoggerInterceptor() grpc.UnaryServerInterceptor{
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	startTime := time.Now()
	result, err := handler(ctx, req)
	duration := time.Since(startTime)

	statusCode := codes.Unknown
	errorMessage := ""
	if st, ok := status.FromError(err); ok {
		statusCode = st.Code()
		errorMessage = st.Message()
	}

	fields := log.Fields{
		"method": info.FullMethod,
		"status_code": int(statusCode),
		"status_text": statusCode.String(),
		"duration": duration,
	}

	if err != nil {
		fields["error"] = errorMessage
		s.log.WithFields(fields).Error("Request unsuccesful")
		return result, err
	}

	s.log.WithFields(fields).Info("Request succesful")

	return result, err
	}
}

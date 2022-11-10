package interceptor

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcLoggerInterceptor(log *logrus.Logger) grpc.UnaryServerInterceptor{
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	result, err := handler(ctx, req)

	statusCode := codes.Unknown
	errorMessage := ""
	if st, ok := status.FromError(err); ok {
		statusCode = st.Code()
		errorMessage = st.Message()
	}

	fields := logrus.Fields{
		"method": info.FullMethod,
		"status_code": int(statusCode),
		"status_text": statusCode.String(),
	}

	if err != nil {
		fields["error"] = errorMessage
		log.WithFields(fields).Error("Request unsuccesful")
		return result, err
	}

	log.WithFields(fields).Info("Request succesful")

	return result, err
	}
}

package services

import (
	"github.com/hudyweas/panshee/services/api_service/config"
	authPb "github.com/hudyweas/panshee/services/api_service/services/auth_service/pb"
	emailPb "github.com/hudyweas/panshee/services/api_service/services/email_service/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Services struct{
	EmailService emailPb.PansheeEmailServiceClient
	AuthService authPb.AuthServiceClient
}

func ConnectServices(config config.Config) (Services, error){
	emailService, err := connectEmailService(config.EMAIL_SERVICE_GRPC_ADDRESS)
	if err != nil {
		return Services{}, err
	}
	authService, err := connectAuthService(config.AUTH_SERVICE_GRPC_ADDRESS)
	if err != nil {
		return Services{}, err
	}

	services := &Services{
		EmailService: emailService,
		AuthService:  authService,
	}

	return *services, nil
}

func connectEmailService(address string) (emailPb.PansheeEmailServiceClient, error) {
	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(address, dialOptions...)
	if err != nil {
		return nil, err
	}

	emailService := emailPb.NewPansheeEmailServiceClient(conn)

	return emailService, nil
}

func connectAuthService(address string) (authPb.AuthServiceClient, error) {
	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(address, dialOptions...)
	if err != nil {
		return nil, err
	}

	authService := authPb.NewAuthServiceClient(conn)

	return authService, nil
}

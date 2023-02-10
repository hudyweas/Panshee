package services

import (
	"github.com/hudyweas/panshee/services/api_service/config"
	authPb "github.com/hudyweas/panshee/services/api_service/services/auth_service/pb"
	emailPb "github.com/hudyweas/panshee/services/api_service/services/email_service/pb"
	walletPb "github.com/hudyweas/panshee/services/api_service/services/wallet_service/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Services struct{
	EmailService emailPb.PansheeEmailServiceClient
	AuthService authPb.AuthServiceClient
	WalletService walletPb.PansheeWalletServiceClient
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
	walletService, err := connectWalletService(config.WALLET_SERVICE_GRPC_ADDRESS)
	if err != nil {
		return Services{}, err
	}

	services := &Services{
		EmailService: emailService,
		AuthService:  authService,
		WalletService:  walletService,
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

func connectWalletService(address string) (walletPb.PansheeWalletServiceClient, error) {
	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(address, dialOptions...)
	if err != nil {
		return nil, err
	}

	walletService := walletPb.NewPansheeWalletServiceClient(conn)

	return walletService, nil
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

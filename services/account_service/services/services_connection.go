package services

import (
	"github.com/hudyweas/panshee/services/account_service/config"
	"github.com/hudyweas/panshee/services/account_service/imported/email_service/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Services struct{
	EmailService pb.PansheeEmailServiceClient
}

func ConnectServices(config config.Config) (Services, error){
	emailService, err := connectEmailService(config.EMAIL_SERVICE_GRPC_ADDRESS)
	if err != nil {
		return Services{}, err
	}

	services := &Services{
		EmailService: emailService,
	}

	return *services, nil
}

func connectEmailService(address string) (pb.PansheeEmailServiceClient, error) {
	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(address, dialOptions...)
	if err != nil {
		return nil, err
	}

	emailService := pb.NewPansheeEmailServiceClient(conn)

	return emailService, nil
}

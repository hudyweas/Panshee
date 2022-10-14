package server

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
	emailService, err := connectEmailService()
	if err != nil {
		return Services{}, err
	}

	services := &Services{
		EmailService: emailService,
	}

	return *services, nil
}

func connectEmailService() (pb.PansheeEmailServiceClient, error) {
	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial("0.0.0.0:50052", dialOptions...)
	if err != nil {
		return nil, err
	}

	emailService := pb.NewPansheeEmailServiceClient(conn)

	return emailService, nil
}

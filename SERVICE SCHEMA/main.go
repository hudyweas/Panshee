package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/sirupsen/logrus"

	"github.com/hudyweas/panshee/services/auth_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/auth_service/config"
	"github.com/hudyweas/panshee/services/auth_service/server"
)

var log = logrus.New()

func main() {
	errChan := make(chan error, 1)

	log.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	//shutting down microservis
	go func() {
		err := <-errChan

		defer func() {
			close(errChan)
		}()

		log.Fatal("Closing microservice %s: ", err.Error())
	}()

	config, err := config.LoadConfigFromFile(".env")
	if err != nil {
		errChan <- fmt.Errorf("cannot load config: %s", err)
	}

	runGrpcServer(config, errChan)
}

func runGrpcServer(config config.Config, errors chan error) {
	server := server.NewServer(config, log)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptors.GrpcLoggerInterceptor(log)))

	pb.RegisterAuthServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPC_ADDRESS)
	if err != nil {
		errors <- fmt.Errorf("cannot create listener: %s", err)
	}

	log.Info("started gRPC server at ", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		errors <- fmt.Errorf("cannot create gRPC server: %s", err)
	}
}

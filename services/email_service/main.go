package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/hudyweas/panshee/services/email_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/email_service/config"
	"github.com/hudyweas/panshee/services/email_service/server"
)

func main() {
	errChan := make(chan error, 1)

	//shutting down microservis
	go func() {
		err := <-errChan

		defer func() {
			close(errChan)
		}()

		log.Fatal(err)
	}()

	config, err := config.LoadConfigFromFile("../../.env")
	if err != nil {
		errChan <- fmt.Errorf("cannot load config: %s", err)
	}

	runGrpcServer(config, errChan)
}

func runGrpcServer(config config.Config, errors chan error) {
	server := server.NewServer(config)

	grpcServer := grpc.NewServer()

	pb.RegisterPansheeEmailServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPC_ADDRESS)
	if err != nil {
		errors <- fmt.Errorf("cannot create listener: %s", err)
	}

	log.Print("started gRPC server at ", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		errors <- fmt.Errorf("cannot create gRPC server: %s", err)
	}
}

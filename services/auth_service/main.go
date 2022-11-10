package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"

	"github.com/hudyweas/panshee/services/auth_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/auth_service/config"
	"github.com/hudyweas/panshee/services/auth_service/database"
	"github.com/hudyweas/panshee/services/auth_service/internal/interceptor"
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

	db, err := database.Connect(pg.Options{
		Addr:     config.DB_HOST + ":" + config.DB_PORT,
		User:     config.DB_USER,
		Password: config.DB_PASSWORD,
		Database: config.DB_DBNAME,
	})
	if err != nil{
		errChan <- err
	}

	if err := db.CheckConnection(); err != nil{
		errChan <- err
	}

	log.Info("Connected to database")

	runGrpcServer(config, db, errChan)
}

func runGrpcServer(config config.Config, db database.Database, errors chan error) {
	server := server.NewServer(config, db, log)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptor.GrpcLoggerInterceptor(log)))

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

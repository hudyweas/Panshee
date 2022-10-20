package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/go-pg/pg/v10"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hudyweas/panshee/services/account_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/account_service/config"
	"github.com/hudyweas/panshee/services/account_service/database"
	"github.com/hudyweas/panshee/services/account_service/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

var(
	//go:embed doc/swagger/*
	res embed.FS
)

var log = logrus.New()

func main() {
	errChan := make(chan error, 1)

	log.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	//shutting down microservis
	go func(){
		err := <- errChan

		defer func(){
			close(errChan)
		}()

		log.Fatal("Closing microservice %s: ", err.Error())
	}()

	config, err := config.LoadConfigFromFile(".env", log)
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

	go runGrpcServer(config, db, errChan)
	runGatewayServer(config, db, errChan)
}

func runGrpcServer(config config.Config, db database.Database, errors chan error){
	server := server.NewServer(config, db, log)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			server.GrpcLoggerInterceptor(),
			server.AuthUnaryServerInterceptor(),
			server.MetadataUnaryServerInterceptor(),
		)),
	)

	pb.RegisterPansheeAccountServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPC_ADDRESS)
	if err != nil{
		errors <- fmt.Errorf("cannot create listener: %s", err)
	}

	log.Infof("started gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil{
		errors <- fmt.Errorf("cannot create gRPC server: %s", err)
	}
}

func runGatewayServer(config config.Config, db database.Database, errors chan error){
	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)
	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := pb.RegisterPansheeAccountServiceHandlerFromEndpoint(c, grpcMux, config.GRPC_ADDRESS, dialOptions); err != nil {
		errors <- fmt.Errorf("cannot register handler server: %s", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	//creating documentation site
	fs, err := fs.Sub(res, "doc/swagger")
    if err != nil {
        log.Error("Cannot create documentation site")
    }
	fileServer := http.FileServer(http.FS(fs))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fileServer))

	//creating listener to host http gateway server
	listener, err := net.Listen("tcp", config.ADDRESS)
	if err != nil{
		errors <- fmt.Errorf("cannot create lisener: %s", err)
	}

	log.Infof("started HTTTP gateway server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil{
		errors <- fmt.Errorf("cannot create HTTTP gateway server: %s", err)
	}
}

package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net"
	"net/http"

	"github.com/rs/cors"

	"github.com/sirupsen/logrus"

	"github.com/hudyweas/panshee/services/api_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/api_service/config"
	"github.com/hudyweas/panshee/services/api_service/database"
	"github.com/hudyweas/panshee/services/api_service/internal/interceptor"
	"github.com/hudyweas/panshee/services/api_service/server"

	"github.com/go-pg/pg/v10"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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

	fmt.Print(config)

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

	go runGrpcServer(config, db, errChan)
	runGatewayServer(config, db, errChan)
}

func runGrpcServer(config config.Config, db database.Database, errors chan error) {
	server := server.NewServer(config, db, log)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			interceptor.GrpcLoggerInterceptor(log),
			interceptor.AuthUnaryServerInterceptor(server.GetServices()),
			interceptor.MetadataUnaryServerInterceptor(),
		)),
	)

	pb.RegisterApiServiceServer(grpcServer, server)
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

	if err := pb.RegisterApiServiceHandlerFromEndpoint(c, grpcMux, config.GRPC_ADDRESS, dialOptions); err != nil {
		errors <- fmt.Errorf("cannot register handler server: %s", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	 co := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	})

	handler := co.Handler(mux)

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
	err = http.Serve(listener, handler)
	if err != nil{
		errors <- fmt.Errorf("cannot create HTTTP gateway server: %s", err)
	}
}

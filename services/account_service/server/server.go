package server

import (
	"github.com/hudyweas/panshee/services/account_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/account_service/config"
	"github.com/hudyweas/panshee/services/account_service/database"
	"github.com/hudyweas/panshee/services/account_service/services"
	"github.com/hudyweas/panshee/services/account_service/token"
	"github.com/sirupsen/logrus"
)

type Server struct {
	pb.UnimplementedPansheeAccountServiceServer

	db             database.Database
	config         config.Config
	tokenGenerator token.TokenGenerator
	log            *logrus.Logger

	services services.Services
}

func NewServer(config config.Config, db database.Database, log *logrus.Logger) (s *Server) {

	tokenGenerator, err := token.NewTokenGenerator(config.TOKEN_KEY)
	if err != nil {
		log.Panic(err)
	}

	services, err := services.ConnectServices(config)
	if err != nil {
		log.Panic(err)
	}

	log.Info("connected to Email Service")

	s = &Server{
		db:             db,
		config:         config,
		tokenGenerator: tokenGenerator,
		log: log,

		services: services,
	}

	s.log.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})


	return
}

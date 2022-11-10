package server

import (
	"github.com/hudyweas/panshee/services/api_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/api_service/config"
	"github.com/hudyweas/panshee/services/api_service/database"
	"github.com/hudyweas/panshee/services/api_service/services"
	"github.com/sirupsen/logrus"
)

type Server struct {
	pb.UnimplementedApiServiceServer

	db             database.Database
	config         config.Config
	log            *logrus.Logger

	services services.Services
}

func NewServer(config config.Config, db database.Database, log *logrus.Logger) (s *Server) {

	services, err := services.ConnectServices(config)
	if err != nil {
		log.Panic(err)
	}

	log.Info("connected to Email Service")

	s = &Server{
		db:             db,
		config:         config,
		log: log,

		services: services,
	}

	s.log.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})


	return
}

func (s *Server) GetServices() *services.Services{
	return &s.services
}

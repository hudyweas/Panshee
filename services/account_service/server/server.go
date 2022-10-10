package server

import (
	"github.com/hudyweas/panshee/services/account_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/account_service/config"
	"github.com/hudyweas/panshee/services/account_service/database"
	"github.com/hudyweas/panshee/services/account_service/token"
)

type Server struct {
	pb.UnimplementedPansheeUserRestServiceServer

	db             database.Database
	config         config.Config
	tokenGenerator token.TokenGenerator

	services Services
}

func NewServer(config config.Config, db database.Database) (s *Server) {
	tokenGenerator, err := token.NewTokenGenerator(config.TOKEN_KEY)
	if err != nil {
		panic(err)
	}

	services, err := ConnectServices(config)
	if err != nil {
		panic(err)
	}

	s = &Server{
		db:             db,
		config:         config,
		tokenGenerator: tokenGenerator,

		services: services,
	}

	return
}

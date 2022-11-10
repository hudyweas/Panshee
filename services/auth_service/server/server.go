package server

import (
	"github.com/hudyweas/panshee/services/auth_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/auth_service/config"
	"github.com/hudyweas/panshee/services/auth_service/database"
	"github.com/hudyweas/panshee/services/auth_service/token"
	"github.com/sirupsen/logrus"
)

type Server struct {
	pb.UnimplementedAuthServiceServer

	db             database.Database
	config config.Config
	log *logrus.Logger

	accessTokenGenerator token.TokenGenerator
	refreshTokenGenerator token.TokenGenerator
}

func NewServer(config config.Config, db database.Database, log *logrus.Logger) (s *Server) {
	accessTokenGenerator, err := token.NewTokenGenerator(config.TOKEN_KEY)
	if err != nil {
		log.Panic(err)
	}

	refreshTokenGenerator, err := token.NewTokenGenerator(config.REFRESH_TOKEN_KEY)
	if err != nil {
		log.Panic(err)
	}

	s = &Server{
		db: db,
		config: config,
		accessTokenGenerator: accessTokenGenerator,
		refreshTokenGenerator: refreshTokenGenerator,
		log: log,
	}

	return
}

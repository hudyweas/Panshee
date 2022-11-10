package server

import (
	"github.com/hudyweas/panshee/services/auth_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/auth_service/config"
	"github.com/sirupsen/logrus"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	config config.Config
	log *logrus.Logger
}

func NewServer(config config.Config, log *logrus.Logger) (s *Server) {
	s = &Server{
		config: config,
		log: log,
	}

	return
}

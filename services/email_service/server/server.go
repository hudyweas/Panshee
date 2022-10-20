package server

import (
	"github.com/hudyweas/panshee/services/email_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/email_service/config"
	"github.com/hudyweas/panshee/services/email_service/email"
	"github.com/sirupsen/logrus"
)

type Server struct {
	pb.UnimplementedPansheeEmailServiceServer
	config config.Config
	emailSender email.EmailSender
	log *logrus.Logger
}

func NewServer(config config.Config, log *logrus.Logger) (s *Server) {
	emailSender, err := email.NewEmailSender(config.EMAIL_ADDRESS, config.EMAIL_PASSWORD)
	if err != nil{
		s.log.Panic(err)
	}

	s = &Server{
		config: config,
		emailSender: emailSender,
		log: log,
	}

	s.log.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	return
}

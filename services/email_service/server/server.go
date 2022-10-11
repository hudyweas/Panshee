package server

import (
	"github.com/hudyweas/panshee/services/email_service/api/panshee/v1/pb"
	"github.com/hudyweas/panshee/services/email_service/config"
	"github.com/hudyweas/panshee/services/email_service/email"
)

type Server struct {
	pb.UnimplementedPansheeEmailServiceServer
	config config.Config
	emailSender email.EmailSender
}

func NewServer(config config.Config) (s *Server) {
	emailSender, err := email.NewEmailSender(config.EMAIL_ADDRESS, config.EMAIL_PASSWORD)
	if err != nil{
		panic(err)
	}

	s = &Server{
		config: config,
		emailSender: emailSender,
	}

	return
}

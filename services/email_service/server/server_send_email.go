package server

import (
	"context"
	"fmt"

	"github.com/hudyweas/panshee/services/email_service/api/panshee/v1/pb"
)

func (s *Server) SendEmail(ctx context.Context,req *pb.SendEmailRequest) (*pb.SendEmailResponse, error) {
	s.log.Info("Received request to send email")

	err := s.emailSender.Send(*req.GetEmail())
	if err != nil {
		return nil, fmt.Errorf("Cannot send email: %s", err)
	}
	s.log.Infof("Email send to: ", req.Email.GetTo())

	res := &pb.SendEmailResponse{
		Email: &pb.Email{},
	}

	s.log.Info("Sending response")
	return res, nil
}

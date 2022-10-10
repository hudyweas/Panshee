package server

import (
	"context"

	"github.com/hudyweas/panshee/services/email_service/api/panshee/v1/pb"
)

func (s *Server) SendEmail(ctx context.Context,req *pb.SendEmailRequest) (*pb.SendEmailResponse, error) {
	s.emailSender.Send(*req.GetEmail())

	res := &pb.SendEmailResponse{
		Email: &pb.Email{},
	}

	return res, nil
}

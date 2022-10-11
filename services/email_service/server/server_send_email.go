package server

import (
	"context"
	"fmt"

	"github.com/hudyweas/panshee/services/email_service/api/panshee/v1/pb"
)

func (s *Server) SendEmail(ctx context.Context,req *pb.SendEmailRequest) (*pb.SendEmailResponse, error) {
	err := s.emailSender.Send(*req.GetEmail())
	if err != nil {
		fmt.Println(err)
	}

	res := &pb.SendEmailResponse{
		Email: &pb.Email{},
	}

	return res, nil
}

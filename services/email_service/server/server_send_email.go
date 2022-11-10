package server

import (
	"context"
	"fmt"

	"github.com/hudyweas/panshee/services/email_service/api/panshee/v1/pb"
	val "github.com/hudyweas/panshee/services/email_service/internal/validators"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SendEmail(ctx context.Context,req *pb.SendEmailRequest) (*pb.SendEmailResponse, error) {
	if validationErrors := validateSendEmailRequest(req); len(validationErrors) > 0{
		return nil, status.Errorf(codes.InvalidArgument, validationErrors)
	}

	err := s.emailSender.Send(*req.GetEmail())
	if err != nil {
		return nil, fmt.Errorf("Cannot send email: %s", err)
	}

	res := &pb.SendEmailResponse{
		Email: req.GetEmail(),
	}
	return res, nil
}

func validateSendEmailRequest(req *pb.SendEmailRequest) (errors string){
	if err := val.ValidateEmail(req.Email.GetTo()); err != nil{
		errors += "invalid email of the receiver"
	}

	return
}

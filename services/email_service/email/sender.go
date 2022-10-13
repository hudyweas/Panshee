package email

import (
	"github.com/hudyweas/panshee/services/email_service/api/panshee/v1/pb"
	gomail "gopkg.in/gomail.v2"
)

type EmailDailer struct {
	Dialer *gomail.Dialer
}

type EmailSender interface{
	Send(email pb.Email) error
}

func NewEmailSender(from string, password string) (EmailSender, error){
	emailDialer := &EmailDailer{
		Dialer:   gomail.NewDialer("smtp.gmail.com", 587, from, password),
	}

	return emailDialer, nil
}

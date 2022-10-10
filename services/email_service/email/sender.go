package email

import (
	"crypto/tls"
	"net"
	"net/smtp"

	"github.com/hudyweas/panshee/services/email_service/api/panshee/v1/pb"
)

type EmailSenderData struct {
	from string
	password string

	config *tls.Config
	auth smtp.Auth

	server string
}

type EmailSender interface{
	Send(email pb.Email) error
}

func NewEmailSender(from string, password string, server string) (EmailSender, error){
	host, _, err := net.SplitHostPort(server)
	if err != nil {
		return nil, err
	}

	auth := smtp.PlainAuth("", from, password, host)

	config := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	emailSenderData := &EmailSenderData{
		from:     from,
		password: password,
		config:   config,
		auth:     auth,
		server:   server,
	}

	return emailSenderData, nil
}

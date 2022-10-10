package email

import (
	"net/smtp"

	"github.com/hudyweas/panshee/services/email_service/api/panshee/v1/pb"
)

func (sender *EmailSenderData) Send(email pb.Email) error {
	c, err := smtp.Dial(sender.server)
	if err != nil {
		return err
	}
	defer c.Close()

	err = c.StartTLS(sender.config)
	if err != nil {
		return err
	}

	err = c.Auth(sender.auth)
	if err != nil {
		return err
	}

	err = c.Mail(sender.from)
	if err != nil {
		return err
	}

	err = c.Rcpt(email.To)
	if err != nil {
		return err
	}

	wr, err := c.Data()
	if err != nil {
		return err
	}

	_, err = wr.Write([]byte(email.Message))
	if err != nil {
		return err
	}

	err = wr.Close()
	if err != nil {
		return err
	}

	return nil
}

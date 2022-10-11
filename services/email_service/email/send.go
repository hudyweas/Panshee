package email

import (
	"github.com/hudyweas/panshee/services/email_service/api/panshee/v1/pb"
	gomail "gopkg.in/gomail.v2"
)

func (d *EmailDailer) Send(email pb.Email) error {
	msg := gomail.NewMessage()
    msg.SetHeader("To", email.To)
    msg.SetHeader("Subject", email.Subject)
    msg.SetBody("text/html", email.Message)

	if err := d.Dialer.DialAndSend(); err != nil{
		return err
	}

	return nil
}

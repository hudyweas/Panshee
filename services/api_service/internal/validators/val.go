package val

import (
	"fmt"
	"net/mail"
)

func ValidateEmail(email string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return fmt.Errorf("invalid email")
	}
	return nil
}

func ValidatePassword(password string) error{
	if len(password) == 0 {
		return fmt.Errorf("no password provided")
	}

	if len(password) < 8 {
		return fmt.Errorf("password is too short")
	}

	return nil
}

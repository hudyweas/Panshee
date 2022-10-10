package val

import (
	"fmt"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

func ValidateEmail(email string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return fmt.Errorf("invalid email")
	}
	return nil
}

func ValidateHash(hashed, plain string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)); err != nil {
		return fmt.Errorf("invalid hash")
	}

	return nil
}

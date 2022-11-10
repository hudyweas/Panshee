package val

import (
	"fmt"
	"net/mail"
	"regexp"

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

func ValidatePasswordLength(password string) error{
	if len(password) == 0 {
		return fmt.Errorf("no password provided")
	}

	if len(password) < 8 {
		return fmt.Errorf("password is too short")
	}

	return nil
}

func ValidateIP(ip string) error {
	re := regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$`)

	if !re.MatchString(ip){
		return fmt.Errorf("Invalid ip address")
	}

	return nil
}

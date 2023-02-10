package val

import (
	"fmt"
	"net/mail"
	"regexp"
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

func ValidateWalletAddress(walletAddress string) error {
	re := regexp.MustCompile(`^0x[a-zA-Z0-9]{40}$`)

	if !re.MatchString(walletAddress) {
		return fmt.Errorf("Invalid wallet address")
	}
	return nil
}

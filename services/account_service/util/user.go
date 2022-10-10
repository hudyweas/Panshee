package util

import (
	"fmt"

	val "github.com/hudyweas/panshee/services/account_service/validators"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func CheckPassword(hash, password string) error {
	if err := val.ValidateHash(hash, password); err != nil {
		return fmt.Errorf("password doesn't match")
	}

	return nil
}

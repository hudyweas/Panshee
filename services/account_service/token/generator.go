package token

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type TokenData struct {
	paseto *paseto.V2
	key    []byte
}

type TokenGenerator interface {
	CreateToken(email string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}

func NewTokenGenerator(key string) (TokenGenerator, error) {
	//checking if key have good length
	if len(key) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters, key provided have %d", chacha20poly1305.KeySize, len(key))
	}

	tokenData := &TokenData{
		paseto: paseto.NewV2(),
		key:    []byte(key),
	}

	return tokenData, nil
}

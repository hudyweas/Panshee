package token

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (tokenData *TokenData) CreateToken(userID uuid.UUID, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(userID, duration)
	if err != nil {
		return "", payload, err
	}

	token, err := tokenData.paseto.Encrypt(tokenData.key, payload, nil)
	return token, payload, err
}

func (tokenData *TokenData) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := tokenData.paseto.Decrypt(token, tokenData.key, payload, nil)
	if err != nil {
		return nil, fmt.Errorf("invalid token")
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}

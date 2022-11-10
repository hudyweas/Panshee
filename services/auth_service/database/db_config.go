package database

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type data struct {
	*pg.DB
}

type Database interface {
	//UserPassword
	IsUserPasswordInDatabaseByUserID(id uuid.UUID) (bool, error)
	UserPasswordCreate(userPassword UserPassword) (*UserPassword, error)
	UserPasswordSelectByUserID(id uuid.UUID) (*UserPassword, error)

	//Session
	SessionCreate(session Session) (*Session, error)
	SessionSelectByID(id uuid.UUID) (*Session, error)

	//General
	Close()
	CheckConnection() error
}

type Session struct {
	ID    uuid.UUID `json:"session_ID"`
	UserID uuid.UUID `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	ClientIp	string `json:"client_ip"`
	UserAgent	string `json:"user_agent"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type UserPassword struct {
	ID uuid.UUID     `json:"id:"`
	UserID uuid.UUID `json:"user_id"`
	Password string  `json:"password"`
}

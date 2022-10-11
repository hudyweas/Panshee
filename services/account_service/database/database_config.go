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
	//User
	UserCreate(email, password string) (*User, error)
	UserSelectByEmail(email string) (*User, error)
	UserSelectById(id int) (*User, error)

	UserPasswordUpdateByID(id int, password string) (*User, error)

	//Session
	SessionCreate(session Session) (*Session, error)
	SessionSelectByID(id uuid.UUID) (*Session, error)

	//General
	Close()
	CheckConnection() error
}

type User struct {
	ID       int `json:"ID"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
}

type Session struct {
	ID       int `json:"ID"`
	SessionID    uuid.UUID `json:"session_ID"`
	Email        string    `json:"email"`
	RefreshToken string    `json:"refresh_token"`
	ClientIp	string `json:"client_ip"`
	UserAgent	string `json:"user_agent"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
}

package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type data struct {
	*bun.DB
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
	Init() error
}

type Session struct {
	bun.BaseModel

	ID    uuid.UUID 		`bun:"type:uuid,unique,notnull,default:gen_random_uuid()"`
	UserID uuid.UUID 		`bun:"type:uuid,notnull"`
	RefreshToken string    	`bun:",unique,notnull"`
	ClientIp	string 		`bun:",notnull"`
	UserAgent	string 		`bun:",notnull"`
	IsBlocked    bool      	`bun:",notnull"`
	ExpiresAt    time.Time 	`bun:",notnull"`
}

type UserPassword struct {
	bun.BaseModel

	ID uuid.UUID     `bun:"type:uuid,unique,notnull,default:gen_random_uuid()"`
	UserID uuid.UUID `bun:"type:uuid,notnull"`
	Password string  `bun:",notnull"`
}

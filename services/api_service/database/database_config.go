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
	//User
	UserCreate(email string) (*User, error)
	UserSelectByEmail(email string) (*User, error)
	UserSelectById(id uuid.UUID) (*User, error)

	//Wallet
	WalletSelectByID(user_id uuid.UUID) (*[]Wallet, error)
	WalletSelectByIDAndWalletAddress(user_id uuid.UUID, walletAddress string) (*[]Wallet, error)
	DeleteWalletByIdAndWalletAddress(user_id uuid.UUID, walletAddress string) error
	WalletCreate(wallet Wallet) (*Wallet, error)

	//General
	Close()
	CheckConnection() error
	Init() error
}

type User struct {
	bun.BaseModel

	ID       uuid.UUID 	`bun:"type:uuid,unique,notnull,default:gen_random_uuid()"`
	Email    string 	`bun:",notnull"`
	Active   bool   	`bun:"type:bool,default:true"`
	CreatedAt time.Time `bun:",notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",notnull,default:current_timestamp"`
	DeletedAt time.Time `bun:",notnull,default:current_timestamp"`
}

type Wallet struct {
	bun.BaseModel

	ID 					int 		`bun:"type:integer"`
	WalletAddress    	string 		`bun:",notnull"`
	UserID       		uuid.UUID 	`bun:"type:uuid"`
}

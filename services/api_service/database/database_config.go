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
}

type User struct {
	ID       uuid.UUID 	`json:"ID"`
	Email    string 	`json:"email"`
	Active   bool   	`json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Wallet struct {
	ID 					int 		`json:"ID"`
	WalletAddress    	string 		`json:"wallet_address"`
	UserID       		uuid.UUID 	`json:"user_id"`
}

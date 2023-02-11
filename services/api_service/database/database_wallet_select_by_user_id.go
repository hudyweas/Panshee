package database

import (
	"context"

	"github.com/google/uuid"
	e "github.com/hudyweas/panshee/services/api_service/errors"
)

func (db *data) WalletSelectByID(user_id uuid.UUID) (*[]Wallet, error) {
	if err := db.CheckConnection(); err != nil{
		return &[]Wallet{}, err
	}

	var walletAddresses []Wallet
	err := db.NewSelect().Model(&walletAddresses).Where("user_id = ?", user_id).Scan(context.Background())
	if err != nil {
		return &walletAddresses, e.DatabaseErrorWrapper(err)
	}

	return &walletAddresses, nil
}

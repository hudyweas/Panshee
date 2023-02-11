package database

import (
	"context"

	e "github.com/hudyweas/panshee/services/api_service/errors"
)

func (db *data) WalletCreate(wallet Wallet) (*Wallet, error) {
	if err := db.CheckConnection(); err != nil {
		return &Wallet{}, err
	}

	_, err := db.NewInsert().Model(&wallet).Exec(context.Background())

	if err != nil {
		return &wallet, e.DatabaseErrorWrapper(err)
	}

	return &wallet, nil
}

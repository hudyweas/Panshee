package database

import (
	"github.com/google/uuid"
	e "github.com/hudyweas/panshee/services/api_service/errors"
)

func (db *data) WalletSelectByIDAndWalletAddress(user_id uuid.UUID, walletAddress string) (*[]Wallet, error) {
	if err := db.CheckConnection(); err != nil{
		return &[]Wallet{}, err
	}

	var walletAddresses []Wallet

	err := db.Model(&walletAddresses).Where("user_id = ? AND wallet_address = ?", user_id, walletAddress).Select()
	if err != nil {
		return &walletAddresses, e.DatabaseErrorWrapper(err)
	}

	return &walletAddresses, nil
}

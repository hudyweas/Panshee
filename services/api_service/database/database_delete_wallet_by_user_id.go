package database

import (
	"fmt"

	"github.com/google/uuid"
	e "github.com/hudyweas/panshee/services/api_service/errors"
)

func (db *data) DeleteWalletByIdAndWalletAddress(user_id uuid.UUID, walletAddress string) error{
	if err := db.CheckConnection(); err != nil{
		return err
	}

	orm , err := db.Model(&Wallet{}).Where("user_id = ? AND wallet_address = ?", user_id, walletAddress).Delete()
	if err != nil {
		return e.DatabaseErrorWrapper(err)
	}

	if orm.RowsAffected() == 0 {
		return fmt.Errorf("no matching wallet address for given user")
	}

	return nil
}

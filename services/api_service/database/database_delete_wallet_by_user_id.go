package database

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	e "github.com/hudyweas/panshee/services/api_service/errors"
)

func (db *data) DeleteWalletByIdAndWalletAddress(user_id uuid.UUID, walletAddress string) error{
	if err := db.CheckConnection(); err != nil{
		return err
	}

	res , err := db.NewDelete().Model(&Wallet{}).Where("user_id = ? AND wallet_address = ?", user_id, walletAddress).Exec(context.Background())
	if err != nil {
		return e.DatabaseErrorWrapper(err)
	}

	rowsAffected, _ := res.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("no matching wallet address for given user")
	}

	return nil
}

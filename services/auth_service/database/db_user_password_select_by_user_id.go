package database

import (
	"context"

	"github.com/google/uuid"
	e "github.com/hudyweas/panshee/services/auth_service/errors"
)

func (db *data) UserPasswordSelectByUserID(id uuid.UUID) (*UserPassword, error) {
	if err := db.CheckConnection(); err != nil{
		return &UserPassword{}, err
	}

	userPassword := &UserPassword{}

	err := db.NewSelect().Model(userPassword).Where("user_id = ?", id).Scan(context.Background())
	if err != nil {
		return userPassword, e.DatabaseErrorWrapper(err)
	}

	return userPassword, nil
}

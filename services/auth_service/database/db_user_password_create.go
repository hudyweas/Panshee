package database

import (
	"context"

	e "github.com/hudyweas/panshee/services/auth_service/errors"
)

func (db *data) UserPasswordCreate(userPassword UserPassword) (*UserPassword, error) {
	if err := db.CheckConnection(); err != nil{
		return &UserPassword{}, err
	}

	_, err := db.NewInsert().Model(&userPassword).Exec(context.Background())
	if err != nil {
		return &userPassword, e.DatabaseErrorWrapper(err)
	}

	return &userPassword, nil
}

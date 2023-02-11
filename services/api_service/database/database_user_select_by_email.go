package database

import (
	"context"

	e "github.com/hudyweas/panshee/services/api_service/errors"
)

func (db *data) UserSelectByEmail(email string) (*User, error) {
	if err := db.CheckConnection(); err != nil{
		return &User{}, err
	}

	var user User
	err := db.NewSelect().Model(&user).Where("email = ?", email).Scan(context.Background())
	if err != nil {
		return &user, e.DatabaseErrorWrapper(err)
	}

	return &user, nil
}

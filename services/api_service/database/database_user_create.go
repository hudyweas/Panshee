package database

import (
	"context"

	e "github.com/hudyweas/panshee/services/api_service/errors"
)

func (db *data) UserCreate(email string) (*User, error) {
	if err := db.CheckConnection(); err != nil{
		return &User{}, err
	}

	user := User{}
	user.Email = email

	_, err := db.NewInsert().Model(&user).Exec(context.Background())

	if err != nil {
		return &user, e.DatabaseErrorWrapper(err)
	}

	return &user, nil
}

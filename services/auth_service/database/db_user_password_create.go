package database

import (
	e "github.com/hudyweas/panshee/services/auth_service/errors"
)

func (db *data) UserPasswordCreate(userPassword UserPassword) (*UserPassword, error) {
	if err := db.CheckConnection(); err != nil{
		return &UserPassword{}, err
	}

	_, err := db.Model(&userPassword).Insert()

	if err != nil {
		return &userPassword, e.DatabaseErrorWrapper(err)
	}

	return &userPassword, nil
}

package database

import (
	e "github.com/hudyweas/panshee/services/api_service/errors"
)

func (db *data) UserCreate(email string) (*User, error) {
	if err := db.CheckConnection(); err != nil{
		return &User{}, err
	}

	user := User{}
	user.Email = email

	_, err := db.Model(&user).Insert()

	if err != nil {
		return &user, e.DatabaseErrorWrapper(err)
	}

	return &user, nil
}

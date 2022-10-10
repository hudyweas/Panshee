package database

import (
	e "github.com/hudyweas/panshee/services/account_service/errors"
)

func (db *data) UserCreate(email, password string) (*User, error) {
	if err := db.CheckConnection(); err != nil{
		return &User{}, err
	}

	user := User{}
	user.Email = email
	user.Password = password

	_, err := db.Model(&user).Insert()

	if err != nil {
		return &user, e.DatabaseErrorWrapper(err)
	}

	return &user, nil
}

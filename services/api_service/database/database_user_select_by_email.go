package database

import (
	e "github.com/hudyweas/panshee/services/api_service/errors"
)

func (db *data) UserSelectByEmail(email string) (*User, error) {
	if err := db.CheckConnection(); err != nil{
		return &User{}, err
	}

	var user User

	err := db.Model(&user).Where("email = ?", email).Select()
	if err != nil {
		return &user, e.DatabaseErrorWrapper(err)
	}

	return &user, nil
}

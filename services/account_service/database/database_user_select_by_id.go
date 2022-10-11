package database

import (
	e "github.com/hudyweas/panshee/services/account_service/errors"
)

func (db *data) UserSelectById(id int) (*User, error) {
	if err := db.CheckConnection(); err != nil{
		return &User{}, err
	}

	var user User

	err := db.Model(&user).Where("id = ?", id).Select()
	if err != nil {
		return &user, e.DatabaseErrorWrapper(err)
	}

	return &user, nil
}

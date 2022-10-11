package database

import e "github.com/hudyweas/panshee/services/account_service/errors"

func (db *data) UserPasswordUpdateByID(id int, password string) (*User, error) {
	if err := db.CheckConnection(); err != nil {
		return &User{}, err
	}

	var user User

	_, err := db.Model(&user).Set("password = ?", password).Where("id = ?", id).Update()
	if err != nil {
		return &user, e.DatabaseErrorWrapper(err)
	}

	return &user, nil
}

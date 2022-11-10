package database

import (
	"github.com/google/uuid"
	e "github.com/hudyweas/panshee/services/api_service/errors"
)

func (db *data) UserSelectById(id uuid.UUID) (*User, error) {
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

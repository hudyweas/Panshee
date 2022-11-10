package database

import (
	"github.com/google/uuid"
	e "github.com/hudyweas/panshee/services/auth_service/errors"
)

func (db *data) IsUserPasswordInDatabaseByUserID(id uuid.UUID) (bool, error) {
	_, err := db.UserPasswordSelectByUserID(id)

	if err == nil{
		return true, nil
	}

	if err.Error() == e.ErrNoDataInDatabase{
		return false, nil
	}

	return true, err
}

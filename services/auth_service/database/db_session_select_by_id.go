package database

import (
	"github.com/google/uuid"
	e "github.com/hudyweas/panshee/services/auth_service/errors"
)

func (db *data) SessionSelectByID(id uuid.UUID) (*Session, error) {
	if err := db.CheckConnection(); err != nil{
		return &Session{}, err
	}

	session := &Session{}

	err := db.Model(session).Where("session_id = ?", id).Select()
	if err != nil {
		return session, e.DatabaseErrorWrapper(err)
	}

	return session, nil
}

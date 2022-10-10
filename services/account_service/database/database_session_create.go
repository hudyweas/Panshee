package database

import (
	e "github.com/hudyweas/panshee/services/account_service/errors"
)

func (db *data) SessionCreate(session Session) (*Session, error) {
	if err := db.CheckConnection(); err != nil{
		return &Session{}, err
	}

	_, err := db.Model(&session).Insert()

	if err != nil {
		return &session, e.DatabaseErrorWrapper(err)
	}

	return &session, nil
}

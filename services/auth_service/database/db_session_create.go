package database

import (
	"context"

	e "github.com/hudyweas/panshee/services/auth_service/errors"
)

func (db *data) SessionCreate(session Session) (*Session, error) {
	if err := db.CheckConnection(); err != nil{
		return &Session{}, err
	}

	_, err := db.NewInsert().Model(&session).Exec(context.Background())

	if err != nil {
		return &session, e.DatabaseErrorWrapper(err)
	}

	return &session, nil
}

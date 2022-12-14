package database

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

func Connect(config pg.Options) (Database, error) {
	db := pg.Connect(&config)

	return &data{db}, nil
}

func (db *data) CheckConnection() error {
	if err:= db.DB.Ping(context.Background()); err != nil{
		return fmt.Errorf("No connection to database: %s", err)
	}

	return nil
}

func (db *data) Close() {
	db.DB.Close()
}

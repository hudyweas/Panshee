package database

import (
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func Connect(address string, port string, username string, password string, database string) (Database){
	dsn := fmt.Sprintf(`postgres://%s:%s@%s:%s/%s?sslmode=disable`, username, password, address, port, database)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())

	return &data{db}
}

func (db *data) CheckConnection() error {
	if err:= db.Ping(); err != nil{
		return fmt.Errorf("No connection to database: %s", err)
	}

	return nil
}

func (db *data) Close() {
	db.DB.Close()
}

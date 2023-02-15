package e

import "fmt"

const (
	//database
	ErrNoDataInDatabase         = "no data in database"
)

func DatabaseErrorWrapper(err error) error{
	switch err.Error(){
	case "sql: no rows in result set":
		return fmt.Errorf(ErrNoDataInDatabase)
	}

	return err
}

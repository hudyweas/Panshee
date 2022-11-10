package e

import "fmt"

const (
	//database
	ErrNoDataInDatabase         = "no data in database"
)

//TODO:
func DatabaseErrorWrapper(err error) error{
	switch err.Error(){
	case "pg: no rows in result set":
		return fmt.Errorf(ErrNoDataInDatabase)
	}

	return err
}

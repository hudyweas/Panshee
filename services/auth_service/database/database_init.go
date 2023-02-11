package database

import "context"

func (db *data) Init() error {
	if err := db.CheckConnection(); err != nil {
		return err
	}

	//TODO:
	db.NewDropTable().Model(&Session{}).Exec(context.Background())
	db.NewDropTable().Model(&UserPassword{}).Exec(context.Background())

	_, err := db.NewCreateTable().IfNotExists().Model(&Session{}).Exec(context.Background())
	if err != nil{
		return err
	}

	_, err = db.NewCreateTable().IfNotExists().Model(&UserPassword{}).Exec(context.Background())
	if err != nil{
		return err
	}
	return nil
}

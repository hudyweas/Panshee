package database

import "context"

func (db *data) Init() error {
	if err := db.CheckConnection(); err != nil {
		return err
	}

	// db.NewDropTable().Model(&User{}).Exec(context.Background())
	// db.NewDropTable().Model(&Wallet{}).Exec(context.Background())

	_, err := db.NewCreateTable().IfNotExists().Model(&User{}).Exec(context.Background())
	if err != nil{
		return err
	}
	_, err =  db.NewCreateTable().IfNotExists().Model(&Wallet{}).Exec(context.Background())
	if err != nil{
		return err
	}
	return nil
}

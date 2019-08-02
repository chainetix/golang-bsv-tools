package database

func (db *DB) CreateTables() error {

	if err := db.CreateTableUser(); err != nil {
		return err
	}


	if err := db.CreateTableUserGroup(); err != nil {
		return err
	}


	if err := db.CreateTableWallet(); err != nil {
		return err
	}


	if err := db.CreateTableCurrency(); err != nil {
		return err
	}


	if err := db.CreateTableInput(); err != nil {
		return err
	}


	if err := db.CreateTableApiToken(); err != nil {
		return err
	}


	if err := db.CreateTableExchange(); err != nil {
		return err
	}


	if err := db.CreateTableFeedMessage(); err != nil {
		return err
	}


	if err := db.CreateTableTransaction(); err != nil {
		return err
	}


	if err := db.CreateTableAgent(); err != nil {
		return err
	}


	if err := db.CreateTableAddress(); err != nil {
		return err
	}


	if err := db.CreateTableBillingRecord(); err != nil {
		return err
	}


	if err := db.CreateTableOutput(); err != nil {
		return err
	}


	if err := db.CreateTableProject(); err != nil {
		return err
	}


	if err := db.CreateTableAsset(); err != nil {
		return err
	}


	if err := db.CreateTableStream(); err != nil {
		return err
	}


	if err := db.CreateTablePermission(); err != nil {
		return err
	}


	if err := db.CreateTableStreamItem(); err != nil {
		return err
	}


	if err := db.CreateTableProjectStats(); err != nil {
		return err
	}


	if err := db.CreateTableNetwork(); err != nil {
		return err
	}


	if err := db.CreateTableNode(); err != nil {
		return err
	}

	return nil
}
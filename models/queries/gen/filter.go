package database

import (
	"fmt"
	//"time"
	//"database/sql"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/filter"
)

func (db *DB) CreateTableFilter() error {
	q := "CREATE TABLE IF NOT EXISTS " + db.dbName + ".\"Filter\" (\"UID\" UUID PRIMARY KEY DEFAULT gen_random_uuid(), \"Salt\" UUID DEFAULT gen_random_uuid(), \"Created\" TIMESTAMP DEFAULT current_timestamp(), \"Name\" STRING NOT NULL, \"Type\" STRING NOT NULL, \"Restrictions\" STRING NOT NULL,\"Code\" STRING NOT NULL);"
	fmt.Println(q)
	_, err := db.DoExec(q)
	return err
}

func (db *DB) InsertFilter(Name string, Type string, Restrictions string, Code string) (*models.Filter, error) {
	row := &models.Filter{}
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Filter\" (\"Name\", \"Type\", \"Description\", \"Code\") VALUES ($1, $2, $3, $4) RETURNING Name, Type, Restrictions, Code;", Name, Type, Restrictions, Code).Scan(&row.Created, &row.Name, &row.Type, &row.Restrictions, &row.Code, &row.Salt, &row.UID); err != nil {
		return nil, err
	}
	return row, nil
}

func (db *DB) DeleteFilter(uid string) error {
	_, err := db.DoQuery("DELETE FROM " + db.dbName + ".\"Filter\" WHERE \"UID\" = $1;", uid)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) PutFilter(model *models.Filter) error {
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Filter\" (\"Name\", \"Type\", \"Restrictions\", \"Code\") VALUES ($1, $2, $3, $4) RETURNING \"Created\", \"Name\", \"Type\", \"Restrictions\", \"Code\", \"Salt\", \"UID\";", model.Name, model.Type, model.Restrictions, model.Code).Scan(&model.Created, &model.Name, &model.Type, &model.Restrictions, &model.Code, &model.Salt, &model.UID); err != nil {
		return err
	}
	return nil
}

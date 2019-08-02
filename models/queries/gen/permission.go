package database

import (
	"fmt"
	"time"
	"database/sql"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/project"
)

func (db *DB) CreateTablePermission() error {
	q := "CREATE TABLE IF NOT EXISTS " + db.dbName + ".\"Permission\" (\"User\" STRING NOT NULL, \"Address\" STRING NOT NULL, \"Action\" STRING NOT NULL, \"UID\" UUID PRIMARY KEY DEFAULT gen_random_uuid(), \"Salt\" UUID DEFAULT gen_random_uuid(), \"Created\" TIMESTAMP DEFAULT current_timestamp(), \"Project\" STRING NOT NULL, \"Wallet\" STRING NOT NULL, \"State\" BOOL NOT NULL);"
	fmt.Println(q)
	_, err := db.DoExec(q)
	return err
}

func (db *DB) InsertPermission(Action string, Address string, Project string, State bool, User string, Wallet string) (*models.Permission, error) {
	row := &models.Permission{}
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Permission\" (\"Action\", \"Address\", \"Project\", \"State\", \"User\", \"Wallet\") VALUES ($1, $2, $3, $4, $5, $6) RETURNING Action, Address, Project, State, User, Wallet;", Action, Address, Project, State, User, Wallet).Scan(&row.Action, &row.Address, &row.Created, &row.Project, &row.Salt, &row.State, &row.UID, &row.User, &row.Wallet); err != nil {
		return nil, err
	}
	return row, nil
}

func (db *DB) DeletePermission(uid string) error {
	_, err := db.DoQuery("DELETE FROM " + db.dbName + ".\"Permission\" WHERE \"UID\" = $1;", uid)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) PutPermission(model *models.Permission) error {
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Permission\" (\"Action\", \"Address\", \"Project\", \"State\", \"User\", \"Wallet\") VALUES ($1, $2, $3, $4, $5, $6) RETURNING \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\";", model.Action, model.Address, model.Project, model.State, model.User, model.Wallet).Scan(&model.Action, &model.Address, &model.Created, &model.Project, &model.Salt, &model.State, &model.UID, &model.User, &model.Wallet); err != nil {
		return err
	}
	return nil
}

func (db *DB) ScanToPermission(rows *sql.Rows) ([]*models.Permission, error) {
	defer rows.Close()
	results := []*models.Permission{}
	for rows.Next() {
		result := &models.Permission{}
		err := rows.Scan(&result.Action, &result.Address, &result.Created, &result.Project, &result.Salt, &result.State, &result.UID, &result.User, &result.Wallet)
		if err != nil {
			return nil, err
		}
		results = append(
			results,
			result,
		)
	}
	return results, nil
}

func (db *DB) CountPermission() (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM " + db.dbName + ".\"Permission\";")
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) QueryPermission() ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\";")
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) CountPermissionByUser(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Permission\" WHERE \"User\" = $1;", value)
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetPermissionByUser(value string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"User\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByUser(value string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"User\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByUserAndAddress(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"User\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByUserAndAddress(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"User\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByUserAndAction(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"User\" = $1 AND \"Action\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByUserAndAction(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"User\" = $1 AND \"Action\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByUserAndUID(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"User\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByUserAndUID(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"User\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByUserAndSalt(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"User\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByUserAndSalt(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"User\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByUserAndCreated(value1 string, value2 time.Time) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"User\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByUserAndCreated(value1 string, value2 time.Time) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"User\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByUserAndProject(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"User\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByUserAndProject(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"User\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByUserAndWallet(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"User\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByUserAndWallet(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"User\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByUserAndState(value1 string, value2 bool) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"User\" = $1 AND \"State\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByUserAndState(value1 string, value2 bool) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"User\" = $1 AND \"State\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) CountPermissionByAddress(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Permission\" WHERE \"Address\" = $1;", value)
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetPermissionByAddress(value string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Address\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByAddress(value string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Address\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByAddressAndUser(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Address\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByAddressAndUser(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Address\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByAddressAndAction(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Address\" = $1 AND \"Action\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByAddressAndAction(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Address\" = $1 AND \"Action\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByAddressAndUID(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Address\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByAddressAndUID(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Address\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByAddressAndSalt(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Address\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByAddressAndSalt(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Address\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByAddressAndCreated(value1 string, value2 time.Time) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Address\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByAddressAndCreated(value1 string, value2 time.Time) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Address\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByAddressAndProject(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Address\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByAddressAndProject(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Address\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByAddressAndWallet(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Address\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByAddressAndWallet(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Address\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByAddressAndState(value1 string, value2 bool) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Address\" = $1 AND \"State\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByAddressAndState(value1 string, value2 bool) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Address\" = $1 AND \"State\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) CountPermissionByAction(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Permission\" WHERE \"Action\" = $1;", value)
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetPermissionByAction(value string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Action\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByAction(value string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Action\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByActionAndUser(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Action\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByActionAndUser(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Action\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByActionAndAddress(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Action\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByActionAndAddress(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Action\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByActionAndUID(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Action\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByActionAndUID(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Action\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByActionAndSalt(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Action\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByActionAndSalt(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Action\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByActionAndCreated(value1 string, value2 time.Time) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Action\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByActionAndCreated(value1 string, value2 time.Time) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Action\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByActionAndProject(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Action\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByActionAndProject(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Action\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByActionAndWallet(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Action\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByActionAndWallet(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Action\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByActionAndState(value1 string, value2 bool) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Action\" = $1 AND \"State\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByActionAndState(value1 string, value2 bool) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Action\" = $1 AND \"State\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) CountPermissionByUID(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Permission\" WHERE \"UID\" = $1;", value)
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetPermissionByUID(value string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"UID\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByUID(value string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"UID\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) CountPermissionBySalt(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Permission\" WHERE \"Salt\" = $1;", value)
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetPermissionBySalt(value string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Salt\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionBySalt(value string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Salt\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionBySaltAndUser(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Salt\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionBySaltAndUser(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Salt\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionBySaltAndAddress(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Salt\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionBySaltAndAddress(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Salt\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionBySaltAndAction(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Salt\" = $1 AND \"Action\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionBySaltAndAction(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Salt\" = $1 AND \"Action\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionBySaltAndUID(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Salt\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionBySaltAndUID(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Salt\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionBySaltAndCreated(value1 string, value2 time.Time) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Salt\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionBySaltAndCreated(value1 string, value2 time.Time) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Salt\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionBySaltAndProject(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Salt\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionBySaltAndProject(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Salt\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionBySaltAndWallet(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Salt\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionBySaltAndWallet(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Salt\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionBySaltAndState(value1 string, value2 bool) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Salt\" = $1 AND \"State\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionBySaltAndState(value1 string, value2 bool) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Salt\" = $1 AND \"State\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) CountPermissionByCreated(value time.Time) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Permission\" WHERE \"Created\" = $1;", value)
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetPermissionByCreated(value time.Time) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Created\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByCreated(value time.Time) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Created\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByCreatedAndUser(value1 time.Time, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Created\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByCreatedAndUser(value1 time.Time, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Created\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByCreatedAndAddress(value1 time.Time, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Created\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByCreatedAndAddress(value1 time.Time, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Created\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByCreatedAndAction(value1 time.Time, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Created\" = $1 AND \"Action\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByCreatedAndAction(value1 time.Time, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Created\" = $1 AND \"Action\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByCreatedAndUID(value1 time.Time, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Created\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByCreatedAndUID(value1 time.Time, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Created\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByCreatedAndSalt(value1 time.Time, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Created\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByCreatedAndSalt(value1 time.Time, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Created\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByCreatedAndProject(value1 time.Time, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Created\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByCreatedAndProject(value1 time.Time, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Created\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByCreatedAndWallet(value1 time.Time, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Created\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByCreatedAndWallet(value1 time.Time, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Created\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByCreatedAndState(value1 time.Time, value2 bool) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Created\" = $1 AND \"State\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByCreatedAndState(value1 time.Time, value2 bool) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Created\" = $1 AND \"State\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) CountPermissionByProject(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Permission\" WHERE \"Project\" = $1;", value)
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetPermissionByProject(value string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Project\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByProject(value string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Project\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByProjectAndUser(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Project\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByProjectAndUser(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Project\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByProjectAndAddress(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Project\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByProjectAndAddress(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Project\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByProjectAndAction(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Project\" = $1 AND \"Action\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByProjectAndAction(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Project\" = $1 AND \"Action\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByProjectAndUID(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Project\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByProjectAndUID(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Project\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByProjectAndSalt(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Project\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByProjectAndSalt(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Project\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByProjectAndCreated(value1 string, value2 time.Time) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Project\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByProjectAndCreated(value1 string, value2 time.Time) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Project\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByProjectAndWallet(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Project\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByProjectAndWallet(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Project\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByProjectAndState(value1 string, value2 bool) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Project\" = $1 AND \"State\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByProjectAndState(value1 string, value2 bool) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Project\" = $1 AND \"State\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) CountPermissionByWallet(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Permission\" WHERE \"Wallet\" = $1;", value)
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetPermissionByWallet(value string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Wallet\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByWallet(value string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Wallet\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByWalletAndUser(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Wallet\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByWalletAndUser(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Wallet\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByWalletAndAddress(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Wallet\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByWalletAndAddress(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Wallet\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByWalletAndAction(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Wallet\" = $1 AND \"Action\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByWalletAndAction(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Wallet\" = $1 AND \"Action\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByWalletAndUID(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Wallet\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByWalletAndUID(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Wallet\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByWalletAndSalt(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Wallet\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByWalletAndSalt(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Wallet\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByWalletAndCreated(value1 string, value2 time.Time) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Wallet\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByWalletAndCreated(value1 string, value2 time.Time) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Wallet\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByWalletAndProject(value1 string, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Wallet\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByWalletAndProject(value1 string, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Wallet\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByWalletAndState(value1 string, value2 bool) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Wallet\" = $1 AND \"State\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByWalletAndState(value1 string, value2 bool) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"Wallet\" = $1 AND \"State\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) CountPermissionByState(value bool) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Permission\" WHERE \"State\" = $1;", value)
	if err != nil {
		return 0, err
	}
	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *DB) GetPermissionByState(value bool) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"State\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByState(value bool) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"State\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByStateAndUser(value1 bool, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"State\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByStateAndUser(value1 bool, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"State\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByStateAndAddress(value1 bool, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"State\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByStateAndAddress(value1 bool, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"State\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByStateAndAction(value1 bool, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"State\" = $1 AND \"Action\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByStateAndAction(value1 bool, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"State\" = $1 AND \"Action\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByStateAndUID(value1 bool, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"State\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByStateAndUID(value1 bool, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"State\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByStateAndSalt(value1 bool, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"State\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByStateAndSalt(value1 bool, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"State\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByStateAndCreated(value1 bool, value2 time.Time) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"State\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByStateAndCreated(value1 bool, value2 time.Time) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"State\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByStateAndProject(value1 bool, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"State\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByStateAndProject(value1 bool, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"State\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}

func (db *DB) GetPermissionByStateAndWallet(value1 bool, value2 string) (bool, *models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"State\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToPermission(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryPermissionByStateAndWallet(value1 bool, value2 string) ([]*models.Permission, error) {
	rows, err := db.DoQuery("SELECT \"Action\", \"Address\", \"Created\", \"Project\", \"Salt\", \"State\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Permission\" WHERE \"State\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToPermission(rows)
}


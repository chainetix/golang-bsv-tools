package database

import (
	"fmt"
	"time"
	"database/sql"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/project"
)

func (db *DB) CreateTableWallet() error {
	q := "CREATE TABLE IF NOT EXISTS " + db.dbName + ".\"Wallet\" (\"User\" STRING NOT NULL, \"Label\" STRING NOT NULL, \"DefaultAddress\" STRING NOT NULL, \"UID\" UUID PRIMARY KEY DEFAULT gen_random_uuid(), \"Salt\" UUID DEFAULT gen_random_uuid(), \"Created\" TIMESTAMP DEFAULT current_timestamp(), \"Project\" STRING NOT NULL);"
	fmt.Println(q)
	_, err := db.DoExec(q)
	return err
}

func (db *DB) InsertWallet(DefaultAddress string, Label string, Project string, User string) (*models.Wallet, error) {
	row := &models.Wallet{}
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Wallet\" (\"DefaultAddress\", \"Label\", \"Project\", \"User\") VALUES ($1, $2, $3, $4) RETURNING DefaultAddress, Label, Project, User;", DefaultAddress, Label, Project, User).Scan(&row.Created, &row.DefaultAddress, &row.Label, &row.Project, &row.Salt, &row.UID, &row.User); err != nil {
		return nil, err
	}
	return row, nil
}

func (db *DB) DeleteWallet(uid string) error {
	_, err := db.DoQuery("DELETE FROM " + db.dbName + ".\"Wallet\" WHERE \"UID\" = $1;", uid)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) PutWallet(model *models.Wallet) error {
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Wallet\" (\"DefaultAddress\", \"Label\", \"Project\", \"User\") VALUES ($1, $2, $3, $4) RETURNING \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\";", model.DefaultAddress, model.Label, model.Project, model.User).Scan(&model.Created, &model.DefaultAddress, &model.Label, &model.Project, &model.Salt, &model.UID, &model.User); err != nil {
		return err
	}
	return nil
}

func (db *DB) ScanToWallet(rows *sql.Rows) ([]*models.Wallet, error) {
	defer rows.Close()
	results := []*models.Wallet{}
	for rows.Next() {
		result := &models.Wallet{}
		err := rows.Scan(&result.Created, &result.DefaultAddress, &result.Label, &result.Project, &result.Salt, &result.UID, &result.User)
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

func (db *DB) CountWallet() (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM " + db.dbName + ".\"Wallet\";")
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

func (db *DB) QueryWallet() ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\";")
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) CountWalletByUser(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Wallet\" WHERE \"User\" = $1;", value)
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

func (db *DB) GetWalletByUser(value string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"User\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByUser(value string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"User\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByUserAndLabel(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"User\" = $1 AND \"Label\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByUserAndLabel(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"User\" = $1 AND \"Label\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByUserAndDefaultAddress(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"User\" = $1 AND \"DefaultAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByUserAndDefaultAddress(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"User\" = $1 AND \"DefaultAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByUserAndUID(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"User\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByUserAndUID(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"User\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByUserAndSalt(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"User\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByUserAndSalt(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"User\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByUserAndCreated(value1 string, value2 time.Time) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"User\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByUserAndCreated(value1 string, value2 time.Time) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"User\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByUserAndProject(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"User\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByUserAndProject(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"User\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) CountWalletByLabel(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Wallet\" WHERE \"Label\" = $1;", value)
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

func (db *DB) GetWalletByLabel(value string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Label\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByLabel(value string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Label\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByLabelAndUser(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Label\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByLabelAndUser(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Label\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByLabelAndDefaultAddress(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Label\" = $1 AND \"DefaultAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByLabelAndDefaultAddress(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Label\" = $1 AND \"DefaultAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByLabelAndUID(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Label\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByLabelAndUID(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Label\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByLabelAndSalt(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Label\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByLabelAndSalt(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Label\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByLabelAndCreated(value1 string, value2 time.Time) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Label\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByLabelAndCreated(value1 string, value2 time.Time) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Label\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByLabelAndProject(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Label\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByLabelAndProject(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Label\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) CountWalletByDefaultAddress(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Wallet\" WHERE \"DefaultAddress\" = $1;", value)
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

func (db *DB) GetWalletByDefaultAddress(value string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"DefaultAddress\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByDefaultAddress(value string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"DefaultAddress\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByDefaultAddressAndUser(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"DefaultAddress\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByDefaultAddressAndUser(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"DefaultAddress\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByDefaultAddressAndLabel(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"DefaultAddress\" = $1 AND \"Label\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByDefaultAddressAndLabel(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"DefaultAddress\" = $1 AND \"Label\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByDefaultAddressAndUID(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"DefaultAddress\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByDefaultAddressAndUID(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"DefaultAddress\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByDefaultAddressAndSalt(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"DefaultAddress\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByDefaultAddressAndSalt(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"DefaultAddress\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByDefaultAddressAndCreated(value1 string, value2 time.Time) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"DefaultAddress\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByDefaultAddressAndCreated(value1 string, value2 time.Time) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"DefaultAddress\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByDefaultAddressAndProject(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"DefaultAddress\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByDefaultAddressAndProject(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"DefaultAddress\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) CountWalletByUID(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Wallet\" WHERE \"UID\" = $1;", value)
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

func (db *DB) GetWalletByUID(value string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"UID\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByUID(value string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"UID\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) CountWalletBySalt(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Wallet\" WHERE \"Salt\" = $1;", value)
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

func (db *DB) GetWalletBySalt(value string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Salt\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletBySalt(value string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Salt\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletBySaltAndUser(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Salt\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletBySaltAndUser(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Salt\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletBySaltAndLabel(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Salt\" = $1 AND \"Label\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletBySaltAndLabel(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Salt\" = $1 AND \"Label\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletBySaltAndDefaultAddress(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Salt\" = $1 AND \"DefaultAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletBySaltAndDefaultAddress(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Salt\" = $1 AND \"DefaultAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletBySaltAndUID(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Salt\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletBySaltAndUID(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Salt\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletBySaltAndCreated(value1 string, value2 time.Time) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Salt\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletBySaltAndCreated(value1 string, value2 time.Time) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Salt\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletBySaltAndProject(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Salt\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletBySaltAndProject(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Salt\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) CountWalletByCreated(value time.Time) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Wallet\" WHERE \"Created\" = $1;", value)
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

func (db *DB) GetWalletByCreated(value time.Time) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Created\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByCreated(value time.Time) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Created\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByCreatedAndUser(value1 time.Time, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Created\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByCreatedAndUser(value1 time.Time, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Created\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByCreatedAndLabel(value1 time.Time, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Created\" = $1 AND \"Label\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByCreatedAndLabel(value1 time.Time, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Created\" = $1 AND \"Label\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByCreatedAndDefaultAddress(value1 time.Time, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Created\" = $1 AND \"DefaultAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByCreatedAndDefaultAddress(value1 time.Time, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Created\" = $1 AND \"DefaultAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByCreatedAndUID(value1 time.Time, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Created\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByCreatedAndUID(value1 time.Time, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Created\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByCreatedAndSalt(value1 time.Time, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Created\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByCreatedAndSalt(value1 time.Time, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Created\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByCreatedAndProject(value1 time.Time, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Created\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByCreatedAndProject(value1 time.Time, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Created\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) CountWalletByProject(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Wallet\" WHERE \"Project\" = $1;", value)
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

func (db *DB) GetWalletByProject(value string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Project\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByProject(value string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Project\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByProjectAndUser(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Project\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByProjectAndUser(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Project\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByProjectAndLabel(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Project\" = $1 AND \"Label\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByProjectAndLabel(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Project\" = $1 AND \"Label\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByProjectAndDefaultAddress(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Project\" = $1 AND \"DefaultAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByProjectAndDefaultAddress(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Project\" = $1 AND \"DefaultAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByProjectAndUID(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Project\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByProjectAndUID(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Project\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByProjectAndSalt(value1 string, value2 string) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Project\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByProjectAndSalt(value1 string, value2 string) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Project\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}

func (db *DB) GetWalletByProjectAndCreated(value1 string, value2 time.Time) (bool, *models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Project\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToWallet(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryWalletByProjectAndCreated(value1 string, value2 time.Time) ([]*models.Wallet, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"DefaultAddress\", \"Label\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"Wallet\" WHERE \"Project\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToWallet(rows)
}


package database

import (
	"fmt"
	"time"
	"database/sql"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/project"
)

func (db *DB) CreateTableAddress() error {
	q := "CREATE TABLE IF NOT EXISTS " + db.dbName + ".\"Address\" (\"Wallet\" STRING NOT NULL, \"IsDefault\" BOOL NOT NULL, \"Seed\" STRING NOT NULL, \"Project\" STRING NOT NULL, \"Salt\" UUID DEFAULT gen_random_uuid(), \"Created\" TIMESTAMP DEFAULT current_timestamp(), \"User\" STRING NOT NULL, \"Addr\" STRING NOT NULL, \"UID\" UUID PRIMARY KEY DEFAULT gen_random_uuid());"
	fmt.Println(q)
	_, err := db.DoExec(q)
	return err
}

func (db *DB) InsertAddress(Addr string, IsDefault bool, Project string, Seed string, User string, Wallet string) (*models.Address, error) {
	row := &models.Address{}
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Address\" (\"Addr\", \"IsDefault\", \"Project\", \"Seed\", \"User\", \"Wallet\") VALUES ($1, $2, $3, $4, $5, $6) RETURNING Addr, IsDefault, Project, Seed, User, Wallet;", Addr, IsDefault, Project, Seed, User, Wallet).Scan(&row.Addr, &row.Created, &row.IsDefault, &row.Project, &row.Salt, &row.Seed, &row.UID, &row.User, &row.Wallet); err != nil {
		return nil, err
	}
	return row, nil
}

func (db *DB) DeleteAddress(uid string) error {
	_, err := db.DoQuery("DELETE FROM " + db.dbName + ".\"Address\" WHERE \"UID\" = $1;", uid)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) PutAddress(model *models.Address) error {
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Address\" (\"Addr\", \"IsDefault\", \"Project\", \"Seed\", \"User\", \"Wallet\") VALUES ($1, $2, $3, $4, $5, $6) RETURNING \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\";", model.Addr, model.IsDefault, model.Project, model.Seed, model.User, model.Wallet).Scan(&model.Addr, &model.Created, &model.IsDefault, &model.Project, &model.Salt, &model.Seed, &model.UID, &model.User, &model.Wallet); err != nil {
		return err
	}
	return nil
}

func (db *DB) ScanToAddress(rows *sql.Rows) ([]*models.Address, error) {
	defer rows.Close()
	results := []*models.Address{}
	for rows.Next() {
		result := &models.Address{}
		err := rows.Scan(&result.Addr, &result.Created, &result.IsDefault, &result.Project, &result.Salt, &result.Seed, &result.UID, &result.User, &result.Wallet)
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

func (db *DB) CountAddress() (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM " + db.dbName + ".\"Address\";")
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

func (db *DB) QueryAddress() ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\";")
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) CountAddressByWallet(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Address\" WHERE \"Wallet\" = $1;", value)
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

func (db *DB) GetAddressByWallet(value string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Wallet\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByWallet(value string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Wallet\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByWalletAndIsDefault(value1 string, value2 bool) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Wallet\" = $1 AND \"IsDefault\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByWalletAndIsDefault(value1 string, value2 bool) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Wallet\" = $1 AND \"IsDefault\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByWalletAndSeed(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Wallet\" = $1 AND \"Seed\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByWalletAndSeed(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Wallet\" = $1 AND \"Seed\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByWalletAndProject(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Wallet\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByWalletAndProject(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Wallet\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByWalletAndSalt(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Wallet\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByWalletAndSalt(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Wallet\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByWalletAndCreated(value1 string, value2 time.Time) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Wallet\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByWalletAndCreated(value1 string, value2 time.Time) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Wallet\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByWalletAndUser(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Wallet\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByWalletAndUser(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Wallet\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByWalletAndAddr(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Wallet\" = $1 AND \"Addr\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByWalletAndAddr(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Wallet\" = $1 AND \"Addr\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByWalletAndUID(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Wallet\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByWalletAndUID(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Wallet\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) CountAddressByIsDefault(value bool) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Address\" WHERE \"IsDefault\" = $1;", value)
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

func (db *DB) GetAddressByIsDefault(value bool) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"IsDefault\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByIsDefault(value bool) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"IsDefault\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByIsDefaultAndWallet(value1 bool, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"IsDefault\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByIsDefaultAndWallet(value1 bool, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"IsDefault\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByIsDefaultAndSeed(value1 bool, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"IsDefault\" = $1 AND \"Seed\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByIsDefaultAndSeed(value1 bool, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"IsDefault\" = $1 AND \"Seed\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByIsDefaultAndProject(value1 bool, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"IsDefault\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByIsDefaultAndProject(value1 bool, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"IsDefault\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByIsDefaultAndSalt(value1 bool, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"IsDefault\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByIsDefaultAndSalt(value1 bool, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"IsDefault\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByIsDefaultAndCreated(value1 bool, value2 time.Time) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"IsDefault\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByIsDefaultAndCreated(value1 bool, value2 time.Time) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"IsDefault\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByIsDefaultAndUser(value1 bool, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"IsDefault\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByIsDefaultAndUser(value1 bool, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"IsDefault\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByIsDefaultAndAddr(value1 bool, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"IsDefault\" = $1 AND \"Addr\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByIsDefaultAndAddr(value1 bool, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"IsDefault\" = $1 AND \"Addr\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByIsDefaultAndUID(value1 bool, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"IsDefault\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByIsDefaultAndUID(value1 bool, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"IsDefault\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) CountAddressBySeed(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Address\" WHERE \"Seed\" = $1;", value)
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

func (db *DB) GetAddressBySeed(value string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Seed\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressBySeed(value string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Seed\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressBySeedAndWallet(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Seed\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressBySeedAndWallet(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Seed\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressBySeedAndIsDefault(value1 string, value2 bool) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Seed\" = $1 AND \"IsDefault\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressBySeedAndIsDefault(value1 string, value2 bool) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Seed\" = $1 AND \"IsDefault\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressBySeedAndProject(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Seed\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressBySeedAndProject(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Seed\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressBySeedAndSalt(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Seed\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressBySeedAndSalt(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Seed\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressBySeedAndCreated(value1 string, value2 time.Time) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Seed\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressBySeedAndCreated(value1 string, value2 time.Time) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Seed\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressBySeedAndUser(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Seed\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressBySeedAndUser(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Seed\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressBySeedAndAddr(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Seed\" = $1 AND \"Addr\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressBySeedAndAddr(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Seed\" = $1 AND \"Addr\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressBySeedAndUID(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Seed\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressBySeedAndUID(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Seed\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) CountAddressByProject(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Address\" WHERE \"Project\" = $1;", value)
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

func (db *DB) GetAddressByProject(value string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Project\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByProject(value string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Project\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByProjectAndWallet(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Project\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByProjectAndWallet(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Project\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByProjectAndIsDefault(value1 string, value2 bool) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Project\" = $1 AND \"IsDefault\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByProjectAndIsDefault(value1 string, value2 bool) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Project\" = $1 AND \"IsDefault\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByProjectAndSeed(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Project\" = $1 AND \"Seed\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByProjectAndSeed(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Project\" = $1 AND \"Seed\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByProjectAndSalt(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Project\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByProjectAndSalt(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Project\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByProjectAndCreated(value1 string, value2 time.Time) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Project\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByProjectAndCreated(value1 string, value2 time.Time) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Project\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByProjectAndUser(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Project\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByProjectAndUser(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Project\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByProjectAndAddr(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Project\" = $1 AND \"Addr\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByProjectAndAddr(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Project\" = $1 AND \"Addr\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByProjectAndUID(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Project\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByProjectAndUID(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Project\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) CountAddressBySalt(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Address\" WHERE \"Salt\" = $1;", value)
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

func (db *DB) GetAddressBySalt(value string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Salt\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressBySalt(value string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Salt\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressBySaltAndWallet(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Salt\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressBySaltAndWallet(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Salt\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressBySaltAndIsDefault(value1 string, value2 bool) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Salt\" = $1 AND \"IsDefault\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressBySaltAndIsDefault(value1 string, value2 bool) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Salt\" = $1 AND \"IsDefault\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressBySaltAndSeed(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Salt\" = $1 AND \"Seed\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressBySaltAndSeed(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Salt\" = $1 AND \"Seed\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressBySaltAndProject(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Salt\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressBySaltAndProject(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Salt\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressBySaltAndCreated(value1 string, value2 time.Time) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Salt\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressBySaltAndCreated(value1 string, value2 time.Time) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Salt\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressBySaltAndUser(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Salt\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressBySaltAndUser(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Salt\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressBySaltAndAddr(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Salt\" = $1 AND \"Addr\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressBySaltAndAddr(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Salt\" = $1 AND \"Addr\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressBySaltAndUID(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Salt\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressBySaltAndUID(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Salt\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) CountAddressByCreated(value time.Time) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Address\" WHERE \"Created\" = $1;", value)
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

func (db *DB) GetAddressByCreated(value time.Time) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Created\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByCreated(value time.Time) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Created\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByCreatedAndWallet(value1 time.Time, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Created\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByCreatedAndWallet(value1 time.Time, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Created\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByCreatedAndIsDefault(value1 time.Time, value2 bool) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Created\" = $1 AND \"IsDefault\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByCreatedAndIsDefault(value1 time.Time, value2 bool) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Created\" = $1 AND \"IsDefault\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByCreatedAndSeed(value1 time.Time, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Created\" = $1 AND \"Seed\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByCreatedAndSeed(value1 time.Time, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Created\" = $1 AND \"Seed\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByCreatedAndProject(value1 time.Time, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Created\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByCreatedAndProject(value1 time.Time, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Created\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByCreatedAndSalt(value1 time.Time, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Created\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByCreatedAndSalt(value1 time.Time, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Created\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByCreatedAndUser(value1 time.Time, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Created\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByCreatedAndUser(value1 time.Time, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Created\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByCreatedAndAddr(value1 time.Time, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Created\" = $1 AND \"Addr\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByCreatedAndAddr(value1 time.Time, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Created\" = $1 AND \"Addr\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByCreatedAndUID(value1 time.Time, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Created\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByCreatedAndUID(value1 time.Time, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Created\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) CountAddressByUser(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Address\" WHERE \"User\" = $1;", value)
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

func (db *DB) GetAddressByUser(value string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"User\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByUser(value string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"User\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByUserAndWallet(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"User\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByUserAndWallet(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"User\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByUserAndIsDefault(value1 string, value2 bool) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"User\" = $1 AND \"IsDefault\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByUserAndIsDefault(value1 string, value2 bool) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"User\" = $1 AND \"IsDefault\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByUserAndSeed(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"User\" = $1 AND \"Seed\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByUserAndSeed(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"User\" = $1 AND \"Seed\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByUserAndProject(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"User\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByUserAndProject(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"User\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByUserAndSalt(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"User\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByUserAndSalt(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"User\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByUserAndCreated(value1 string, value2 time.Time) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"User\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByUserAndCreated(value1 string, value2 time.Time) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"User\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByUserAndAddr(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"User\" = $1 AND \"Addr\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByUserAndAddr(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"User\" = $1 AND \"Addr\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByUserAndUID(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"User\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByUserAndUID(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"User\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) CountAddressByAddr(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Address\" WHERE \"Addr\" = $1;", value)
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

func (db *DB) GetAddressByAddr(value string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Addr\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByAddr(value string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Addr\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByAddrAndWallet(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Addr\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByAddrAndWallet(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Addr\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByAddrAndIsDefault(value1 string, value2 bool) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Addr\" = $1 AND \"IsDefault\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByAddrAndIsDefault(value1 string, value2 bool) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Addr\" = $1 AND \"IsDefault\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByAddrAndSeed(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Addr\" = $1 AND \"Seed\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByAddrAndSeed(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Addr\" = $1 AND \"Seed\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByAddrAndProject(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Addr\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByAddrAndProject(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Addr\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByAddrAndSalt(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Addr\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByAddrAndSalt(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Addr\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByAddrAndCreated(value1 string, value2 time.Time) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Addr\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByAddrAndCreated(value1 string, value2 time.Time) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Addr\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByAddrAndUser(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Addr\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByAddrAndUser(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Addr\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) GetAddressByAddrAndUID(value1 string, value2 string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Addr\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByAddrAndUID(value1 string, value2 string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"Addr\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}

func (db *DB) CountAddressByUID(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Address\" WHERE \"UID\" = $1;", value)
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

func (db *DB) GetAddressByUID(value string) (bool, *models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"UID\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAddress(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAddressByUID(value string) ([]*models.Address, error) {
	rows, err := db.DoQuery("SELECT \"Addr\", \"Created\", \"IsDefault\", \"Project\", \"Salt\", \"Seed\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Address\" WHERE \"UID\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToAddress(rows)
}


package database

import (
	"fmt"
	"time"
	"database/sql"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/project"
)

func (db *DB) CreateTableCurrency() error {
	q := "CREATE TABLE IF NOT EXISTS " + db.dbName + ".\"Currency\" (\"Open\" BOOL NOT NULL, \"UID\" UUID PRIMARY KEY DEFAULT gen_random_uuid(), \"Salt\" UUID DEFAULT gen_random_uuid(), \"Created\" TIMESTAMP DEFAULT current_timestamp(), \"Project\" STRING NOT NULL, \"Name\" STRING NOT NULL, \"Alias\" STRING NOT NULL, \"Units\" INT NOT NULL);"
	fmt.Println(q)
	_, err := db.DoExec(q)
	return err
}

func (db *DB) InsertCurrency(Alias string, Name string, Open bool, Project string, Units int) (*models.Currency, error) {
	row := &models.Currency{}
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Currency\" (\"Alias\", \"Name\", \"Open\", \"Project\", \"Units\") VALUES ($1, $2, $3, $4, $5) RETURNING Alias, Name, Open, Project, Units;", Alias, Name, Open, Project, Units).Scan(&row.Alias, &row.Created, &row.Name, &row.Open, &row.Project, &row.Salt, &row.UID, &row.Units); err != nil {
		return nil, err
	}
	return row, nil
}

func (db *DB) DeleteCurrency(uid string) error {
	_, err := db.DoQuery("DELETE FROM " + db.dbName + ".\"Currency\" WHERE \"UID\" = $1;", uid)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) PutCurrency(model *models.Currency) error {
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Currency\" (\"Alias\", \"Name\", \"Open\", \"Project\", \"Units\") VALUES ($1, $2, $3, $4, $5) RETURNING \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\";", model.Alias, model.Name, model.Open, model.Project, model.Units).Scan(&model.Alias, &model.Created, &model.Name, &model.Open, &model.Project, &model.Salt, &model.UID, &model.Units); err != nil {
		return err
	}
	return nil
}

func (db *DB) ScanToCurrency(rows *sql.Rows) ([]*models.Currency, error) {
	defer rows.Close()
	results := []*models.Currency{}
	for rows.Next() {
		result := &models.Currency{}
		err := rows.Scan(&result.Alias, &result.Created, &result.Name, &result.Open, &result.Project, &result.Salt, &result.UID, &result.Units)
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

func (db *DB) CountCurrency() (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM " + db.dbName + ".\"Currency\";")
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

func (db *DB) QueryCurrency() ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\";")
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) CountCurrencyByOpen(value bool) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Currency\" WHERE \"Open\" = $1;", value)
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

func (db *DB) GetCurrencyByOpen(value bool) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Open\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByOpen(value bool) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Open\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByOpenAndUID(value1 bool, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Open\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByOpenAndUID(value1 bool, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Open\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByOpenAndSalt(value1 bool, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Open\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByOpenAndSalt(value1 bool, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Open\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByOpenAndCreated(value1 bool, value2 time.Time) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Open\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByOpenAndCreated(value1 bool, value2 time.Time) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Open\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByOpenAndProject(value1 bool, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Open\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByOpenAndProject(value1 bool, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Open\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByOpenAndName(value1 bool, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Open\" = $1 AND \"Name\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByOpenAndName(value1 bool, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Open\" = $1 AND \"Name\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByOpenAndAlias(value1 bool, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Open\" = $1 AND \"Alias\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByOpenAndAlias(value1 bool, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Open\" = $1 AND \"Alias\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByOpenAndUnits(value1 bool, value2 int) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Open\" = $1 AND \"Units\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByOpenAndUnits(value1 bool, value2 int) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Open\" = $1 AND \"Units\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) CountCurrencyByUID(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Currency\" WHERE \"UID\" = $1;", value)
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

func (db *DB) GetCurrencyByUID(value string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"UID\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByUID(value string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"UID\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) CountCurrencyBySalt(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Currency\" WHERE \"Salt\" = $1;", value)
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

func (db *DB) GetCurrencyBySalt(value string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Salt\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyBySalt(value string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Salt\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyBySaltAndOpen(value1 string, value2 bool) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Salt\" = $1 AND \"Open\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyBySaltAndOpen(value1 string, value2 bool) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Salt\" = $1 AND \"Open\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyBySaltAndUID(value1 string, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Salt\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyBySaltAndUID(value1 string, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Salt\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyBySaltAndCreated(value1 string, value2 time.Time) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Salt\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyBySaltAndCreated(value1 string, value2 time.Time) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Salt\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyBySaltAndProject(value1 string, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Salt\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyBySaltAndProject(value1 string, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Salt\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyBySaltAndName(value1 string, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Salt\" = $1 AND \"Name\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyBySaltAndName(value1 string, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Salt\" = $1 AND \"Name\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyBySaltAndAlias(value1 string, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Salt\" = $1 AND \"Alias\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyBySaltAndAlias(value1 string, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Salt\" = $1 AND \"Alias\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyBySaltAndUnits(value1 string, value2 int) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Salt\" = $1 AND \"Units\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyBySaltAndUnits(value1 string, value2 int) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Salt\" = $1 AND \"Units\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) CountCurrencyByCreated(value time.Time) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Currency\" WHERE \"Created\" = $1;", value)
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

func (db *DB) GetCurrencyByCreated(value time.Time) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Created\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByCreated(value time.Time) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Created\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByCreatedAndOpen(value1 time.Time, value2 bool) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Created\" = $1 AND \"Open\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByCreatedAndOpen(value1 time.Time, value2 bool) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Created\" = $1 AND \"Open\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByCreatedAndUID(value1 time.Time, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Created\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByCreatedAndUID(value1 time.Time, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Created\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByCreatedAndSalt(value1 time.Time, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Created\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByCreatedAndSalt(value1 time.Time, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Created\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByCreatedAndProject(value1 time.Time, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Created\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByCreatedAndProject(value1 time.Time, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Created\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByCreatedAndName(value1 time.Time, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Created\" = $1 AND \"Name\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByCreatedAndName(value1 time.Time, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Created\" = $1 AND \"Name\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByCreatedAndAlias(value1 time.Time, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Created\" = $1 AND \"Alias\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByCreatedAndAlias(value1 time.Time, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Created\" = $1 AND \"Alias\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByCreatedAndUnits(value1 time.Time, value2 int) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Created\" = $1 AND \"Units\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByCreatedAndUnits(value1 time.Time, value2 int) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Created\" = $1 AND \"Units\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) CountCurrencyByProject(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Currency\" WHERE \"Project\" = $1;", value)
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

func (db *DB) GetCurrencyByProject(value string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Project\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByProject(value string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Project\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByProjectAndOpen(value1 string, value2 bool) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Project\" = $1 AND \"Open\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByProjectAndOpen(value1 string, value2 bool) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Project\" = $1 AND \"Open\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByProjectAndUID(value1 string, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Project\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByProjectAndUID(value1 string, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Project\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByProjectAndSalt(value1 string, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Project\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByProjectAndSalt(value1 string, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Project\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByProjectAndCreated(value1 string, value2 time.Time) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Project\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByProjectAndCreated(value1 string, value2 time.Time) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Project\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByProjectAndName(value1 string, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Project\" = $1 AND \"Name\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByProjectAndName(value1 string, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Project\" = $1 AND \"Name\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByProjectAndAlias(value1 string, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Project\" = $1 AND \"Alias\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByProjectAndAlias(value1 string, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Project\" = $1 AND \"Alias\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByProjectAndUnits(value1 string, value2 int) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Project\" = $1 AND \"Units\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByProjectAndUnits(value1 string, value2 int) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Project\" = $1 AND \"Units\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) CountCurrencyByName(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Currency\" WHERE \"Name\" = $1;", value)
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

func (db *DB) GetCurrencyByName(value string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Name\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByName(value string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Name\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByNameAndOpen(value1 string, value2 bool) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Name\" = $1 AND \"Open\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByNameAndOpen(value1 string, value2 bool) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Name\" = $1 AND \"Open\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByNameAndUID(value1 string, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Name\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByNameAndUID(value1 string, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Name\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByNameAndSalt(value1 string, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Name\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByNameAndSalt(value1 string, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Name\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByNameAndCreated(value1 string, value2 time.Time) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Name\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByNameAndCreated(value1 string, value2 time.Time) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Name\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByNameAndProject(value1 string, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Name\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByNameAndProject(value1 string, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Name\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByNameAndAlias(value1 string, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Name\" = $1 AND \"Alias\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByNameAndAlias(value1 string, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Name\" = $1 AND \"Alias\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByNameAndUnits(value1 string, value2 int) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Name\" = $1 AND \"Units\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByNameAndUnits(value1 string, value2 int) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Name\" = $1 AND \"Units\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) CountCurrencyByAlias(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Currency\" WHERE \"Alias\" = $1;", value)
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

func (db *DB) GetCurrencyByAlias(value string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Alias\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByAlias(value string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Alias\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByAliasAndOpen(value1 string, value2 bool) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Alias\" = $1 AND \"Open\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByAliasAndOpen(value1 string, value2 bool) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Alias\" = $1 AND \"Open\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByAliasAndUID(value1 string, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Alias\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByAliasAndUID(value1 string, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Alias\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByAliasAndSalt(value1 string, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Alias\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByAliasAndSalt(value1 string, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Alias\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByAliasAndCreated(value1 string, value2 time.Time) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Alias\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByAliasAndCreated(value1 string, value2 time.Time) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Alias\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByAliasAndProject(value1 string, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Alias\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByAliasAndProject(value1 string, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Alias\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByAliasAndName(value1 string, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Alias\" = $1 AND \"Name\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByAliasAndName(value1 string, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Alias\" = $1 AND \"Name\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByAliasAndUnits(value1 string, value2 int) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Alias\" = $1 AND \"Units\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByAliasAndUnits(value1 string, value2 int) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Alias\" = $1 AND \"Units\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) CountCurrencyByUnits(value int) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Currency\" WHERE \"Units\" = $1;", value)
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

func (db *DB) GetCurrencyByUnits(value int) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Units\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByUnits(value int) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Units\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByUnitsAndOpen(value1 int, value2 bool) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Units\" = $1 AND \"Open\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByUnitsAndOpen(value1 int, value2 bool) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Units\" = $1 AND \"Open\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByUnitsAndUID(value1 int, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Units\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByUnitsAndUID(value1 int, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Units\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByUnitsAndSalt(value1 int, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Units\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByUnitsAndSalt(value1 int, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Units\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByUnitsAndCreated(value1 int, value2 time.Time) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Units\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByUnitsAndCreated(value1 int, value2 time.Time) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Units\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByUnitsAndProject(value1 int, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Units\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByUnitsAndProject(value1 int, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Units\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByUnitsAndName(value1 int, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Units\" = $1 AND \"Name\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByUnitsAndName(value1 int, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Units\" = $1 AND \"Name\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}

func (db *DB) GetCurrencyByUnitsAndAlias(value1 int, value2 string) (bool, *models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Units\" = $1 AND \"Alias\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToCurrency(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryCurrencyByUnitsAndAlias(value1 int, value2 string) ([]*models.Currency, error) {
	rows, err := db.DoQuery("SELECT \"Alias\", \"Created\", \"Name\", \"Open\", \"Project\", \"Salt\", \"UID\", \"Units\" FROM " + db.dbName + ".\"Currency\" WHERE \"Units\" = $1 AND \"Alias\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToCurrency(rows)
}


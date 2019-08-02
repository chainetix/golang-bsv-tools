package database

import (
	"fmt"
	"time"
	"database/sql"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/project"
)

func (db *DB) CreateTableExchange() error {
	q := "CREATE TABLE IF NOT EXISTS " + db.dbName + ".\"Exchange\" (\"UID\" UUID PRIMARY KEY DEFAULT gen_random_uuid(), \"Salt\" UUID DEFAULT gen_random_uuid(), \"Project\" STRING NOT NULL, \"Expiry\" int64 NOT NULL, \"Tx\" STRING NOT NULL, \"Created\" TIMESTAMP DEFAULT current_timestamp(), \"GiveCurrency\" STRING NOT NULL, \"GiveQuantity\" FLOAT NOT NULL, \"RecvCurrency\" STRING NOT NULL, \"RecvQuantity\" FLOAT NOT NULL);"
	fmt.Println(q)
	_, err := db.DoExec(q)
	return err
}

func (db *DB) InsertExchange(Expiry int64, GiveCurrency string, GiveQuantity float64, Project string, RecvCurrency string, RecvQuantity float64, Tx string) (*models.Exchange, error) {
	row := &models.Exchange{}
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Exchange\" (\"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Tx\") VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING Expiry, GiveCurrency, GiveQuantity, Project, RecvCurrency, RecvQuantity, Tx;", Expiry, GiveCurrency, GiveQuantity, Project, RecvCurrency, RecvQuantity, Tx).Scan(&row.Created, &row.Expiry, &row.GiveCurrency, &row.GiveQuantity, &row.Project, &row.RecvCurrency, &row.RecvQuantity, &row.Salt, &row.Tx, &row.UID); err != nil {
		return nil, err
	}
	return row, nil
}

func (db *DB) DeleteExchange(uid string) error {
	_, err := db.DoQuery("DELETE FROM " + db.dbName + ".\"Exchange\" WHERE \"UID\" = $1;", uid)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) PutExchange(model *models.Exchange) error {
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Exchange\" (\"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Tx\") VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\";", model.Expiry, model.GiveCurrency, model.GiveQuantity, model.Project, model.RecvCurrency, model.RecvQuantity, model.Tx).Scan(&model.Created, &model.Expiry, &model.GiveCurrency, &model.GiveQuantity, &model.Project, &model.RecvCurrency, &model.RecvQuantity, &model.Salt, &model.Tx, &model.UID); err != nil {
		return err
	}
	return nil
}

func (db *DB) ScanToExchange(rows *sql.Rows) ([]*models.Exchange, error) {
	defer rows.Close()
	results := []*models.Exchange{}
	for rows.Next() {
		result := &models.Exchange{}
		err := rows.Scan(&result.Created, &result.Expiry, &result.GiveCurrency, &result.GiveQuantity, &result.Project, &result.RecvCurrency, &result.RecvQuantity, &result.Salt, &result.Tx, &result.UID)
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

func (db *DB) CountExchange() (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM " + db.dbName + ".\"Exchange\";")
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

func (db *DB) QueryExchange() ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\";")
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) CountExchangeByUID(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Exchange\" WHERE \"UID\" = $1;", value)
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

func (db *DB) GetExchangeByUID(value string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"UID\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByUID(value string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"UID\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) CountExchangeBySalt(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Exchange\" WHERE \"Salt\" = $1;", value)
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

func (db *DB) GetExchangeBySalt(value string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Salt\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeBySalt(value string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Salt\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeBySaltAndUID(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Salt\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeBySaltAndUID(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Salt\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeBySaltAndProject(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Salt\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeBySaltAndProject(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Salt\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeBySaltAndExpiry(value1 string, value2 int64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Salt\" = $1 AND \"Expiry\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeBySaltAndExpiry(value1 string, value2 int64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Salt\" = $1 AND \"Expiry\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeBySaltAndTx(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Salt\" = $1 AND \"Tx\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeBySaltAndTx(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Salt\" = $1 AND \"Tx\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeBySaltAndCreated(value1 string, value2 time.Time) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Salt\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeBySaltAndCreated(value1 string, value2 time.Time) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Salt\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeBySaltAndGiveCurrency(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Salt\" = $1 AND \"GiveCurrency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeBySaltAndGiveCurrency(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Salt\" = $1 AND \"GiveCurrency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeBySaltAndGiveQuantity(value1 string, value2 float64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Salt\" = $1 AND \"GiveQuantity\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeBySaltAndGiveQuantity(value1 string, value2 float64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Salt\" = $1 AND \"GiveQuantity\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeBySaltAndRecvCurrency(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Salt\" = $1 AND \"RecvCurrency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeBySaltAndRecvCurrency(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Salt\" = $1 AND \"RecvCurrency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeBySaltAndRecvQuantity(value1 string, value2 float64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Salt\" = $1 AND \"RecvQuantity\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeBySaltAndRecvQuantity(value1 string, value2 float64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Salt\" = $1 AND \"RecvQuantity\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) CountExchangeByProject(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Exchange\" WHERE \"Project\" = $1;", value)
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

func (db *DB) GetExchangeByProject(value string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Project\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByProject(value string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Project\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByProjectAndUID(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Project\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByProjectAndUID(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Project\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByProjectAndSalt(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Project\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByProjectAndSalt(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Project\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByProjectAndExpiry(value1 string, value2 int64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Project\" = $1 AND \"Expiry\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByProjectAndExpiry(value1 string, value2 int64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Project\" = $1 AND \"Expiry\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByProjectAndTx(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Project\" = $1 AND \"Tx\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByProjectAndTx(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Project\" = $1 AND \"Tx\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByProjectAndCreated(value1 string, value2 time.Time) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Project\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByProjectAndCreated(value1 string, value2 time.Time) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Project\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByProjectAndGiveCurrency(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Project\" = $1 AND \"GiveCurrency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByProjectAndGiveCurrency(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Project\" = $1 AND \"GiveCurrency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByProjectAndGiveQuantity(value1 string, value2 float64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Project\" = $1 AND \"GiveQuantity\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByProjectAndGiveQuantity(value1 string, value2 float64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Project\" = $1 AND \"GiveQuantity\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByProjectAndRecvCurrency(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Project\" = $1 AND \"RecvCurrency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByProjectAndRecvCurrency(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Project\" = $1 AND \"RecvCurrency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByProjectAndRecvQuantity(value1 string, value2 float64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Project\" = $1 AND \"RecvQuantity\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByProjectAndRecvQuantity(value1 string, value2 float64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Project\" = $1 AND \"RecvQuantity\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) CountExchangeByExpiry(value int64) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Exchange\" WHERE \"Expiry\" = $1;", value)
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

func (db *DB) GetExchangeByExpiry(value int64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Expiry\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByExpiry(value int64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Expiry\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByExpiryAndUID(value1 int64, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Expiry\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByExpiryAndUID(value1 int64, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Expiry\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByExpiryAndSalt(value1 int64, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Expiry\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByExpiryAndSalt(value1 int64, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Expiry\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByExpiryAndProject(value1 int64, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Expiry\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByExpiryAndProject(value1 int64, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Expiry\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByExpiryAndTx(value1 int64, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Expiry\" = $1 AND \"Tx\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByExpiryAndTx(value1 int64, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Expiry\" = $1 AND \"Tx\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByExpiryAndCreated(value1 int64, value2 time.Time) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Expiry\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByExpiryAndCreated(value1 int64, value2 time.Time) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Expiry\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByExpiryAndGiveCurrency(value1 int64, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Expiry\" = $1 AND \"GiveCurrency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByExpiryAndGiveCurrency(value1 int64, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Expiry\" = $1 AND \"GiveCurrency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByExpiryAndGiveQuantity(value1 int64, value2 float64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Expiry\" = $1 AND \"GiveQuantity\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByExpiryAndGiveQuantity(value1 int64, value2 float64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Expiry\" = $1 AND \"GiveQuantity\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByExpiryAndRecvCurrency(value1 int64, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Expiry\" = $1 AND \"RecvCurrency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByExpiryAndRecvCurrency(value1 int64, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Expiry\" = $1 AND \"RecvCurrency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByExpiryAndRecvQuantity(value1 int64, value2 float64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Expiry\" = $1 AND \"RecvQuantity\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByExpiryAndRecvQuantity(value1 int64, value2 float64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Expiry\" = $1 AND \"RecvQuantity\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) CountExchangeByTx(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Exchange\" WHERE \"Tx\" = $1;", value)
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

func (db *DB) GetExchangeByTx(value string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Tx\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByTx(value string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Tx\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByTxAndUID(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Tx\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByTxAndUID(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Tx\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByTxAndSalt(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Tx\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByTxAndSalt(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Tx\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByTxAndProject(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Tx\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByTxAndProject(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Tx\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByTxAndExpiry(value1 string, value2 int64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Tx\" = $1 AND \"Expiry\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByTxAndExpiry(value1 string, value2 int64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Tx\" = $1 AND \"Expiry\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByTxAndCreated(value1 string, value2 time.Time) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Tx\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByTxAndCreated(value1 string, value2 time.Time) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Tx\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByTxAndGiveCurrency(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Tx\" = $1 AND \"GiveCurrency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByTxAndGiveCurrency(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Tx\" = $1 AND \"GiveCurrency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByTxAndGiveQuantity(value1 string, value2 float64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Tx\" = $1 AND \"GiveQuantity\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByTxAndGiveQuantity(value1 string, value2 float64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Tx\" = $1 AND \"GiveQuantity\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByTxAndRecvCurrency(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Tx\" = $1 AND \"RecvCurrency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByTxAndRecvCurrency(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Tx\" = $1 AND \"RecvCurrency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByTxAndRecvQuantity(value1 string, value2 float64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Tx\" = $1 AND \"RecvQuantity\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByTxAndRecvQuantity(value1 string, value2 float64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Tx\" = $1 AND \"RecvQuantity\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) CountExchangeByCreated(value time.Time) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Exchange\" WHERE \"Created\" = $1;", value)
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

func (db *DB) GetExchangeByCreated(value time.Time) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Created\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByCreated(value time.Time) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Created\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByCreatedAndUID(value1 time.Time, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Created\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByCreatedAndUID(value1 time.Time, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Created\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByCreatedAndSalt(value1 time.Time, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Created\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByCreatedAndSalt(value1 time.Time, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Created\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByCreatedAndProject(value1 time.Time, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Created\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByCreatedAndProject(value1 time.Time, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Created\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByCreatedAndExpiry(value1 time.Time, value2 int64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Created\" = $1 AND \"Expiry\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByCreatedAndExpiry(value1 time.Time, value2 int64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Created\" = $1 AND \"Expiry\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByCreatedAndTx(value1 time.Time, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Created\" = $1 AND \"Tx\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByCreatedAndTx(value1 time.Time, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Created\" = $1 AND \"Tx\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByCreatedAndGiveCurrency(value1 time.Time, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Created\" = $1 AND \"GiveCurrency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByCreatedAndGiveCurrency(value1 time.Time, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Created\" = $1 AND \"GiveCurrency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByCreatedAndGiveQuantity(value1 time.Time, value2 float64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Created\" = $1 AND \"GiveQuantity\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByCreatedAndGiveQuantity(value1 time.Time, value2 float64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Created\" = $1 AND \"GiveQuantity\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByCreatedAndRecvCurrency(value1 time.Time, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Created\" = $1 AND \"RecvCurrency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByCreatedAndRecvCurrency(value1 time.Time, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Created\" = $1 AND \"RecvCurrency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByCreatedAndRecvQuantity(value1 time.Time, value2 float64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Created\" = $1 AND \"RecvQuantity\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByCreatedAndRecvQuantity(value1 time.Time, value2 float64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"Created\" = $1 AND \"RecvQuantity\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) CountExchangeByGiveCurrency(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Exchange\" WHERE \"GiveCurrency\" = $1;", value)
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

func (db *DB) GetExchangeByGiveCurrency(value string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveCurrency\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByGiveCurrency(value string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveCurrency\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByGiveCurrencyAndUID(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveCurrency\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByGiveCurrencyAndUID(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveCurrency\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByGiveCurrencyAndSalt(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveCurrency\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByGiveCurrencyAndSalt(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveCurrency\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByGiveCurrencyAndProject(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveCurrency\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByGiveCurrencyAndProject(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveCurrency\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByGiveCurrencyAndExpiry(value1 string, value2 int64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveCurrency\" = $1 AND \"Expiry\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByGiveCurrencyAndExpiry(value1 string, value2 int64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveCurrency\" = $1 AND \"Expiry\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByGiveCurrencyAndTx(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveCurrency\" = $1 AND \"Tx\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByGiveCurrencyAndTx(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveCurrency\" = $1 AND \"Tx\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByGiveCurrencyAndCreated(value1 string, value2 time.Time) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveCurrency\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByGiveCurrencyAndCreated(value1 string, value2 time.Time) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveCurrency\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByGiveCurrencyAndGiveQuantity(value1 string, value2 float64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveCurrency\" = $1 AND \"GiveQuantity\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByGiveCurrencyAndGiveQuantity(value1 string, value2 float64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveCurrency\" = $1 AND \"GiveQuantity\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByGiveCurrencyAndRecvCurrency(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveCurrency\" = $1 AND \"RecvCurrency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByGiveCurrencyAndRecvCurrency(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveCurrency\" = $1 AND \"RecvCurrency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByGiveCurrencyAndRecvQuantity(value1 string, value2 float64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveCurrency\" = $1 AND \"RecvQuantity\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByGiveCurrencyAndRecvQuantity(value1 string, value2 float64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveCurrency\" = $1 AND \"RecvQuantity\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) CountExchangeByGiveQuantity(value float64) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Exchange\" WHERE \"GiveQuantity\" = $1;", value)
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

func (db *DB) GetExchangeByGiveQuantity(value float64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveQuantity\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByGiveQuantity(value float64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveQuantity\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByGiveQuantityAndUID(value1 float64, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveQuantity\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByGiveQuantityAndUID(value1 float64, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveQuantity\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByGiveQuantityAndSalt(value1 float64, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveQuantity\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByGiveQuantityAndSalt(value1 float64, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveQuantity\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByGiveQuantityAndProject(value1 float64, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveQuantity\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByGiveQuantityAndProject(value1 float64, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveQuantity\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByGiveQuantityAndExpiry(value1 float64, value2 int64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveQuantity\" = $1 AND \"Expiry\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByGiveQuantityAndExpiry(value1 float64, value2 int64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveQuantity\" = $1 AND \"Expiry\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByGiveQuantityAndTx(value1 float64, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveQuantity\" = $1 AND \"Tx\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByGiveQuantityAndTx(value1 float64, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveQuantity\" = $1 AND \"Tx\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByGiveQuantityAndCreated(value1 float64, value2 time.Time) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveQuantity\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByGiveQuantityAndCreated(value1 float64, value2 time.Time) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveQuantity\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByGiveQuantityAndGiveCurrency(value1 float64, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveQuantity\" = $1 AND \"GiveCurrency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByGiveQuantityAndGiveCurrency(value1 float64, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveQuantity\" = $1 AND \"GiveCurrency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByGiveQuantityAndRecvCurrency(value1 float64, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveQuantity\" = $1 AND \"RecvCurrency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByGiveQuantityAndRecvCurrency(value1 float64, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveQuantity\" = $1 AND \"RecvCurrency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByGiveQuantityAndRecvQuantity(value1 float64, value2 float64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveQuantity\" = $1 AND \"RecvQuantity\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByGiveQuantityAndRecvQuantity(value1 float64, value2 float64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"GiveQuantity\" = $1 AND \"RecvQuantity\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) CountExchangeByRecvCurrency(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Exchange\" WHERE \"RecvCurrency\" = $1;", value)
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

func (db *DB) GetExchangeByRecvCurrency(value string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvCurrency\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByRecvCurrency(value string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvCurrency\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByRecvCurrencyAndUID(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvCurrency\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByRecvCurrencyAndUID(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvCurrency\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByRecvCurrencyAndSalt(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvCurrency\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByRecvCurrencyAndSalt(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvCurrency\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByRecvCurrencyAndProject(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvCurrency\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByRecvCurrencyAndProject(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvCurrency\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByRecvCurrencyAndExpiry(value1 string, value2 int64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvCurrency\" = $1 AND \"Expiry\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByRecvCurrencyAndExpiry(value1 string, value2 int64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvCurrency\" = $1 AND \"Expiry\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByRecvCurrencyAndTx(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvCurrency\" = $1 AND \"Tx\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByRecvCurrencyAndTx(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvCurrency\" = $1 AND \"Tx\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByRecvCurrencyAndCreated(value1 string, value2 time.Time) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvCurrency\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByRecvCurrencyAndCreated(value1 string, value2 time.Time) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvCurrency\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByRecvCurrencyAndGiveCurrency(value1 string, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvCurrency\" = $1 AND \"GiveCurrency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByRecvCurrencyAndGiveCurrency(value1 string, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvCurrency\" = $1 AND \"GiveCurrency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByRecvCurrencyAndGiveQuantity(value1 string, value2 float64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvCurrency\" = $1 AND \"GiveQuantity\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByRecvCurrencyAndGiveQuantity(value1 string, value2 float64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvCurrency\" = $1 AND \"GiveQuantity\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByRecvCurrencyAndRecvQuantity(value1 string, value2 float64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvCurrency\" = $1 AND \"RecvQuantity\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByRecvCurrencyAndRecvQuantity(value1 string, value2 float64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvCurrency\" = $1 AND \"RecvQuantity\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) CountExchangeByRecvQuantity(value float64) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Exchange\" WHERE \"RecvQuantity\" = $1;", value)
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

func (db *DB) GetExchangeByRecvQuantity(value float64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvQuantity\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByRecvQuantity(value float64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvQuantity\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByRecvQuantityAndUID(value1 float64, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvQuantity\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByRecvQuantityAndUID(value1 float64, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvQuantity\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByRecvQuantityAndSalt(value1 float64, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvQuantity\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByRecvQuantityAndSalt(value1 float64, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvQuantity\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByRecvQuantityAndProject(value1 float64, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvQuantity\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByRecvQuantityAndProject(value1 float64, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvQuantity\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByRecvQuantityAndExpiry(value1 float64, value2 int64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvQuantity\" = $1 AND \"Expiry\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByRecvQuantityAndExpiry(value1 float64, value2 int64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvQuantity\" = $1 AND \"Expiry\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByRecvQuantityAndTx(value1 float64, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvQuantity\" = $1 AND \"Tx\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByRecvQuantityAndTx(value1 float64, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvQuantity\" = $1 AND \"Tx\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByRecvQuantityAndCreated(value1 float64, value2 time.Time) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvQuantity\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByRecvQuantityAndCreated(value1 float64, value2 time.Time) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvQuantity\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByRecvQuantityAndGiveCurrency(value1 float64, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvQuantity\" = $1 AND \"GiveCurrency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByRecvQuantityAndGiveCurrency(value1 float64, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvQuantity\" = $1 AND \"GiveCurrency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByRecvQuantityAndGiveQuantity(value1 float64, value2 float64) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvQuantity\" = $1 AND \"GiveQuantity\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByRecvQuantityAndGiveQuantity(value1 float64, value2 float64) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvQuantity\" = $1 AND \"GiveQuantity\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}

func (db *DB) GetExchangeByRecvQuantityAndRecvCurrency(value1 float64, value2 string) (bool, *models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvQuantity\" = $1 AND \"RecvCurrency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToExchange(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryExchangeByRecvQuantityAndRecvCurrency(value1 float64, value2 string) ([]*models.Exchange, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Expiry\", \"GiveCurrency\", \"GiveQuantity\", \"Project\", \"RecvCurrency\", \"RecvQuantity\", \"Salt\", \"Tx\", \"UID\" FROM " + db.dbName + ".\"Exchange\" WHERE \"RecvQuantity\" = $1 AND \"RecvCurrency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToExchange(rows)
}


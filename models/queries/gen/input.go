package database

import (
	"fmt"
	"time"
	"database/sql"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/project"
)

func (db *DB) CreateTableInput() error {
	q := "CREATE TABLE IF NOT EXISTS " + db.dbName + ".\"Input\" (\"TxId\" STRING NOT NULL, \"Vout\" INT NOT NULL, \"Value\" FLOAT NOT NULL, \"Created\" TIMESTAMP DEFAULT current_timestamp());"
	fmt.Println(q)
	_, err := db.DoExec(q)
	return err
}

func (db *DB) InsertInput(TxId string, Value float64, Vout int) (*models.Input, error) {
	row := &models.Input{}
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Input\" (\"TxId\", \"Value\", \"Vout\") VALUES ($1, $2, $3) RETURNING TxId, Value, Vout;", TxId, Value, Vout).Scan(&row.Created, &row.TxId, &row.Value, &row.Vout); err != nil {
		return nil, err
	}
	return row, nil
}

func (db *DB) DeleteInput(uid string) error {
	_, err := db.DoQuery("DELETE FROM " + db.dbName + ".\"Input\" WHERE \"UID\" = $1;", uid)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) PutInput(model *models.Input) error {
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Input\" (\"TxId\", \"Value\", \"Vout\") VALUES ($1, $2, $3) RETURNING \"Created\", \"TxId\", \"Value\", \"Vout\";", model.TxId, model.Value, model.Vout).Scan(&model.Created, &model.TxId, &model.Value, &model.Vout); err != nil {
		return err
	}
	return nil
}

func (db *DB) ScanToInput(rows *sql.Rows) ([]*models.Input, error) {
	defer rows.Close()
	results := []*models.Input{}
	for rows.Next() {
		result := &models.Input{}
		err := rows.Scan(&result.Created, &result.TxId, &result.Value, &result.Vout)
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

func (db *DB) CountInput() (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM " + db.dbName + ".\"Input\";")
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

func (db *DB) QueryInput() ([]*models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\";")
	if err != nil {
		return nil, err
	}
	return db.ScanToInput(rows)
}

func (db *DB) CountInputByTxId(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Input\" WHERE \"TxId\" = $1;", value)
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

func (db *DB) GetInputByTxId(value string) (bool, *models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"TxId\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToInput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryInputByTxId(value string) ([]*models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"TxId\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToInput(rows)
}

func (db *DB) GetInputByTxIdAndVout(value1 string, value2 int) (bool, *models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"TxId\" = $1 AND \"Vout\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToInput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryInputByTxIdAndVout(value1 string, value2 int) ([]*models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"TxId\" = $1 AND \"Vout\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToInput(rows)
}

func (db *DB) GetInputByTxIdAndValue(value1 string, value2 float64) (bool, *models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"TxId\" = $1 AND \"Value\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToInput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryInputByTxIdAndValue(value1 string, value2 float64) ([]*models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"TxId\" = $1 AND \"Value\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToInput(rows)
}

func (db *DB) GetInputByTxIdAndCreated(value1 string, value2 time.Time) (bool, *models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"TxId\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToInput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryInputByTxIdAndCreated(value1 string, value2 time.Time) ([]*models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"TxId\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToInput(rows)
}

func (db *DB) CountInputByVout(value int) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Input\" WHERE \"Vout\" = $1;", value)
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

func (db *DB) GetInputByVout(value int) (bool, *models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"Vout\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToInput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryInputByVout(value int) ([]*models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"Vout\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToInput(rows)
}

func (db *DB) GetInputByVoutAndTxId(value1 int, value2 string) (bool, *models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"Vout\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToInput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryInputByVoutAndTxId(value1 int, value2 string) ([]*models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"Vout\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToInput(rows)
}

func (db *DB) GetInputByVoutAndValue(value1 int, value2 float64) (bool, *models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"Vout\" = $1 AND \"Value\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToInput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryInputByVoutAndValue(value1 int, value2 float64) ([]*models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"Vout\" = $1 AND \"Value\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToInput(rows)
}

func (db *DB) GetInputByVoutAndCreated(value1 int, value2 time.Time) (bool, *models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"Vout\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToInput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryInputByVoutAndCreated(value1 int, value2 time.Time) ([]*models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"Vout\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToInput(rows)
}

func (db *DB) CountInputByValue(value float64) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Input\" WHERE \"Value\" = $1;", value)
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

func (db *DB) GetInputByValue(value float64) (bool, *models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"Value\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToInput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryInputByValue(value float64) ([]*models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"Value\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToInput(rows)
}

func (db *DB) GetInputByValueAndTxId(value1 float64, value2 string) (bool, *models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"Value\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToInput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryInputByValueAndTxId(value1 float64, value2 string) ([]*models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"Value\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToInput(rows)
}

func (db *DB) GetInputByValueAndVout(value1 float64, value2 int) (bool, *models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"Value\" = $1 AND \"Vout\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToInput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryInputByValueAndVout(value1 float64, value2 int) ([]*models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"Value\" = $1 AND \"Vout\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToInput(rows)
}

func (db *DB) GetInputByValueAndCreated(value1 float64, value2 time.Time) (bool, *models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"Value\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToInput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryInputByValueAndCreated(value1 float64, value2 time.Time) ([]*models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"Value\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToInput(rows)
}

func (db *DB) CountInputByCreated(value time.Time) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Input\" WHERE \"Created\" = $1;", value)
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

func (db *DB) GetInputByCreated(value time.Time) (bool, *models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"Created\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToInput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryInputByCreated(value time.Time) ([]*models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"Created\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToInput(rows)
}

func (db *DB) GetInputByCreatedAndTxId(value1 time.Time, value2 string) (bool, *models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"Created\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToInput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryInputByCreatedAndTxId(value1 time.Time, value2 string) ([]*models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"Created\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToInput(rows)
}

func (db *DB) GetInputByCreatedAndVout(value1 time.Time, value2 int) (bool, *models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"Created\" = $1 AND \"Vout\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToInput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryInputByCreatedAndVout(value1 time.Time, value2 int) ([]*models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"Created\" = $1 AND \"Vout\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToInput(rows)
}

func (db *DB) GetInputByCreatedAndValue(value1 time.Time, value2 float64) (bool, *models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"Created\" = $1 AND \"Value\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToInput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryInputByCreatedAndValue(value1 time.Time, value2 float64) ([]*models.Input, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"TxId\", \"Value\", \"Vout\" FROM " + db.dbName + ".\"Input\" WHERE \"Created\" = $1 AND \"Value\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToInput(rows)
}


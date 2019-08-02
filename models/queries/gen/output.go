package database

import (
	"fmt"
	"time"
	"database/sql"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/project"
)

func (db *DB) CreateTableOutput() error {
	q := "CREATE TABLE IF NOT EXISTS " + db.dbName + ".\"Output\" (\"Asset\" STRING NOT NULL, \"Amount\" FLOAT NOT NULL, \"Created\" TIMESTAMP DEFAULT current_timestamp());"
	fmt.Println(q)
	_, err := db.DoExec(q)
	return err
}

func (db *DB) InsertOutput(Amount float64, Asset string) (*models.Output, error) {
	row := &models.Output{}
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Output\" (\"Amount\", \"Asset\") VALUES ($1, $2) RETURNING Amount, Asset;", Amount, Asset).Scan(&row.Amount, &row.Asset, &row.Created); err != nil {
		return nil, err
	}
	return row, nil
}

func (db *DB) DeleteOutput(uid string) error {
	_, err := db.DoQuery("DELETE FROM " + db.dbName + ".\"Output\" WHERE \"UID\" = $1;", uid)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) PutOutput(model *models.Output) error {
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Output\" (\"Amount\", \"Asset\") VALUES ($1, $2) RETURNING \"Amount\", \"Asset\", \"Created\";", model.Amount, model.Asset).Scan(&model.Amount, &model.Asset, &model.Created); err != nil {
		return err
	}
	return nil
}

func (db *DB) ScanToOutput(rows *sql.Rows) ([]*models.Output, error) {
	defer rows.Close()
	results := []*models.Output{}
	for rows.Next() {
		result := &models.Output{}
		err := rows.Scan(&result.Amount, &result.Asset, &result.Created)
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

func (db *DB) CountOutput() (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM " + db.dbName + ".\"Output\";")
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

func (db *DB) QueryOutput() ([]*models.Output, error) {
	rows, err := db.DoQuery("SELECT \"Amount\", \"Asset\", \"Created\" FROM " + db.dbName + ".\"Output\";")
	if err != nil {
		return nil, err
	}
	return db.ScanToOutput(rows)
}

func (db *DB) CountOutputByAsset(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Output\" WHERE \"Asset\" = $1;", value)
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

func (db *DB) GetOutputByAsset(value string) (bool, *models.Output, error) {
	rows, err := db.DoQuery("SELECT \"Amount\", \"Asset\", \"Created\" FROM " + db.dbName + ".\"Output\" WHERE \"Asset\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToOutput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryOutputByAsset(value string) ([]*models.Output, error) {
	rows, err := db.DoQuery("SELECT \"Amount\", \"Asset\", \"Created\" FROM " + db.dbName + ".\"Output\" WHERE \"Asset\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToOutput(rows)
}

func (db *DB) GetOutputByAssetAndAmount(value1 string, value2 float64) (bool, *models.Output, error) {
	rows, err := db.DoQuery("SELECT \"Amount\", \"Asset\", \"Created\" FROM " + db.dbName + ".\"Output\" WHERE \"Asset\" = $1 AND \"Amount\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToOutput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryOutputByAssetAndAmount(value1 string, value2 float64) ([]*models.Output, error) {
	rows, err := db.DoQuery("SELECT \"Amount\", \"Asset\", \"Created\" FROM " + db.dbName + ".\"Output\" WHERE \"Asset\" = $1 AND \"Amount\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToOutput(rows)
}

func (db *DB) GetOutputByAssetAndCreated(value1 string, value2 time.Time) (bool, *models.Output, error) {
	rows, err := db.DoQuery("SELECT \"Amount\", \"Asset\", \"Created\" FROM " + db.dbName + ".\"Output\" WHERE \"Asset\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToOutput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryOutputByAssetAndCreated(value1 string, value2 time.Time) ([]*models.Output, error) {
	rows, err := db.DoQuery("SELECT \"Amount\", \"Asset\", \"Created\" FROM " + db.dbName + ".\"Output\" WHERE \"Asset\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToOutput(rows)
}

func (db *DB) CountOutputByAmount(value float64) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Output\" WHERE \"Amount\" = $1;", value)
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

func (db *DB) GetOutputByAmount(value float64) (bool, *models.Output, error) {
	rows, err := db.DoQuery("SELECT \"Amount\", \"Asset\", \"Created\" FROM " + db.dbName + ".\"Output\" WHERE \"Amount\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToOutput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryOutputByAmount(value float64) ([]*models.Output, error) {
	rows, err := db.DoQuery("SELECT \"Amount\", \"Asset\", \"Created\" FROM " + db.dbName + ".\"Output\" WHERE \"Amount\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToOutput(rows)
}

func (db *DB) GetOutputByAmountAndAsset(value1 float64, value2 string) (bool, *models.Output, error) {
	rows, err := db.DoQuery("SELECT \"Amount\", \"Asset\", \"Created\" FROM " + db.dbName + ".\"Output\" WHERE \"Amount\" = $1 AND \"Asset\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToOutput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryOutputByAmountAndAsset(value1 float64, value2 string) ([]*models.Output, error) {
	rows, err := db.DoQuery("SELECT \"Amount\", \"Asset\", \"Created\" FROM " + db.dbName + ".\"Output\" WHERE \"Amount\" = $1 AND \"Asset\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToOutput(rows)
}

func (db *DB) GetOutputByAmountAndCreated(value1 float64, value2 time.Time) (bool, *models.Output, error) {
	rows, err := db.DoQuery("SELECT \"Amount\", \"Asset\", \"Created\" FROM " + db.dbName + ".\"Output\" WHERE \"Amount\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToOutput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryOutputByAmountAndCreated(value1 float64, value2 time.Time) ([]*models.Output, error) {
	rows, err := db.DoQuery("SELECT \"Amount\", \"Asset\", \"Created\" FROM " + db.dbName + ".\"Output\" WHERE \"Amount\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToOutput(rows)
}

func (db *DB) CountOutputByCreated(value time.Time) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Output\" WHERE \"Created\" = $1;", value)
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

func (db *DB) GetOutputByCreated(value time.Time) (bool, *models.Output, error) {
	rows, err := db.DoQuery("SELECT \"Amount\", \"Asset\", \"Created\" FROM " + db.dbName + ".\"Output\" WHERE \"Created\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToOutput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryOutputByCreated(value time.Time) ([]*models.Output, error) {
	rows, err := db.DoQuery("SELECT \"Amount\", \"Asset\", \"Created\" FROM " + db.dbName + ".\"Output\" WHERE \"Created\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToOutput(rows)
}

func (db *DB) GetOutputByCreatedAndAsset(value1 time.Time, value2 string) (bool, *models.Output, error) {
	rows, err := db.DoQuery("SELECT \"Amount\", \"Asset\", \"Created\" FROM " + db.dbName + ".\"Output\" WHERE \"Created\" = $1 AND \"Asset\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToOutput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryOutputByCreatedAndAsset(value1 time.Time, value2 string) ([]*models.Output, error) {
	rows, err := db.DoQuery("SELECT \"Amount\", \"Asset\", \"Created\" FROM " + db.dbName + ".\"Output\" WHERE \"Created\" = $1 AND \"Asset\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToOutput(rows)
}

func (db *DB) GetOutputByCreatedAndAmount(value1 time.Time, value2 float64) (bool, *models.Output, error) {
	rows, err := db.DoQuery("SELECT \"Amount\", \"Asset\", \"Created\" FROM " + db.dbName + ".\"Output\" WHERE \"Created\" = $1 AND \"Amount\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToOutput(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryOutputByCreatedAndAmount(value1 time.Time, value2 float64) ([]*models.Output, error) {
	rows, err := db.DoQuery("SELECT \"Amount\", \"Asset\", \"Created\" FROM " + db.dbName + ".\"Output\" WHERE \"Created\" = $1 AND \"Amount\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToOutput(rows)
}


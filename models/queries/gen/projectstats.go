package database

import (
	"fmt"
	"time"
	"database/sql"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/project"
)

func (db *DB) CreateTableProjectStats() error {
	q := "CREATE TABLE IF NOT EXISTS " + db.dbName + ".\"ProjectStats\" (\"BytesConsumed\" int64 NOT NULL, \"Salt\" UUID DEFAULT gen_random_uuid(), \"Created\" TIMESTAMP DEFAULT current_timestamp(), \"Project\" STRING NOT NULL, \"Users\" INT NOT NULL, \"Streams\" INT NOT NULL, \"Exchanges\" INT NOT NULL, \"Wallets\" INT NOT NULL, \"UID\" UUID PRIMARY KEY DEFAULT gen_random_uuid(), \"Title\" STRING NOT NULL, \"TxCount\" INT NOT NULL, \"Currencies\" INT NOT NULL);"
	fmt.Println(q)
	_, err := db.DoExec(q)
	return err
}

func (db *DB) InsertProjectStats(BytesConsumed int64, Currencies int, Exchanges int, Project string, Streams int, Title string, TxCount int, Users int, Wallets int) (*models.ProjectStats, error) {
	row := &models.ProjectStats{}
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"ProjectStats\" (\"BytesConsumed\", \"Currencies\", \"Exchanges\", \"Project\", \"Streams\", \"Title\", \"TxCount\", \"Users\", \"Wallets\") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING BytesConsumed, Currencies, Exchanges, Project, Streams, Title, TxCount, Users, Wallets;", BytesConsumed, Currencies, Exchanges, Project, Streams, Title, TxCount, Users, Wallets).Scan(&row.BytesConsumed, &row.Created, &row.Currencies, &row.Exchanges, &row.Project, &row.Salt, &row.Streams, &row.Title, &row.TxCount, &row.UID, &row.Users, &row.Wallets); err != nil {
		return nil, err
	}
	return row, nil
}

func (db *DB) DeleteProjectStats(uid string) error {
	_, err := db.DoQuery("DELETE FROM " + db.dbName + ".\"ProjectStats\" WHERE \"UID\" = $1;", uid)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) PutProjectStats(model *models.ProjectStats) error {
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"ProjectStats\" (\"BytesConsumed\", \"Currencies\", \"Exchanges\", \"Project\", \"Streams\", \"Title\", \"TxCount\", \"Users\", \"Wallets\") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\";", model.BytesConsumed, model.Currencies, model.Exchanges, model.Project, model.Streams, model.Title, model.TxCount, model.Users, model.Wallets).Scan(&model.BytesConsumed, &model.Created, &model.Currencies, &model.Exchanges, &model.Project, &model.Salt, &model.Streams, &model.Title, &model.TxCount, &model.UID, &model.Users, &model.Wallets); err != nil {
		return err
	}
	return nil
}

func (db *DB) ScanToProjectStats(rows *sql.Rows) ([]*models.ProjectStats, error) {
	defer rows.Close()
	results := []*models.ProjectStats{}
	for rows.Next() {
		result := &models.ProjectStats{}
		err := rows.Scan(&result.BytesConsumed, &result.Created, &result.Currencies, &result.Exchanges, &result.Project, &result.Salt, &result.Streams, &result.Title, &result.TxCount, &result.UID, &result.Users, &result.Wallets)
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

func (db *DB) CountProjectStats() (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM " + db.dbName + ".\"ProjectStats\";")
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

func (db *DB) QueryProjectStats() ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\";")
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) CountProjectStatsByBytesConsumed(value int64) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"ProjectStats\" WHERE \"BytesConsumed\" = $1;", value)
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

func (db *DB) GetProjectStatsByBytesConsumed(value int64) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"BytesConsumed\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByBytesConsumed(value int64) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"BytesConsumed\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByBytesConsumedAndSalt(value1 int64, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"BytesConsumed\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByBytesConsumedAndSalt(value1 int64, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"BytesConsumed\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByBytesConsumedAndCreated(value1 int64, value2 time.Time) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"BytesConsumed\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByBytesConsumedAndCreated(value1 int64, value2 time.Time) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"BytesConsumed\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByBytesConsumedAndProject(value1 int64, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"BytesConsumed\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByBytesConsumedAndProject(value1 int64, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"BytesConsumed\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByBytesConsumedAndUsers(value1 int64, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"BytesConsumed\" = $1 AND \"Users\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByBytesConsumedAndUsers(value1 int64, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"BytesConsumed\" = $1 AND \"Users\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByBytesConsumedAndStreams(value1 int64, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"BytesConsumed\" = $1 AND \"Streams\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByBytesConsumedAndStreams(value1 int64, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"BytesConsumed\" = $1 AND \"Streams\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByBytesConsumedAndExchanges(value1 int64, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"BytesConsumed\" = $1 AND \"Exchanges\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByBytesConsumedAndExchanges(value1 int64, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"BytesConsumed\" = $1 AND \"Exchanges\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByBytesConsumedAndWallets(value1 int64, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"BytesConsumed\" = $1 AND \"Wallets\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByBytesConsumedAndWallets(value1 int64, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"BytesConsumed\" = $1 AND \"Wallets\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByBytesConsumedAndUID(value1 int64, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"BytesConsumed\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByBytesConsumedAndUID(value1 int64, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"BytesConsumed\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByBytesConsumedAndTitle(value1 int64, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"BytesConsumed\" = $1 AND \"Title\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByBytesConsumedAndTitle(value1 int64, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"BytesConsumed\" = $1 AND \"Title\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByBytesConsumedAndTxCount(value1 int64, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"BytesConsumed\" = $1 AND \"TxCount\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByBytesConsumedAndTxCount(value1 int64, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"BytesConsumed\" = $1 AND \"TxCount\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByBytesConsumedAndCurrencies(value1 int64, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"BytesConsumed\" = $1 AND \"Currencies\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByBytesConsumedAndCurrencies(value1 int64, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"BytesConsumed\" = $1 AND \"Currencies\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) CountProjectStatsBySalt(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"ProjectStats\" WHERE \"Salt\" = $1;", value)
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

func (db *DB) GetProjectStatsBySalt(value string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Salt\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsBySalt(value string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Salt\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsBySaltAndBytesConsumed(value1 string, value2 int64) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Salt\" = $1 AND \"BytesConsumed\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsBySaltAndBytesConsumed(value1 string, value2 int64) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Salt\" = $1 AND \"BytesConsumed\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsBySaltAndCreated(value1 string, value2 time.Time) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Salt\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsBySaltAndCreated(value1 string, value2 time.Time) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Salt\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsBySaltAndProject(value1 string, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Salt\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsBySaltAndProject(value1 string, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Salt\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsBySaltAndUsers(value1 string, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Salt\" = $1 AND \"Users\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsBySaltAndUsers(value1 string, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Salt\" = $1 AND \"Users\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsBySaltAndStreams(value1 string, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Salt\" = $1 AND \"Streams\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsBySaltAndStreams(value1 string, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Salt\" = $1 AND \"Streams\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsBySaltAndExchanges(value1 string, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Salt\" = $1 AND \"Exchanges\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsBySaltAndExchanges(value1 string, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Salt\" = $1 AND \"Exchanges\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsBySaltAndWallets(value1 string, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Salt\" = $1 AND \"Wallets\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsBySaltAndWallets(value1 string, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Salt\" = $1 AND \"Wallets\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsBySaltAndUID(value1 string, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Salt\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsBySaltAndUID(value1 string, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Salt\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsBySaltAndTitle(value1 string, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Salt\" = $1 AND \"Title\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsBySaltAndTitle(value1 string, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Salt\" = $1 AND \"Title\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsBySaltAndTxCount(value1 string, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Salt\" = $1 AND \"TxCount\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsBySaltAndTxCount(value1 string, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Salt\" = $1 AND \"TxCount\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsBySaltAndCurrencies(value1 string, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Salt\" = $1 AND \"Currencies\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsBySaltAndCurrencies(value1 string, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Salt\" = $1 AND \"Currencies\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) CountProjectStatsByCreated(value time.Time) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"ProjectStats\" WHERE \"Created\" = $1;", value)
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

func (db *DB) GetProjectStatsByCreated(value time.Time) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Created\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByCreated(value time.Time) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Created\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByCreatedAndBytesConsumed(value1 time.Time, value2 int64) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Created\" = $1 AND \"BytesConsumed\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByCreatedAndBytesConsumed(value1 time.Time, value2 int64) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Created\" = $1 AND \"BytesConsumed\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByCreatedAndSalt(value1 time.Time, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Created\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByCreatedAndSalt(value1 time.Time, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Created\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByCreatedAndProject(value1 time.Time, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Created\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByCreatedAndProject(value1 time.Time, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Created\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByCreatedAndUsers(value1 time.Time, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Created\" = $1 AND \"Users\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByCreatedAndUsers(value1 time.Time, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Created\" = $1 AND \"Users\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByCreatedAndStreams(value1 time.Time, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Created\" = $1 AND \"Streams\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByCreatedAndStreams(value1 time.Time, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Created\" = $1 AND \"Streams\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByCreatedAndExchanges(value1 time.Time, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Created\" = $1 AND \"Exchanges\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByCreatedAndExchanges(value1 time.Time, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Created\" = $1 AND \"Exchanges\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByCreatedAndWallets(value1 time.Time, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Created\" = $1 AND \"Wallets\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByCreatedAndWallets(value1 time.Time, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Created\" = $1 AND \"Wallets\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByCreatedAndUID(value1 time.Time, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Created\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByCreatedAndUID(value1 time.Time, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Created\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByCreatedAndTitle(value1 time.Time, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Created\" = $1 AND \"Title\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByCreatedAndTitle(value1 time.Time, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Created\" = $1 AND \"Title\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByCreatedAndTxCount(value1 time.Time, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Created\" = $1 AND \"TxCount\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByCreatedAndTxCount(value1 time.Time, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Created\" = $1 AND \"TxCount\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByCreatedAndCurrencies(value1 time.Time, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Created\" = $1 AND \"Currencies\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByCreatedAndCurrencies(value1 time.Time, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Created\" = $1 AND \"Currencies\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) CountProjectStatsByProject(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"ProjectStats\" WHERE \"Project\" = $1;", value)
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

func (db *DB) GetProjectStatsByProject(value string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Project\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByProject(value string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Project\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByProjectAndBytesConsumed(value1 string, value2 int64) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Project\" = $1 AND \"BytesConsumed\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByProjectAndBytesConsumed(value1 string, value2 int64) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Project\" = $1 AND \"BytesConsumed\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByProjectAndSalt(value1 string, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Project\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByProjectAndSalt(value1 string, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Project\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByProjectAndCreated(value1 string, value2 time.Time) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Project\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByProjectAndCreated(value1 string, value2 time.Time) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Project\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByProjectAndUsers(value1 string, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Project\" = $1 AND \"Users\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByProjectAndUsers(value1 string, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Project\" = $1 AND \"Users\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByProjectAndStreams(value1 string, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Project\" = $1 AND \"Streams\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByProjectAndStreams(value1 string, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Project\" = $1 AND \"Streams\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByProjectAndExchanges(value1 string, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Project\" = $1 AND \"Exchanges\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByProjectAndExchanges(value1 string, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Project\" = $1 AND \"Exchanges\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByProjectAndWallets(value1 string, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Project\" = $1 AND \"Wallets\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByProjectAndWallets(value1 string, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Project\" = $1 AND \"Wallets\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByProjectAndUID(value1 string, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Project\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByProjectAndUID(value1 string, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Project\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByProjectAndTitle(value1 string, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Project\" = $1 AND \"Title\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByProjectAndTitle(value1 string, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Project\" = $1 AND \"Title\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByProjectAndTxCount(value1 string, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Project\" = $1 AND \"TxCount\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByProjectAndTxCount(value1 string, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Project\" = $1 AND \"TxCount\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByProjectAndCurrencies(value1 string, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Project\" = $1 AND \"Currencies\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByProjectAndCurrencies(value1 string, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Project\" = $1 AND \"Currencies\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) CountProjectStatsByUsers(value int) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"ProjectStats\" WHERE \"Users\" = $1;", value)
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

func (db *DB) GetProjectStatsByUsers(value int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Users\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByUsers(value int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Users\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByUsersAndBytesConsumed(value1 int, value2 int64) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Users\" = $1 AND \"BytesConsumed\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByUsersAndBytesConsumed(value1 int, value2 int64) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Users\" = $1 AND \"BytesConsumed\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByUsersAndSalt(value1 int, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Users\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByUsersAndSalt(value1 int, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Users\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByUsersAndCreated(value1 int, value2 time.Time) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Users\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByUsersAndCreated(value1 int, value2 time.Time) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Users\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByUsersAndProject(value1 int, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Users\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByUsersAndProject(value1 int, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Users\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByUsersAndStreams(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Users\" = $1 AND \"Streams\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByUsersAndStreams(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Users\" = $1 AND \"Streams\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByUsersAndExchanges(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Users\" = $1 AND \"Exchanges\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByUsersAndExchanges(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Users\" = $1 AND \"Exchanges\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByUsersAndWallets(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Users\" = $1 AND \"Wallets\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByUsersAndWallets(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Users\" = $1 AND \"Wallets\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByUsersAndUID(value1 int, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Users\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByUsersAndUID(value1 int, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Users\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByUsersAndTitle(value1 int, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Users\" = $1 AND \"Title\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByUsersAndTitle(value1 int, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Users\" = $1 AND \"Title\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByUsersAndTxCount(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Users\" = $1 AND \"TxCount\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByUsersAndTxCount(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Users\" = $1 AND \"TxCount\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByUsersAndCurrencies(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Users\" = $1 AND \"Currencies\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByUsersAndCurrencies(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Users\" = $1 AND \"Currencies\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) CountProjectStatsByStreams(value int) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"ProjectStats\" WHERE \"Streams\" = $1;", value)
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

func (db *DB) GetProjectStatsByStreams(value int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Streams\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByStreams(value int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Streams\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByStreamsAndBytesConsumed(value1 int, value2 int64) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Streams\" = $1 AND \"BytesConsumed\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByStreamsAndBytesConsumed(value1 int, value2 int64) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Streams\" = $1 AND \"BytesConsumed\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByStreamsAndSalt(value1 int, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Streams\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByStreamsAndSalt(value1 int, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Streams\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByStreamsAndCreated(value1 int, value2 time.Time) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Streams\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByStreamsAndCreated(value1 int, value2 time.Time) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Streams\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByStreamsAndProject(value1 int, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Streams\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByStreamsAndProject(value1 int, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Streams\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByStreamsAndUsers(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Streams\" = $1 AND \"Users\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByStreamsAndUsers(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Streams\" = $1 AND \"Users\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByStreamsAndExchanges(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Streams\" = $1 AND \"Exchanges\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByStreamsAndExchanges(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Streams\" = $1 AND \"Exchanges\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByStreamsAndWallets(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Streams\" = $1 AND \"Wallets\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByStreamsAndWallets(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Streams\" = $1 AND \"Wallets\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByStreamsAndUID(value1 int, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Streams\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByStreamsAndUID(value1 int, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Streams\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByStreamsAndTitle(value1 int, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Streams\" = $1 AND \"Title\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByStreamsAndTitle(value1 int, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Streams\" = $1 AND \"Title\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByStreamsAndTxCount(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Streams\" = $1 AND \"TxCount\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByStreamsAndTxCount(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Streams\" = $1 AND \"TxCount\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByStreamsAndCurrencies(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Streams\" = $1 AND \"Currencies\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByStreamsAndCurrencies(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Streams\" = $1 AND \"Currencies\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) CountProjectStatsByExchanges(value int) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"ProjectStats\" WHERE \"Exchanges\" = $1;", value)
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

func (db *DB) GetProjectStatsByExchanges(value int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Exchanges\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByExchanges(value int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Exchanges\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByExchangesAndBytesConsumed(value1 int, value2 int64) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Exchanges\" = $1 AND \"BytesConsumed\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByExchangesAndBytesConsumed(value1 int, value2 int64) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Exchanges\" = $1 AND \"BytesConsumed\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByExchangesAndSalt(value1 int, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Exchanges\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByExchangesAndSalt(value1 int, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Exchanges\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByExchangesAndCreated(value1 int, value2 time.Time) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Exchanges\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByExchangesAndCreated(value1 int, value2 time.Time) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Exchanges\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByExchangesAndProject(value1 int, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Exchanges\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByExchangesAndProject(value1 int, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Exchanges\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByExchangesAndUsers(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Exchanges\" = $1 AND \"Users\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByExchangesAndUsers(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Exchanges\" = $1 AND \"Users\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByExchangesAndStreams(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Exchanges\" = $1 AND \"Streams\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByExchangesAndStreams(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Exchanges\" = $1 AND \"Streams\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByExchangesAndWallets(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Exchanges\" = $1 AND \"Wallets\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByExchangesAndWallets(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Exchanges\" = $1 AND \"Wallets\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByExchangesAndUID(value1 int, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Exchanges\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByExchangesAndUID(value1 int, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Exchanges\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByExchangesAndTitle(value1 int, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Exchanges\" = $1 AND \"Title\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByExchangesAndTitle(value1 int, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Exchanges\" = $1 AND \"Title\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByExchangesAndTxCount(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Exchanges\" = $1 AND \"TxCount\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByExchangesAndTxCount(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Exchanges\" = $1 AND \"TxCount\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByExchangesAndCurrencies(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Exchanges\" = $1 AND \"Currencies\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByExchangesAndCurrencies(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Exchanges\" = $1 AND \"Currencies\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) CountProjectStatsByWallets(value int) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"ProjectStats\" WHERE \"Wallets\" = $1;", value)
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

func (db *DB) GetProjectStatsByWallets(value int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Wallets\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByWallets(value int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Wallets\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByWalletsAndBytesConsumed(value1 int, value2 int64) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Wallets\" = $1 AND \"BytesConsumed\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByWalletsAndBytesConsumed(value1 int, value2 int64) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Wallets\" = $1 AND \"BytesConsumed\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByWalletsAndSalt(value1 int, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Wallets\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByWalletsAndSalt(value1 int, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Wallets\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByWalletsAndCreated(value1 int, value2 time.Time) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Wallets\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByWalletsAndCreated(value1 int, value2 time.Time) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Wallets\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByWalletsAndProject(value1 int, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Wallets\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByWalletsAndProject(value1 int, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Wallets\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByWalletsAndUsers(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Wallets\" = $1 AND \"Users\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByWalletsAndUsers(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Wallets\" = $1 AND \"Users\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByWalletsAndStreams(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Wallets\" = $1 AND \"Streams\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByWalletsAndStreams(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Wallets\" = $1 AND \"Streams\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByWalletsAndExchanges(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Wallets\" = $1 AND \"Exchanges\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByWalletsAndExchanges(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Wallets\" = $1 AND \"Exchanges\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByWalletsAndUID(value1 int, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Wallets\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByWalletsAndUID(value1 int, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Wallets\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByWalletsAndTitle(value1 int, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Wallets\" = $1 AND \"Title\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByWalletsAndTitle(value1 int, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Wallets\" = $1 AND \"Title\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByWalletsAndTxCount(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Wallets\" = $1 AND \"TxCount\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByWalletsAndTxCount(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Wallets\" = $1 AND \"TxCount\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByWalletsAndCurrencies(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Wallets\" = $1 AND \"Currencies\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByWalletsAndCurrencies(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Wallets\" = $1 AND \"Currencies\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) CountProjectStatsByUID(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"ProjectStats\" WHERE \"UID\" = $1;", value)
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

func (db *DB) GetProjectStatsByUID(value string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"UID\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByUID(value string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"UID\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) CountProjectStatsByTitle(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"ProjectStats\" WHERE \"Title\" = $1;", value)
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

func (db *DB) GetProjectStatsByTitle(value string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Title\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByTitle(value string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Title\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByTitleAndBytesConsumed(value1 string, value2 int64) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Title\" = $1 AND \"BytesConsumed\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByTitleAndBytesConsumed(value1 string, value2 int64) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Title\" = $1 AND \"BytesConsumed\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByTitleAndSalt(value1 string, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Title\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByTitleAndSalt(value1 string, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Title\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByTitleAndCreated(value1 string, value2 time.Time) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Title\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByTitleAndCreated(value1 string, value2 time.Time) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Title\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByTitleAndProject(value1 string, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Title\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByTitleAndProject(value1 string, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Title\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByTitleAndUsers(value1 string, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Title\" = $1 AND \"Users\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByTitleAndUsers(value1 string, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Title\" = $1 AND \"Users\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByTitleAndStreams(value1 string, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Title\" = $1 AND \"Streams\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByTitleAndStreams(value1 string, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Title\" = $1 AND \"Streams\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByTitleAndExchanges(value1 string, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Title\" = $1 AND \"Exchanges\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByTitleAndExchanges(value1 string, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Title\" = $1 AND \"Exchanges\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByTitleAndWallets(value1 string, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Title\" = $1 AND \"Wallets\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByTitleAndWallets(value1 string, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Title\" = $1 AND \"Wallets\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByTitleAndUID(value1 string, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Title\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByTitleAndUID(value1 string, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Title\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByTitleAndTxCount(value1 string, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Title\" = $1 AND \"TxCount\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByTitleAndTxCount(value1 string, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Title\" = $1 AND \"TxCount\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByTitleAndCurrencies(value1 string, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Title\" = $1 AND \"Currencies\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByTitleAndCurrencies(value1 string, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Title\" = $1 AND \"Currencies\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) CountProjectStatsByTxCount(value int) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"ProjectStats\" WHERE \"TxCount\" = $1;", value)
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

func (db *DB) GetProjectStatsByTxCount(value int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"TxCount\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByTxCount(value int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"TxCount\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByTxCountAndBytesConsumed(value1 int, value2 int64) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"TxCount\" = $1 AND \"BytesConsumed\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByTxCountAndBytesConsumed(value1 int, value2 int64) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"TxCount\" = $1 AND \"BytesConsumed\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByTxCountAndSalt(value1 int, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"TxCount\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByTxCountAndSalt(value1 int, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"TxCount\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByTxCountAndCreated(value1 int, value2 time.Time) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"TxCount\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByTxCountAndCreated(value1 int, value2 time.Time) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"TxCount\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByTxCountAndProject(value1 int, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"TxCount\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByTxCountAndProject(value1 int, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"TxCount\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByTxCountAndUsers(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"TxCount\" = $1 AND \"Users\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByTxCountAndUsers(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"TxCount\" = $1 AND \"Users\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByTxCountAndStreams(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"TxCount\" = $1 AND \"Streams\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByTxCountAndStreams(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"TxCount\" = $1 AND \"Streams\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByTxCountAndExchanges(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"TxCount\" = $1 AND \"Exchanges\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByTxCountAndExchanges(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"TxCount\" = $1 AND \"Exchanges\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByTxCountAndWallets(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"TxCount\" = $1 AND \"Wallets\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByTxCountAndWallets(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"TxCount\" = $1 AND \"Wallets\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByTxCountAndUID(value1 int, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"TxCount\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByTxCountAndUID(value1 int, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"TxCount\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByTxCountAndTitle(value1 int, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"TxCount\" = $1 AND \"Title\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByTxCountAndTitle(value1 int, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"TxCount\" = $1 AND \"Title\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByTxCountAndCurrencies(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"TxCount\" = $1 AND \"Currencies\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByTxCountAndCurrencies(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"TxCount\" = $1 AND \"Currencies\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) CountProjectStatsByCurrencies(value int) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"ProjectStats\" WHERE \"Currencies\" = $1;", value)
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

func (db *DB) GetProjectStatsByCurrencies(value int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Currencies\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByCurrencies(value int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Currencies\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByCurrenciesAndBytesConsumed(value1 int, value2 int64) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Currencies\" = $1 AND \"BytesConsumed\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByCurrenciesAndBytesConsumed(value1 int, value2 int64) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Currencies\" = $1 AND \"BytesConsumed\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByCurrenciesAndSalt(value1 int, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Currencies\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByCurrenciesAndSalt(value1 int, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Currencies\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByCurrenciesAndCreated(value1 int, value2 time.Time) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Currencies\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByCurrenciesAndCreated(value1 int, value2 time.Time) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Currencies\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByCurrenciesAndProject(value1 int, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Currencies\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByCurrenciesAndProject(value1 int, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Currencies\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByCurrenciesAndUsers(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Currencies\" = $1 AND \"Users\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByCurrenciesAndUsers(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Currencies\" = $1 AND \"Users\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByCurrenciesAndStreams(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Currencies\" = $1 AND \"Streams\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByCurrenciesAndStreams(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Currencies\" = $1 AND \"Streams\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByCurrenciesAndExchanges(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Currencies\" = $1 AND \"Exchanges\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByCurrenciesAndExchanges(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Currencies\" = $1 AND \"Exchanges\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByCurrenciesAndWallets(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Currencies\" = $1 AND \"Wallets\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByCurrenciesAndWallets(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Currencies\" = $1 AND \"Wallets\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByCurrenciesAndUID(value1 int, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Currencies\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByCurrenciesAndUID(value1 int, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Currencies\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByCurrenciesAndTitle(value1 int, value2 string) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Currencies\" = $1 AND \"Title\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByCurrenciesAndTitle(value1 int, value2 string) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Currencies\" = $1 AND \"Title\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}

func (db *DB) GetProjectStatsByCurrenciesAndTxCount(value1 int, value2 int) (bool, *models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Currencies\" = $1 AND \"TxCount\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProjectStats(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectStatsByCurrenciesAndTxCount(value1 int, value2 int) ([]*models.ProjectStats, error) {
	rows, err := db.DoQuery("SELECT \"BytesConsumed\", \"Created\", \"Currencies\", \"Exchanges\", \"Project\", \"Salt\", \"Streams\", \"Title\", \"TxCount\", \"UID\", \"Users\", \"Wallets\" FROM " + db.dbName + ".\"ProjectStats\" WHERE \"Currencies\" = $1 AND \"TxCount\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProjectStats(rows)
}


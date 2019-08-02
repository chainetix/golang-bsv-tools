package database

import (
	"fmt"
	"time"
	"database/sql"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/project"
)

func (db *DB) CreateTableStreamItem() error {
	q := "CREATE TABLE IF NOT EXISTS " + db.dbName + ".\"StreamItem\" (\"Blockhash\" STRING NOT NULL, \"Data\" STRING NOT NULL, \"Keys\" STRING ARRAY NOT NULL, \"Vout\" INT NOT NULL, \"Blocktime\" INT NOT NULL, \"Blockindex\" INT NOT NULL, \"Confirmations\" INT NOT NULL, \"Publishers\" STRING ARRAY NOT NULL, \"Time\" INT NOT NULL, \"Timereceived\" INT NOT NULL, \"UID\" UUID PRIMARY KEY DEFAULT gen_random_uuid(), \"Salt\" UUID DEFAULT gen_random_uuid(), \"Created\" TIMESTAMP DEFAULT current_timestamp(), \"Txid\" STRING NOT NULL, \"Valid\" BOOL NOT NULL);"
	fmt.Println(q)
	_, err := db.DoExec(q)
	return err
}

func (db *DB) InsertStreamItem(Blockhash string, Blockindex int, Blocktime int, Confirmations int, Data string, Keys []string, Publishers []string, Time int, Timereceived int, Txid string, Valid bool, Vout int) (*models.StreamItem, error) {
	row := &models.StreamItem{}
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"StreamItem\" (\"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Data\", \"Keys\", \"Publishers\", \"Time\", \"Timereceived\", \"Txid\", \"Valid\", \"Vout\") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING Blockhash, Blockindex, Blocktime, Confirmations, Data, Keys, Publishers, Time, Timereceived, Txid, Valid, Vout;", Blockhash, Blockindex, Blocktime, Confirmations, Data, Keys, Publishers, Time, Timereceived, Txid, Valid, Vout).Scan(&row.Blockhash, &row.Blockindex, &row.Blocktime, &row.Confirmations, &row.Created, &row.Data, &row.Keys, &row.Publishers, &row.Salt, &row.Time, &row.Timereceived, &row.Txid, &row.UID, &row.Valid, &row.Vout); err != nil {
		return nil, err
	}
	return row, nil
}

func (db *DB) DeleteStreamItem(uid string) error {
	_, err := db.DoQuery("DELETE FROM " + db.dbName + ".\"StreamItem\" WHERE \"UID\" = $1;", uid)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) PutStreamItem(model *models.StreamItem) error {
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"StreamItem\" (\"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Data\", \"Keys\", \"Publishers\", \"Time\", \"Timereceived\", \"Txid\", \"Valid\", \"Vout\") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\";", model.Blockhash, model.Blockindex, model.Blocktime, model.Confirmations, model.Data, model.Keys, model.Publishers, model.Time, model.Timereceived, model.Txid, model.Valid, model.Vout).Scan(&model.Blockhash, &model.Blockindex, &model.Blocktime, &model.Confirmations, &model.Created, &model.Data, &model.Keys, &model.Publishers, &model.Salt, &model.Time, &model.Timereceived, &model.Txid, &model.UID, &model.Valid, &model.Vout); err != nil {
		return err
	}
	return nil
}

func (db *DB) ScanToStreamItem(rows *sql.Rows) ([]*models.StreamItem, error) {
	defer rows.Close()
	results := []*models.StreamItem{}
	for rows.Next() {
		result := &models.StreamItem{}
		err := rows.Scan(&result.Blockhash, &result.Blockindex, &result.Blocktime, &result.Confirmations, &result.Created, &result.Data, &result.Keys, &result.Publishers, &result.Salt, &result.Time, &result.Timereceived, &result.Txid, &result.UID, &result.Valid, &result.Vout)
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

func (db *DB) CountStreamItem() (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM " + db.dbName + ".\"StreamItem\";")
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

func (db *DB) QueryStreamItem() ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\";")
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) CountStreamItemByBlockhash(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"StreamItem\" WHERE \"Blockhash\" = $1;", value)
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

func (db *DB) GetStreamItemByBlockhash(value string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockhash(value string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockhashAndData(value1 string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Data\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockhashAndData(value1 string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Data\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockhashAndKeys(value1 string, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Keys\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockhashAndKeys(value1 string, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Keys\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockhashAndVout(value1 string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Vout\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockhashAndVout(value1 string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Vout\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockhashAndBlocktime(value1 string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Blocktime\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockhashAndBlocktime(value1 string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Blocktime\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockhashAndBlockindex(value1 string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Blockindex\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockhashAndBlockindex(value1 string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Blockindex\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockhashAndConfirmations(value1 string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Confirmations\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockhashAndConfirmations(value1 string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Confirmations\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockhashAndPublishers(value1 string, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Publishers\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockhashAndPublishers(value1 string, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Publishers\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockhashAndTime(value1 string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockhashAndTime(value1 string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockhashAndTimereceived(value1 string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Timereceived\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockhashAndTimereceived(value1 string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Timereceived\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockhashAndUID(value1 string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockhashAndUID(value1 string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockhashAndSalt(value1 string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockhashAndSalt(value1 string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockhashAndCreated(value1 string, value2 time.Time) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockhashAndCreated(value1 string, value2 time.Time) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockhashAndTxid(value1 string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Txid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockhashAndTxid(value1 string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Txid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockhashAndValid(value1 string, value2 bool) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Valid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockhashAndValid(value1 string, value2 bool) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockhash\" = $1 AND \"Valid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) CountStreamItemByData(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"StreamItem\" WHERE \"Data\" = $1;", value)
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

func (db *DB) GetStreamItemByData(value string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByData(value string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByDataAndBlockhash(value1 string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Blockhash\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByDataAndBlockhash(value1 string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Blockhash\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByDataAndKeys(value1 string, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Keys\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByDataAndKeys(value1 string, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Keys\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByDataAndVout(value1 string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Vout\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByDataAndVout(value1 string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Vout\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByDataAndBlocktime(value1 string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Blocktime\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByDataAndBlocktime(value1 string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Blocktime\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByDataAndBlockindex(value1 string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Blockindex\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByDataAndBlockindex(value1 string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Blockindex\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByDataAndConfirmations(value1 string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Confirmations\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByDataAndConfirmations(value1 string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Confirmations\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByDataAndPublishers(value1 string, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Publishers\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByDataAndPublishers(value1 string, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Publishers\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByDataAndTime(value1 string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByDataAndTime(value1 string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByDataAndTimereceived(value1 string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Timereceived\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByDataAndTimereceived(value1 string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Timereceived\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByDataAndUID(value1 string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByDataAndUID(value1 string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByDataAndSalt(value1 string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByDataAndSalt(value1 string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByDataAndCreated(value1 string, value2 time.Time) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByDataAndCreated(value1 string, value2 time.Time) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByDataAndTxid(value1 string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Txid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByDataAndTxid(value1 string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Txid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByDataAndValid(value1 string, value2 bool) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Valid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByDataAndValid(value1 string, value2 bool) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Data\" = $1 AND \"Valid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) CountStreamItemByKeys(value []string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"StreamItem\" WHERE \"Keys\" = $1;", value)
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

func (db *DB) GetStreamItemByKeys(value []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByKeys(value []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByKeysAndBlockhash(value1 []string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Blockhash\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByKeysAndBlockhash(value1 []string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Blockhash\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByKeysAndData(value1 []string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Data\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByKeysAndData(value1 []string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Data\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByKeysAndVout(value1 []string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Vout\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByKeysAndVout(value1 []string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Vout\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByKeysAndBlocktime(value1 []string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Blocktime\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByKeysAndBlocktime(value1 []string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Blocktime\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByKeysAndBlockindex(value1 []string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Blockindex\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByKeysAndBlockindex(value1 []string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Blockindex\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByKeysAndConfirmations(value1 []string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Confirmations\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByKeysAndConfirmations(value1 []string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Confirmations\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByKeysAndPublishers(value1 []string, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Publishers\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByKeysAndPublishers(value1 []string, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Publishers\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByKeysAndTime(value1 []string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByKeysAndTime(value1 []string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByKeysAndTimereceived(value1 []string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Timereceived\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByKeysAndTimereceived(value1 []string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Timereceived\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByKeysAndUID(value1 []string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByKeysAndUID(value1 []string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByKeysAndSalt(value1 []string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByKeysAndSalt(value1 []string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByKeysAndCreated(value1 []string, value2 time.Time) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByKeysAndCreated(value1 []string, value2 time.Time) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByKeysAndTxid(value1 []string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Txid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByKeysAndTxid(value1 []string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Txid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByKeysAndValid(value1 []string, value2 bool) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Valid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByKeysAndValid(value1 []string, value2 bool) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Keys\" = $1 AND \"Valid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) CountStreamItemByVout(value int) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"StreamItem\" WHERE \"Vout\" = $1;", value)
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

func (db *DB) GetStreamItemByVout(value int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByVout(value int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByVoutAndBlockhash(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Blockhash\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByVoutAndBlockhash(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Blockhash\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByVoutAndData(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Data\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByVoutAndData(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Data\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByVoutAndKeys(value1 int, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Keys\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByVoutAndKeys(value1 int, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Keys\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByVoutAndBlocktime(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Blocktime\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByVoutAndBlocktime(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Blocktime\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByVoutAndBlockindex(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Blockindex\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByVoutAndBlockindex(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Blockindex\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByVoutAndConfirmations(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Confirmations\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByVoutAndConfirmations(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Confirmations\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByVoutAndPublishers(value1 int, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Publishers\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByVoutAndPublishers(value1 int, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Publishers\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByVoutAndTime(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByVoutAndTime(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByVoutAndTimereceived(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Timereceived\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByVoutAndTimereceived(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Timereceived\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByVoutAndUID(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByVoutAndUID(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByVoutAndSalt(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByVoutAndSalt(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByVoutAndCreated(value1 int, value2 time.Time) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByVoutAndCreated(value1 int, value2 time.Time) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByVoutAndTxid(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Txid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByVoutAndTxid(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Txid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByVoutAndValid(value1 int, value2 bool) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Valid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByVoutAndValid(value1 int, value2 bool) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Vout\" = $1 AND \"Valid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) CountStreamItemByBlocktime(value int) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"StreamItem\" WHERE \"Blocktime\" = $1;", value)
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

func (db *DB) GetStreamItemByBlocktime(value int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlocktime(value int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlocktimeAndBlockhash(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Blockhash\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlocktimeAndBlockhash(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Blockhash\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlocktimeAndData(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Data\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlocktimeAndData(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Data\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlocktimeAndKeys(value1 int, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Keys\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlocktimeAndKeys(value1 int, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Keys\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlocktimeAndVout(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Vout\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlocktimeAndVout(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Vout\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlocktimeAndBlockindex(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Blockindex\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlocktimeAndBlockindex(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Blockindex\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlocktimeAndConfirmations(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Confirmations\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlocktimeAndConfirmations(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Confirmations\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlocktimeAndPublishers(value1 int, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Publishers\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlocktimeAndPublishers(value1 int, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Publishers\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlocktimeAndTime(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlocktimeAndTime(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlocktimeAndTimereceived(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Timereceived\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlocktimeAndTimereceived(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Timereceived\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlocktimeAndUID(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlocktimeAndUID(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlocktimeAndSalt(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlocktimeAndSalt(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlocktimeAndCreated(value1 int, value2 time.Time) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlocktimeAndCreated(value1 int, value2 time.Time) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlocktimeAndTxid(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Txid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlocktimeAndTxid(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Txid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlocktimeAndValid(value1 int, value2 bool) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Valid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlocktimeAndValid(value1 int, value2 bool) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blocktime\" = $1 AND \"Valid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) CountStreamItemByBlockindex(value int) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"StreamItem\" WHERE \"Blockindex\" = $1;", value)
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

func (db *DB) GetStreamItemByBlockindex(value int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockindex(value int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockindexAndBlockhash(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Blockhash\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockindexAndBlockhash(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Blockhash\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockindexAndData(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Data\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockindexAndData(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Data\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockindexAndKeys(value1 int, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Keys\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockindexAndKeys(value1 int, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Keys\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockindexAndVout(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Vout\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockindexAndVout(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Vout\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockindexAndBlocktime(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Blocktime\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockindexAndBlocktime(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Blocktime\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockindexAndConfirmations(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Confirmations\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockindexAndConfirmations(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Confirmations\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockindexAndPublishers(value1 int, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Publishers\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockindexAndPublishers(value1 int, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Publishers\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockindexAndTime(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockindexAndTime(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockindexAndTimereceived(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Timereceived\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockindexAndTimereceived(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Timereceived\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockindexAndUID(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockindexAndUID(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockindexAndSalt(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockindexAndSalt(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockindexAndCreated(value1 int, value2 time.Time) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockindexAndCreated(value1 int, value2 time.Time) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockindexAndTxid(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Txid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockindexAndTxid(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Txid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByBlockindexAndValid(value1 int, value2 bool) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Valid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByBlockindexAndValid(value1 int, value2 bool) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Blockindex\" = $1 AND \"Valid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) CountStreamItemByConfirmations(value int) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"StreamItem\" WHERE \"Confirmations\" = $1;", value)
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

func (db *DB) GetStreamItemByConfirmations(value int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByConfirmations(value int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByConfirmationsAndBlockhash(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Blockhash\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByConfirmationsAndBlockhash(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Blockhash\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByConfirmationsAndData(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Data\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByConfirmationsAndData(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Data\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByConfirmationsAndKeys(value1 int, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Keys\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByConfirmationsAndKeys(value1 int, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Keys\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByConfirmationsAndVout(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Vout\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByConfirmationsAndVout(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Vout\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByConfirmationsAndBlocktime(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Blocktime\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByConfirmationsAndBlocktime(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Blocktime\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByConfirmationsAndBlockindex(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Blockindex\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByConfirmationsAndBlockindex(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Blockindex\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByConfirmationsAndPublishers(value1 int, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Publishers\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByConfirmationsAndPublishers(value1 int, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Publishers\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByConfirmationsAndTime(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByConfirmationsAndTime(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByConfirmationsAndTimereceived(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Timereceived\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByConfirmationsAndTimereceived(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Timereceived\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByConfirmationsAndUID(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByConfirmationsAndUID(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByConfirmationsAndSalt(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByConfirmationsAndSalt(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByConfirmationsAndCreated(value1 int, value2 time.Time) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByConfirmationsAndCreated(value1 int, value2 time.Time) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByConfirmationsAndTxid(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Txid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByConfirmationsAndTxid(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Txid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByConfirmationsAndValid(value1 int, value2 bool) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Valid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByConfirmationsAndValid(value1 int, value2 bool) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Confirmations\" = $1 AND \"Valid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) CountStreamItemByPublishers(value []string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"StreamItem\" WHERE \"Publishers\" = $1;", value)
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

func (db *DB) GetStreamItemByPublishers(value []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByPublishers(value []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByPublishersAndBlockhash(value1 []string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Blockhash\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByPublishersAndBlockhash(value1 []string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Blockhash\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByPublishersAndData(value1 []string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Data\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByPublishersAndData(value1 []string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Data\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByPublishersAndKeys(value1 []string, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Keys\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByPublishersAndKeys(value1 []string, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Keys\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByPublishersAndVout(value1 []string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Vout\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByPublishersAndVout(value1 []string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Vout\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByPublishersAndBlocktime(value1 []string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Blocktime\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByPublishersAndBlocktime(value1 []string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Blocktime\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByPublishersAndBlockindex(value1 []string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Blockindex\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByPublishersAndBlockindex(value1 []string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Blockindex\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByPublishersAndConfirmations(value1 []string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Confirmations\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByPublishersAndConfirmations(value1 []string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Confirmations\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByPublishersAndTime(value1 []string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByPublishersAndTime(value1 []string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByPublishersAndTimereceived(value1 []string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Timereceived\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByPublishersAndTimereceived(value1 []string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Timereceived\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByPublishersAndUID(value1 []string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByPublishersAndUID(value1 []string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByPublishersAndSalt(value1 []string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByPublishersAndSalt(value1 []string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByPublishersAndCreated(value1 []string, value2 time.Time) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByPublishersAndCreated(value1 []string, value2 time.Time) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByPublishersAndTxid(value1 []string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Txid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByPublishersAndTxid(value1 []string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Txid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByPublishersAndValid(value1 []string, value2 bool) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Valid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByPublishersAndValid(value1 []string, value2 bool) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Publishers\" = $1 AND \"Valid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) CountStreamItemByTime(value int) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"StreamItem\" WHERE \"Time\" = $1;", value)
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

func (db *DB) GetStreamItemByTime(value int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTime(value int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimeAndBlockhash(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Blockhash\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimeAndBlockhash(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Blockhash\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimeAndData(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Data\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimeAndData(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Data\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimeAndKeys(value1 int, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Keys\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimeAndKeys(value1 int, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Keys\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimeAndVout(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Vout\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimeAndVout(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Vout\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimeAndBlocktime(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Blocktime\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimeAndBlocktime(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Blocktime\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimeAndBlockindex(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Blockindex\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimeAndBlockindex(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Blockindex\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimeAndConfirmations(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Confirmations\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimeAndConfirmations(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Confirmations\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimeAndPublishers(value1 int, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Publishers\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimeAndPublishers(value1 int, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Publishers\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimeAndTimereceived(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Timereceived\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimeAndTimereceived(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Timereceived\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimeAndUID(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimeAndUID(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimeAndSalt(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimeAndSalt(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimeAndCreated(value1 int, value2 time.Time) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimeAndCreated(value1 int, value2 time.Time) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimeAndTxid(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Txid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimeAndTxid(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Txid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimeAndValid(value1 int, value2 bool) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Valid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimeAndValid(value1 int, value2 bool) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Time\" = $1 AND \"Valid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) CountStreamItemByTimereceived(value int) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"StreamItem\" WHERE \"Timereceived\" = $1;", value)
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

func (db *DB) GetStreamItemByTimereceived(value int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimereceived(value int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimereceivedAndBlockhash(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Blockhash\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimereceivedAndBlockhash(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Blockhash\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimereceivedAndData(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Data\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimereceivedAndData(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Data\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimereceivedAndKeys(value1 int, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Keys\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimereceivedAndKeys(value1 int, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Keys\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimereceivedAndVout(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Vout\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimereceivedAndVout(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Vout\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimereceivedAndBlocktime(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Blocktime\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimereceivedAndBlocktime(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Blocktime\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimereceivedAndBlockindex(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Blockindex\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimereceivedAndBlockindex(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Blockindex\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimereceivedAndConfirmations(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Confirmations\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimereceivedAndConfirmations(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Confirmations\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimereceivedAndPublishers(value1 int, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Publishers\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimereceivedAndPublishers(value1 int, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Publishers\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimereceivedAndTime(value1 int, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimereceivedAndTime(value1 int, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimereceivedAndUID(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimereceivedAndUID(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimereceivedAndSalt(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimereceivedAndSalt(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimereceivedAndCreated(value1 int, value2 time.Time) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimereceivedAndCreated(value1 int, value2 time.Time) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimereceivedAndTxid(value1 int, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Txid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimereceivedAndTxid(value1 int, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Txid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTimereceivedAndValid(value1 int, value2 bool) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Valid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTimereceivedAndValid(value1 int, value2 bool) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Timereceived\" = $1 AND \"Valid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) CountStreamItemByUID(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"StreamItem\" WHERE \"UID\" = $1;", value)
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

func (db *DB) GetStreamItemByUID(value string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"UID\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByUID(value string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"UID\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) CountStreamItemBySalt(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"StreamItem\" WHERE \"Salt\" = $1;", value)
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

func (db *DB) GetStreamItemBySalt(value string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemBySalt(value string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemBySaltAndBlockhash(value1 string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Blockhash\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemBySaltAndBlockhash(value1 string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Blockhash\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemBySaltAndData(value1 string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Data\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemBySaltAndData(value1 string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Data\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemBySaltAndKeys(value1 string, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Keys\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemBySaltAndKeys(value1 string, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Keys\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemBySaltAndVout(value1 string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Vout\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemBySaltAndVout(value1 string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Vout\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemBySaltAndBlocktime(value1 string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Blocktime\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemBySaltAndBlocktime(value1 string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Blocktime\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemBySaltAndBlockindex(value1 string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Blockindex\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemBySaltAndBlockindex(value1 string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Blockindex\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemBySaltAndConfirmations(value1 string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Confirmations\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemBySaltAndConfirmations(value1 string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Confirmations\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemBySaltAndPublishers(value1 string, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Publishers\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemBySaltAndPublishers(value1 string, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Publishers\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemBySaltAndTime(value1 string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemBySaltAndTime(value1 string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemBySaltAndTimereceived(value1 string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Timereceived\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemBySaltAndTimereceived(value1 string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Timereceived\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemBySaltAndUID(value1 string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemBySaltAndUID(value1 string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemBySaltAndCreated(value1 string, value2 time.Time) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemBySaltAndCreated(value1 string, value2 time.Time) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemBySaltAndTxid(value1 string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Txid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemBySaltAndTxid(value1 string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Txid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemBySaltAndValid(value1 string, value2 bool) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Valid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemBySaltAndValid(value1 string, value2 bool) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Salt\" = $1 AND \"Valid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) CountStreamItemByCreated(value time.Time) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"StreamItem\" WHERE \"Created\" = $1;", value)
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

func (db *DB) GetStreamItemByCreated(value time.Time) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByCreated(value time.Time) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByCreatedAndBlockhash(value1 time.Time, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Blockhash\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByCreatedAndBlockhash(value1 time.Time, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Blockhash\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByCreatedAndData(value1 time.Time, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Data\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByCreatedAndData(value1 time.Time, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Data\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByCreatedAndKeys(value1 time.Time, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Keys\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByCreatedAndKeys(value1 time.Time, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Keys\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByCreatedAndVout(value1 time.Time, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Vout\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByCreatedAndVout(value1 time.Time, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Vout\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByCreatedAndBlocktime(value1 time.Time, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Blocktime\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByCreatedAndBlocktime(value1 time.Time, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Blocktime\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByCreatedAndBlockindex(value1 time.Time, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Blockindex\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByCreatedAndBlockindex(value1 time.Time, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Blockindex\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByCreatedAndConfirmations(value1 time.Time, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Confirmations\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByCreatedAndConfirmations(value1 time.Time, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Confirmations\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByCreatedAndPublishers(value1 time.Time, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Publishers\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByCreatedAndPublishers(value1 time.Time, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Publishers\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByCreatedAndTime(value1 time.Time, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByCreatedAndTime(value1 time.Time, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByCreatedAndTimereceived(value1 time.Time, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Timereceived\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByCreatedAndTimereceived(value1 time.Time, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Timereceived\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByCreatedAndUID(value1 time.Time, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByCreatedAndUID(value1 time.Time, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByCreatedAndSalt(value1 time.Time, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByCreatedAndSalt(value1 time.Time, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByCreatedAndTxid(value1 time.Time, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Txid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByCreatedAndTxid(value1 time.Time, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Txid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByCreatedAndValid(value1 time.Time, value2 bool) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Valid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByCreatedAndValid(value1 time.Time, value2 bool) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Created\" = $1 AND \"Valid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) CountStreamItemByTxid(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"StreamItem\" WHERE \"Txid\" = $1;", value)
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

func (db *DB) GetStreamItemByTxid(value string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTxid(value string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTxidAndBlockhash(value1 string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Blockhash\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTxidAndBlockhash(value1 string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Blockhash\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTxidAndData(value1 string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Data\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTxidAndData(value1 string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Data\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTxidAndKeys(value1 string, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Keys\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTxidAndKeys(value1 string, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Keys\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTxidAndVout(value1 string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Vout\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTxidAndVout(value1 string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Vout\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTxidAndBlocktime(value1 string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Blocktime\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTxidAndBlocktime(value1 string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Blocktime\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTxidAndBlockindex(value1 string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Blockindex\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTxidAndBlockindex(value1 string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Blockindex\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTxidAndConfirmations(value1 string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Confirmations\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTxidAndConfirmations(value1 string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Confirmations\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTxidAndPublishers(value1 string, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Publishers\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTxidAndPublishers(value1 string, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Publishers\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTxidAndTime(value1 string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTxidAndTime(value1 string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTxidAndTimereceived(value1 string, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Timereceived\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTxidAndTimereceived(value1 string, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Timereceived\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTxidAndUID(value1 string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTxidAndUID(value1 string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTxidAndSalt(value1 string, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTxidAndSalt(value1 string, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTxidAndCreated(value1 string, value2 time.Time) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTxidAndCreated(value1 string, value2 time.Time) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByTxidAndValid(value1 string, value2 bool) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Valid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByTxidAndValid(value1 string, value2 bool) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Txid\" = $1 AND \"Valid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) CountStreamItemByValid(value bool) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"StreamItem\" WHERE \"Valid\" = $1;", value)
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

func (db *DB) GetStreamItemByValid(value bool) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByValid(value bool) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByValidAndBlockhash(value1 bool, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Blockhash\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByValidAndBlockhash(value1 bool, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Blockhash\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByValidAndData(value1 bool, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Data\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByValidAndData(value1 bool, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Data\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByValidAndKeys(value1 bool, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Keys\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByValidAndKeys(value1 bool, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Keys\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByValidAndVout(value1 bool, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Vout\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByValidAndVout(value1 bool, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Vout\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByValidAndBlocktime(value1 bool, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Blocktime\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByValidAndBlocktime(value1 bool, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Blocktime\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByValidAndBlockindex(value1 bool, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Blockindex\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByValidAndBlockindex(value1 bool, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Blockindex\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByValidAndConfirmations(value1 bool, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Confirmations\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByValidAndConfirmations(value1 bool, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Confirmations\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByValidAndPublishers(value1 bool, value2 []string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Publishers\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByValidAndPublishers(value1 bool, value2 []string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Publishers\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByValidAndTime(value1 bool, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByValidAndTime(value1 bool, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByValidAndTimereceived(value1 bool, value2 int) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Timereceived\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByValidAndTimereceived(value1 bool, value2 int) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Timereceived\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByValidAndUID(value1 bool, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByValidAndUID(value1 bool, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByValidAndSalt(value1 bool, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByValidAndSalt(value1 bool, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByValidAndCreated(value1 bool, value2 time.Time) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByValidAndCreated(value1 bool, value2 time.Time) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}

func (db *DB) GetStreamItemByValidAndTxid(value1 bool, value2 string) (bool, *models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Txid\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStreamItem(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamItemByValidAndTxid(value1 bool, value2 string) ([]*models.StreamItem, error) {
	rows, err := db.DoQuery("SELECT \"Blockhash\", \"Blockindex\", \"Blocktime\", \"Confirmations\", \"Created\", \"Data\", \"Keys\", \"Publishers\", \"Salt\", \"Time\", \"Timereceived\", \"Txid\", \"UID\", \"Valid\", \"Vout\" FROM " + db.dbName + ".\"StreamItem\" WHERE \"Valid\" = $1 AND \"Txid\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStreamItem(rows)
}


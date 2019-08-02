package database

import (
	"fmt"
	"time"
	"database/sql"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/project"
)

func (db *DB) CreateTableBillingRecord() error {
	q := "CREATE TABLE IF NOT EXISTS " + db.dbName + ".\"BillingRecord\" (\"Salt\" UUID DEFAULT gen_random_uuid(), \"SendAddressUID\" STRING NOT NULL, \"OutputsTotal\" FLOAT NOT NULL, \"Time\" int64 NOT NULL, \"UID\" UUID PRIMARY KEY DEFAULT gen_random_uuid(), \"Created\" TIMESTAMP DEFAULT current_timestamp(), \"User\" STRING NOT NULL, \"SendAddress\" STRING NOT NULL, \"Currency\" STRING NOT NULL, \"Address\" STRING NOT NULL, \"TxId\" STRING NOT NULL, \"Sender\" STRING NOT NULL, \"Project\" STRING NOT NULL, \"Wallet\" STRING NOT NULL, \"Receiver\" STRING NOT NULL, \"ReceiveAddress\" STRING NOT NULL, \"ReceiveAddressUID\" STRING NOT NULL, \"Size\" INT NOT NULL);"
	fmt.Println(q)
	_, err := db.DoExec(q)
	return err
}

func (db *DB) InsertBillingRecord(Address string, Currency string, OutputsTotal float64, Project string, ReceiveAddress string, ReceiveAddressUID string, Receiver string, SendAddress string, SendAddressUID string, Sender string, Size int, Time int64, TxId string, User string, Wallet string) (*models.BillingRecord, error) {
	row := &models.BillingRecord{}
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"BillingRecord\" (\"Address\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"User\", \"Wallet\") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING Address, Currency, OutputsTotal, Project, ReceiveAddress, ReceiveAddressUID, Receiver, SendAddress, SendAddressUID, Sender, Size, Time, TxId, User, Wallet;", Address, Currency, OutputsTotal, Project, ReceiveAddress, ReceiveAddressUID, Receiver, SendAddress, SendAddressUID, Sender, Size, Time, TxId, User, Wallet).Scan(&row.Address, &row.Created, &row.Currency, &row.OutputsTotal, &row.Project, &row.ReceiveAddress, &row.ReceiveAddressUID, &row.Receiver, &row.Salt, &row.SendAddress, &row.SendAddressUID, &row.Sender, &row.Size, &row.Time, &row.TxId, &row.UID, &row.User, &row.Wallet); err != nil {
		return nil, err
	}
	return row, nil
}

func (db *DB) DeleteBillingRecord(uid string) error {
	_, err := db.DoQuery("DELETE FROM " + db.dbName + ".\"BillingRecord\" WHERE \"UID\" = $1;", uid)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) PutBillingRecord(model *models.BillingRecord) error {
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"BillingRecord\" (\"Address\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"User\", \"Wallet\") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\";", model.Address, model.Currency, model.OutputsTotal, model.Project, model.ReceiveAddress, model.ReceiveAddressUID, model.Receiver, model.SendAddress, model.SendAddressUID, model.Sender, model.Size, model.Time, model.TxId, model.User, model.Wallet).Scan(&model.Address, &model.Created, &model.Currency, &model.OutputsTotal, &model.Project, &model.ReceiveAddress, &model.ReceiveAddressUID, &model.Receiver, &model.Salt, &model.SendAddress, &model.SendAddressUID, &model.Sender, &model.Size, &model.Time, &model.TxId, &model.UID, &model.User, &model.Wallet); err != nil {
		return err
	}
	return nil
}

func (db *DB) ScanToBillingRecord(rows *sql.Rows) ([]*models.BillingRecord, error) {
	defer rows.Close()
	results := []*models.BillingRecord{}
	for rows.Next() {
		result := &models.BillingRecord{}
		err := rows.Scan(&result.Address, &result.Created, &result.Currency, &result.OutputsTotal, &result.Project, &result.ReceiveAddress, &result.ReceiveAddressUID, &result.Receiver, &result.Salt, &result.SendAddress, &result.SendAddressUID, &result.Sender, &result.Size, &result.Time, &result.TxId, &result.UID, &result.User, &result.Wallet)
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

func (db *DB) CountBillingRecord() (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM " + db.dbName + ".\"BillingRecord\";")
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

func (db *DB) QueryBillingRecord() ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\";")
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) CountBillingRecordBySalt(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"BillingRecord\" WHERE \"Salt\" = $1;", value)
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

func (db *DB) GetBillingRecordBySalt(value string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySalt(value string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySaltAndSendAddressUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySaltAndSendAddressUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySaltAndOutputsTotal(value1 string, value2 float64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySaltAndOutputsTotal(value1 string, value2 float64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySaltAndTime(value1 string, value2 int64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySaltAndTime(value1 string, value2 int64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySaltAndUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySaltAndUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySaltAndCreated(value1 string, value2 time.Time) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySaltAndCreated(value1 string, value2 time.Time) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySaltAndUser(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySaltAndUser(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySaltAndSendAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySaltAndSendAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySaltAndCurrency(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySaltAndCurrency(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySaltAndAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySaltAndAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySaltAndTxId(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySaltAndTxId(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySaltAndSender(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySaltAndSender(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySaltAndProject(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySaltAndProject(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySaltAndWallet(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySaltAndWallet(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySaltAndReceiver(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySaltAndReceiver(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySaltAndReceiveAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySaltAndReceiveAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySaltAndReceiveAddressUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySaltAndReceiveAddressUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySaltAndSize(value1 string, value2 int) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySaltAndSize(value1 string, value2 int) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Salt\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) CountBillingRecordBySendAddressUID(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"BillingRecord\" WHERE \"SendAddressUID\" = $1;", value)
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

func (db *DB) GetBillingRecordBySendAddressUID(value string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressUID(value string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressUIDAndSalt(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressUIDAndSalt(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressUIDAndOutputsTotal(value1 string, value2 float64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressUIDAndOutputsTotal(value1 string, value2 float64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressUIDAndTime(value1 string, value2 int64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressUIDAndTime(value1 string, value2 int64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressUIDAndUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressUIDAndUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressUIDAndCreated(value1 string, value2 time.Time) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressUIDAndCreated(value1 string, value2 time.Time) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressUIDAndUser(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressUIDAndUser(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressUIDAndSendAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressUIDAndSendAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressUIDAndCurrency(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressUIDAndCurrency(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressUIDAndAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressUIDAndAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressUIDAndTxId(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressUIDAndTxId(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressUIDAndSender(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressUIDAndSender(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressUIDAndProject(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressUIDAndProject(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressUIDAndWallet(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressUIDAndWallet(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressUIDAndReceiver(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressUIDAndReceiver(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressUIDAndReceiveAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressUIDAndReceiveAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressUIDAndReceiveAddressUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressUIDAndReceiveAddressUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressUIDAndSize(value1 string, value2 int) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressUIDAndSize(value1 string, value2 int) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddressUID\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) CountBillingRecordByOutputsTotal(value float64) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"BillingRecord\" WHERE \"OutputsTotal\" = $1;", value)
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

func (db *DB) GetBillingRecordByOutputsTotal(value float64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByOutputsTotal(value float64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByOutputsTotalAndSalt(value1 float64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByOutputsTotalAndSalt(value1 float64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByOutputsTotalAndSendAddressUID(value1 float64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByOutputsTotalAndSendAddressUID(value1 float64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByOutputsTotalAndTime(value1 float64, value2 int64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByOutputsTotalAndTime(value1 float64, value2 int64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByOutputsTotalAndUID(value1 float64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByOutputsTotalAndUID(value1 float64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByOutputsTotalAndCreated(value1 float64, value2 time.Time) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByOutputsTotalAndCreated(value1 float64, value2 time.Time) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByOutputsTotalAndUser(value1 float64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByOutputsTotalAndUser(value1 float64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByOutputsTotalAndSendAddress(value1 float64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByOutputsTotalAndSendAddress(value1 float64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByOutputsTotalAndCurrency(value1 float64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByOutputsTotalAndCurrency(value1 float64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByOutputsTotalAndAddress(value1 float64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByOutputsTotalAndAddress(value1 float64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByOutputsTotalAndTxId(value1 float64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByOutputsTotalAndTxId(value1 float64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByOutputsTotalAndSender(value1 float64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByOutputsTotalAndSender(value1 float64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByOutputsTotalAndProject(value1 float64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByOutputsTotalAndProject(value1 float64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByOutputsTotalAndWallet(value1 float64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByOutputsTotalAndWallet(value1 float64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByOutputsTotalAndReceiver(value1 float64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByOutputsTotalAndReceiver(value1 float64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByOutputsTotalAndReceiveAddress(value1 float64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByOutputsTotalAndReceiveAddress(value1 float64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByOutputsTotalAndReceiveAddressUID(value1 float64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByOutputsTotalAndReceiveAddressUID(value1 float64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByOutputsTotalAndSize(value1 float64, value2 int) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByOutputsTotalAndSize(value1 float64, value2 int) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"OutputsTotal\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) CountBillingRecordByTime(value int64) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"BillingRecord\" WHERE \"Time\" = $1;", value)
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

func (db *DB) GetBillingRecordByTime(value int64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTime(value int64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTimeAndSalt(value1 int64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTimeAndSalt(value1 int64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTimeAndSendAddressUID(value1 int64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTimeAndSendAddressUID(value1 int64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTimeAndOutputsTotal(value1 int64, value2 float64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTimeAndOutputsTotal(value1 int64, value2 float64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTimeAndUID(value1 int64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTimeAndUID(value1 int64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTimeAndCreated(value1 int64, value2 time.Time) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTimeAndCreated(value1 int64, value2 time.Time) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTimeAndUser(value1 int64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTimeAndUser(value1 int64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTimeAndSendAddress(value1 int64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTimeAndSendAddress(value1 int64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTimeAndCurrency(value1 int64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTimeAndCurrency(value1 int64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTimeAndAddress(value1 int64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTimeAndAddress(value1 int64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTimeAndTxId(value1 int64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTimeAndTxId(value1 int64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTimeAndSender(value1 int64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTimeAndSender(value1 int64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTimeAndProject(value1 int64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTimeAndProject(value1 int64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTimeAndWallet(value1 int64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTimeAndWallet(value1 int64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTimeAndReceiver(value1 int64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTimeAndReceiver(value1 int64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTimeAndReceiveAddress(value1 int64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTimeAndReceiveAddress(value1 int64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTimeAndReceiveAddressUID(value1 int64, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTimeAndReceiveAddressUID(value1 int64, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTimeAndSize(value1 int64, value2 int) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTimeAndSize(value1 int64, value2 int) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Time\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) CountBillingRecordByUID(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"BillingRecord\" WHERE \"UID\" = $1;", value)
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

func (db *DB) GetBillingRecordByUID(value string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"UID\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByUID(value string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"UID\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) CountBillingRecordByCreated(value time.Time) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"BillingRecord\" WHERE \"Created\" = $1;", value)
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

func (db *DB) GetBillingRecordByCreated(value time.Time) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCreated(value time.Time) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCreatedAndSalt(value1 time.Time, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCreatedAndSalt(value1 time.Time, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCreatedAndSendAddressUID(value1 time.Time, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCreatedAndSendAddressUID(value1 time.Time, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCreatedAndOutputsTotal(value1 time.Time, value2 float64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCreatedAndOutputsTotal(value1 time.Time, value2 float64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCreatedAndTime(value1 time.Time, value2 int64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCreatedAndTime(value1 time.Time, value2 int64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCreatedAndUID(value1 time.Time, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCreatedAndUID(value1 time.Time, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCreatedAndUser(value1 time.Time, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCreatedAndUser(value1 time.Time, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCreatedAndSendAddress(value1 time.Time, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCreatedAndSendAddress(value1 time.Time, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCreatedAndCurrency(value1 time.Time, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCreatedAndCurrency(value1 time.Time, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCreatedAndAddress(value1 time.Time, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCreatedAndAddress(value1 time.Time, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCreatedAndTxId(value1 time.Time, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCreatedAndTxId(value1 time.Time, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCreatedAndSender(value1 time.Time, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCreatedAndSender(value1 time.Time, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCreatedAndProject(value1 time.Time, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCreatedAndProject(value1 time.Time, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCreatedAndWallet(value1 time.Time, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCreatedAndWallet(value1 time.Time, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCreatedAndReceiver(value1 time.Time, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCreatedAndReceiver(value1 time.Time, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCreatedAndReceiveAddress(value1 time.Time, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCreatedAndReceiveAddress(value1 time.Time, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCreatedAndReceiveAddressUID(value1 time.Time, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCreatedAndReceiveAddressUID(value1 time.Time, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCreatedAndSize(value1 time.Time, value2 int) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCreatedAndSize(value1 time.Time, value2 int) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Created\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) CountBillingRecordByUser(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"BillingRecord\" WHERE \"User\" = $1;", value)
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

func (db *DB) GetBillingRecordByUser(value string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByUser(value string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByUserAndSalt(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByUserAndSalt(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByUserAndSendAddressUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByUserAndSendAddressUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByUserAndOutputsTotal(value1 string, value2 float64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByUserAndOutputsTotal(value1 string, value2 float64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByUserAndTime(value1 string, value2 int64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByUserAndTime(value1 string, value2 int64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByUserAndUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByUserAndUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByUserAndCreated(value1 string, value2 time.Time) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByUserAndCreated(value1 string, value2 time.Time) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByUserAndSendAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByUserAndSendAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByUserAndCurrency(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByUserAndCurrency(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByUserAndAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByUserAndAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByUserAndTxId(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByUserAndTxId(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByUserAndSender(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByUserAndSender(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByUserAndProject(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByUserAndProject(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByUserAndWallet(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByUserAndWallet(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByUserAndReceiver(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByUserAndReceiver(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByUserAndReceiveAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByUserAndReceiveAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByUserAndReceiveAddressUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByUserAndReceiveAddressUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByUserAndSize(value1 string, value2 int) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByUserAndSize(value1 string, value2 int) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"User\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) CountBillingRecordBySendAddress(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"BillingRecord\" WHERE \"SendAddress\" = $1;", value)
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

func (db *DB) GetBillingRecordBySendAddress(value string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddress(value string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressAndSalt(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressAndSalt(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressAndSendAddressUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressAndSendAddressUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressAndOutputsTotal(value1 string, value2 float64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressAndOutputsTotal(value1 string, value2 float64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressAndTime(value1 string, value2 int64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressAndTime(value1 string, value2 int64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressAndUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressAndUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressAndCreated(value1 string, value2 time.Time) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressAndCreated(value1 string, value2 time.Time) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressAndUser(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressAndUser(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressAndCurrency(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressAndCurrency(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressAndAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressAndAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressAndTxId(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressAndTxId(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressAndSender(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressAndSender(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressAndProject(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressAndProject(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressAndWallet(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressAndWallet(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressAndReceiver(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressAndReceiver(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressAndReceiveAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressAndReceiveAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressAndReceiveAddressUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressAndReceiveAddressUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySendAddressAndSize(value1 string, value2 int) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySendAddressAndSize(value1 string, value2 int) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"SendAddress\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) CountBillingRecordByCurrency(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"BillingRecord\" WHERE \"Currency\" = $1;", value)
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

func (db *DB) GetBillingRecordByCurrency(value string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCurrency(value string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCurrencyAndSalt(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCurrencyAndSalt(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCurrencyAndSendAddressUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCurrencyAndSendAddressUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCurrencyAndOutputsTotal(value1 string, value2 float64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCurrencyAndOutputsTotal(value1 string, value2 float64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCurrencyAndTime(value1 string, value2 int64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCurrencyAndTime(value1 string, value2 int64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCurrencyAndUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCurrencyAndUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCurrencyAndCreated(value1 string, value2 time.Time) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCurrencyAndCreated(value1 string, value2 time.Time) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCurrencyAndUser(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCurrencyAndUser(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCurrencyAndSendAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCurrencyAndSendAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCurrencyAndAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCurrencyAndAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCurrencyAndTxId(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCurrencyAndTxId(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCurrencyAndSender(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCurrencyAndSender(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCurrencyAndProject(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCurrencyAndProject(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCurrencyAndWallet(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCurrencyAndWallet(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCurrencyAndReceiver(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCurrencyAndReceiver(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCurrencyAndReceiveAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCurrencyAndReceiveAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCurrencyAndReceiveAddressUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCurrencyAndReceiveAddressUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByCurrencyAndSize(value1 string, value2 int) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByCurrencyAndSize(value1 string, value2 int) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Currency\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) CountBillingRecordByAddress(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"BillingRecord\" WHERE \"Address\" = $1;", value)
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

func (db *DB) GetBillingRecordByAddress(value string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByAddress(value string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByAddressAndSalt(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByAddressAndSalt(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByAddressAndSendAddressUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByAddressAndSendAddressUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByAddressAndOutputsTotal(value1 string, value2 float64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByAddressAndOutputsTotal(value1 string, value2 float64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByAddressAndTime(value1 string, value2 int64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByAddressAndTime(value1 string, value2 int64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByAddressAndUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByAddressAndUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByAddressAndCreated(value1 string, value2 time.Time) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByAddressAndCreated(value1 string, value2 time.Time) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByAddressAndUser(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByAddressAndUser(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByAddressAndSendAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByAddressAndSendAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByAddressAndCurrency(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByAddressAndCurrency(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByAddressAndTxId(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByAddressAndTxId(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByAddressAndSender(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByAddressAndSender(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByAddressAndProject(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByAddressAndProject(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByAddressAndWallet(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByAddressAndWallet(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByAddressAndReceiver(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByAddressAndReceiver(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByAddressAndReceiveAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByAddressAndReceiveAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByAddressAndReceiveAddressUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByAddressAndReceiveAddressUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByAddressAndSize(value1 string, value2 int) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByAddressAndSize(value1 string, value2 int) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Address\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) CountBillingRecordByTxId(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"BillingRecord\" WHERE \"TxId\" = $1;", value)
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

func (db *DB) GetBillingRecordByTxId(value string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTxId(value string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTxIdAndSalt(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTxIdAndSalt(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTxIdAndSendAddressUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTxIdAndSendAddressUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTxIdAndOutputsTotal(value1 string, value2 float64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTxIdAndOutputsTotal(value1 string, value2 float64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTxIdAndTime(value1 string, value2 int64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTxIdAndTime(value1 string, value2 int64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTxIdAndUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTxIdAndUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTxIdAndCreated(value1 string, value2 time.Time) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTxIdAndCreated(value1 string, value2 time.Time) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTxIdAndUser(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTxIdAndUser(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTxIdAndSendAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTxIdAndSendAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTxIdAndCurrency(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTxIdAndCurrency(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTxIdAndAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTxIdAndAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTxIdAndSender(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTxIdAndSender(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTxIdAndProject(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTxIdAndProject(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTxIdAndWallet(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTxIdAndWallet(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTxIdAndReceiver(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTxIdAndReceiver(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTxIdAndReceiveAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTxIdAndReceiveAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTxIdAndReceiveAddressUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTxIdAndReceiveAddressUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByTxIdAndSize(value1 string, value2 int) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByTxIdAndSize(value1 string, value2 int) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"TxId\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) CountBillingRecordBySender(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"BillingRecord\" WHERE \"Sender\" = $1;", value)
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

func (db *DB) GetBillingRecordBySender(value string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySender(value string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySenderAndSalt(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySenderAndSalt(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySenderAndSendAddressUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySenderAndSendAddressUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySenderAndOutputsTotal(value1 string, value2 float64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySenderAndOutputsTotal(value1 string, value2 float64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySenderAndTime(value1 string, value2 int64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySenderAndTime(value1 string, value2 int64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySenderAndUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySenderAndUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySenderAndCreated(value1 string, value2 time.Time) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySenderAndCreated(value1 string, value2 time.Time) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySenderAndUser(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySenderAndUser(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySenderAndSendAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySenderAndSendAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySenderAndCurrency(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySenderAndCurrency(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySenderAndAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySenderAndAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySenderAndTxId(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySenderAndTxId(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySenderAndProject(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySenderAndProject(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySenderAndWallet(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySenderAndWallet(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySenderAndReceiver(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySenderAndReceiver(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySenderAndReceiveAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySenderAndReceiveAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySenderAndReceiveAddressUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySenderAndReceiveAddressUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySenderAndSize(value1 string, value2 int) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySenderAndSize(value1 string, value2 int) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Sender\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) CountBillingRecordByProject(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"BillingRecord\" WHERE \"Project\" = $1;", value)
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

func (db *DB) GetBillingRecordByProject(value string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByProject(value string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByProjectAndSalt(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByProjectAndSalt(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByProjectAndSendAddressUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByProjectAndSendAddressUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByProjectAndOutputsTotal(value1 string, value2 float64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByProjectAndOutputsTotal(value1 string, value2 float64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByProjectAndTime(value1 string, value2 int64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByProjectAndTime(value1 string, value2 int64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByProjectAndUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByProjectAndUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByProjectAndCreated(value1 string, value2 time.Time) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByProjectAndCreated(value1 string, value2 time.Time) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByProjectAndUser(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByProjectAndUser(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByProjectAndSendAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByProjectAndSendAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByProjectAndCurrency(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByProjectAndCurrency(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByProjectAndAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByProjectAndAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByProjectAndTxId(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByProjectAndTxId(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByProjectAndSender(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByProjectAndSender(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByProjectAndWallet(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByProjectAndWallet(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByProjectAndReceiver(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByProjectAndReceiver(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByProjectAndReceiveAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByProjectAndReceiveAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByProjectAndReceiveAddressUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByProjectAndReceiveAddressUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByProjectAndSize(value1 string, value2 int) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByProjectAndSize(value1 string, value2 int) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Project\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) CountBillingRecordByWallet(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"BillingRecord\" WHERE \"Wallet\" = $1;", value)
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

func (db *DB) GetBillingRecordByWallet(value string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByWallet(value string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByWalletAndSalt(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByWalletAndSalt(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByWalletAndSendAddressUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByWalletAndSendAddressUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByWalletAndOutputsTotal(value1 string, value2 float64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByWalletAndOutputsTotal(value1 string, value2 float64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByWalletAndTime(value1 string, value2 int64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByWalletAndTime(value1 string, value2 int64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByWalletAndUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByWalletAndUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByWalletAndCreated(value1 string, value2 time.Time) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByWalletAndCreated(value1 string, value2 time.Time) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByWalletAndUser(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByWalletAndUser(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByWalletAndSendAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByWalletAndSendAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByWalletAndCurrency(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByWalletAndCurrency(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByWalletAndAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByWalletAndAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByWalletAndTxId(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByWalletAndTxId(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByWalletAndSender(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByWalletAndSender(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByWalletAndProject(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByWalletAndProject(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByWalletAndReceiver(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByWalletAndReceiver(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByWalletAndReceiveAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByWalletAndReceiveAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByWalletAndReceiveAddressUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByWalletAndReceiveAddressUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByWalletAndSize(value1 string, value2 int) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByWalletAndSize(value1 string, value2 int) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Wallet\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) CountBillingRecordByReceiver(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"BillingRecord\" WHERE \"Receiver\" = $1;", value)
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

func (db *DB) GetBillingRecordByReceiver(value string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiver(value string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiverAndSalt(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiverAndSalt(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiverAndSendAddressUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiverAndSendAddressUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiverAndOutputsTotal(value1 string, value2 float64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiverAndOutputsTotal(value1 string, value2 float64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiverAndTime(value1 string, value2 int64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiverAndTime(value1 string, value2 int64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiverAndUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiverAndUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiverAndCreated(value1 string, value2 time.Time) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiverAndCreated(value1 string, value2 time.Time) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiverAndUser(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiverAndUser(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiverAndSendAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiverAndSendAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiverAndCurrency(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiverAndCurrency(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiverAndAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiverAndAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiverAndTxId(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiverAndTxId(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiverAndSender(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiverAndSender(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiverAndProject(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiverAndProject(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiverAndWallet(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiverAndWallet(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiverAndReceiveAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiverAndReceiveAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiverAndReceiveAddressUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiverAndReceiveAddressUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiverAndSize(value1 string, value2 int) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiverAndSize(value1 string, value2 int) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Receiver\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) CountBillingRecordByReceiveAddress(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"BillingRecord\" WHERE \"ReceiveAddress\" = $1;", value)
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

func (db *DB) GetBillingRecordByReceiveAddress(value string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddress(value string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressAndSalt(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressAndSalt(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressAndSendAddressUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressAndSendAddressUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressAndOutputsTotal(value1 string, value2 float64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressAndOutputsTotal(value1 string, value2 float64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressAndTime(value1 string, value2 int64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressAndTime(value1 string, value2 int64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressAndUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressAndUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressAndCreated(value1 string, value2 time.Time) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressAndCreated(value1 string, value2 time.Time) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressAndUser(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressAndUser(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressAndSendAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressAndSendAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressAndCurrency(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressAndCurrency(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressAndAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressAndAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressAndTxId(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressAndTxId(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressAndSender(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressAndSender(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressAndProject(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressAndProject(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressAndWallet(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressAndWallet(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressAndReceiver(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressAndReceiver(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressAndReceiveAddressUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressAndReceiveAddressUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressAndSize(value1 string, value2 int) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressAndSize(value1 string, value2 int) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddress\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) CountBillingRecordByReceiveAddressUID(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1;", value)
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

func (db *DB) GetBillingRecordByReceiveAddressUID(value string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressUID(value string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressUIDAndSalt(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressUIDAndSalt(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressUIDAndSendAddressUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressUIDAndSendAddressUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressUIDAndOutputsTotal(value1 string, value2 float64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressUIDAndOutputsTotal(value1 string, value2 float64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressUIDAndTime(value1 string, value2 int64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressUIDAndTime(value1 string, value2 int64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressUIDAndUID(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressUIDAndUID(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressUIDAndCreated(value1 string, value2 time.Time) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressUIDAndCreated(value1 string, value2 time.Time) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressUIDAndUser(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressUIDAndUser(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressUIDAndSendAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressUIDAndSendAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressUIDAndCurrency(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressUIDAndCurrency(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressUIDAndAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressUIDAndAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressUIDAndTxId(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressUIDAndTxId(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressUIDAndSender(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressUIDAndSender(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressUIDAndProject(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressUIDAndProject(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressUIDAndWallet(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressUIDAndWallet(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressUIDAndReceiver(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressUIDAndReceiver(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressUIDAndReceiveAddress(value1 string, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressUIDAndReceiveAddress(value1 string, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordByReceiveAddressUIDAndSize(value1 string, value2 int) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordByReceiveAddressUIDAndSize(value1 string, value2 int) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"ReceiveAddressUID\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) CountBillingRecordBySize(value int) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"BillingRecord\" WHERE \"Size\" = $1;", value)
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

func (db *DB) GetBillingRecordBySize(value int) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySize(value int) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySizeAndSalt(value1 int, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySizeAndSalt(value1 int, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySizeAndSendAddressUID(value1 int, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySizeAndSendAddressUID(value1 int, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySizeAndOutputsTotal(value1 int, value2 float64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySizeAndOutputsTotal(value1 int, value2 float64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySizeAndTime(value1 int, value2 int64) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySizeAndTime(value1 int, value2 int64) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySizeAndUID(value1 int, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySizeAndUID(value1 int, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySizeAndCreated(value1 int, value2 time.Time) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySizeAndCreated(value1 int, value2 time.Time) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySizeAndUser(value1 int, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySizeAndUser(value1 int, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySizeAndSendAddress(value1 int, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySizeAndSendAddress(value1 int, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySizeAndCurrency(value1 int, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySizeAndCurrency(value1 int, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySizeAndAddress(value1 int, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySizeAndAddress(value1 int, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySizeAndTxId(value1 int, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySizeAndTxId(value1 int, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySizeAndSender(value1 int, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySizeAndSender(value1 int, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySizeAndProject(value1 int, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySizeAndProject(value1 int, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySizeAndWallet(value1 int, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySizeAndWallet(value1 int, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySizeAndReceiver(value1 int, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySizeAndReceiver(value1 int, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySizeAndReceiveAddress(value1 int, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySizeAndReceiveAddress(value1 int, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}

func (db *DB) GetBillingRecordBySizeAndReceiveAddressUID(value1 int, value2 string) (bool, *models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToBillingRecord(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryBillingRecordBySizeAndReceiveAddressUID(value1 int, value2 string) ([]*models.BillingRecord, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"BillingRecord\" WHERE \"Size\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToBillingRecord(rows)
}


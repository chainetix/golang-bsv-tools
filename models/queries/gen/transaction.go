package database

import (
	"fmt"
	"time"
	"database/sql"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/project"
)

func (db *DB) CreateTableTransaction() error {
	q := "CREATE TABLE IF NOT EXISTS " + db.dbName + ".\"Transaction\" (\"SendAddress\" STRING NOT NULL, \"Currency\" STRING NOT NULL, \"Size\" INT NOT NULL, \"UID\" UUID PRIMARY KEY DEFAULT gen_random_uuid(), \"User\" STRING NOT NULL, \"Sender\" STRING NOT NULL, \"SendAddressUID\" STRING NOT NULL, \"ReceiveAddress\" STRING NOT NULL, \"OutputsTotal\" FLOAT NOT NULL, \"Salt\" UUID DEFAULT gen_random_uuid(), \"TxId\" STRING NOT NULL, \"Receiver\" STRING NOT NULL, \"Address\" STRING NOT NULL, \"ReceiveAddressUID\" STRING NOT NULL, \"Time\" int64 NOT NULL, \"Created\" TIMESTAMP DEFAULT current_timestamp(), \"Project\" STRING NOT NULL, \"Wallet\" STRING NOT NULL);"
	fmt.Println(q)
	_, err := db.DoExec(q)
	return err
}

func (db *DB) InsertTransaction(Address string, Currency string, OutputsTotal float64, Project string, ReceiveAddress string, ReceiveAddressUID string, Receiver string, SendAddress string, SendAddressUID string, Sender string, Size int, Time int64, TxId string, User string, Wallet string) (*models.Transaction, error) {
	row := &models.Transaction{}
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Transaction\" (\"Address\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"User\", \"Wallet\") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING Address, Currency, OutputsTotal, Project, ReceiveAddress, ReceiveAddressUID, Receiver, SendAddress, SendAddressUID, Sender, Size, Time, TxId, User, Wallet;", Address, Currency, OutputsTotal, Project, ReceiveAddress, ReceiveAddressUID, Receiver, SendAddress, SendAddressUID, Sender, Size, Time, TxId, User, Wallet).Scan(&row.Address, &row.Created, &row.Currency, &row.OutputsTotal, &row.Project, &row.ReceiveAddress, &row.ReceiveAddressUID, &row.Receiver, &row.Salt, &row.SendAddress, &row.SendAddressUID, &row.Sender, &row.Size, &row.Time, &row.TxId, &row.UID, &row.User, &row.Wallet); err != nil {
		return nil, err
	}
	return row, nil
}

func (db *DB) DeleteTransaction(uid string) error {
	_, err := db.DoQuery("DELETE FROM " + db.dbName + ".\"Transaction\" WHERE \"UID\" = $1;", uid)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) PutTransaction(model *models.Transaction) error {
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Transaction\" (\"Address\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"User\", \"Wallet\") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\";", model.Address, model.Currency, model.OutputsTotal, model.Project, model.ReceiveAddress, model.ReceiveAddressUID, model.Receiver, model.SendAddress, model.SendAddressUID, model.Sender, model.Size, model.Time, model.TxId, model.User, model.Wallet).Scan(&model.Address, &model.Created, &model.Currency, &model.OutputsTotal, &model.Project, &model.ReceiveAddress, &model.ReceiveAddressUID, &model.Receiver, &model.Salt, &model.SendAddress, &model.SendAddressUID, &model.Sender, &model.Size, &model.Time, &model.TxId, &model.UID, &model.User, &model.Wallet); err != nil {
		return err
	}
	return nil
}

func (db *DB) ScanToTransaction(rows *sql.Rows) ([]*models.Transaction, error) {
	defer rows.Close()
	results := []*models.Transaction{}
	for rows.Next() {
		result := &models.Transaction{}
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

func (db *DB) CountTransaction() (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM " + db.dbName + ".\"Transaction\";")
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

func (db *DB) QueryTransaction() ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\";")
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) CountTransactionBySendAddress(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Transaction\" WHERE \"SendAddress\" = $1;", value)
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

func (db *DB) GetTransactionBySendAddress(value string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddress(value string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressAndCurrency(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressAndCurrency(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressAndSize(value1 string, value2 int) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressAndSize(value1 string, value2 int) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressAndUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressAndUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressAndUser(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressAndUser(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressAndSender(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressAndSender(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressAndSendAddressUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressAndSendAddressUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressAndReceiveAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressAndReceiveAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressAndOutputsTotal(value1 string, value2 float64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressAndOutputsTotal(value1 string, value2 float64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressAndSalt(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressAndSalt(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressAndTxId(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressAndTxId(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressAndReceiver(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressAndReceiver(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressAndAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressAndAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressAndReceiveAddressUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressAndReceiveAddressUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressAndTime(value1 string, value2 int64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressAndTime(value1 string, value2 int64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressAndCreated(value1 string, value2 time.Time) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressAndCreated(value1 string, value2 time.Time) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressAndProject(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressAndProject(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressAndWallet(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressAndWallet(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddress\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) CountTransactionByCurrency(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Transaction\" WHERE \"Currency\" = $1;", value)
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

func (db *DB) GetTransactionByCurrency(value string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCurrency(value string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCurrencyAndSendAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCurrencyAndSendAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCurrencyAndSize(value1 string, value2 int) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCurrencyAndSize(value1 string, value2 int) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCurrencyAndUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCurrencyAndUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCurrencyAndUser(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCurrencyAndUser(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCurrencyAndSender(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCurrencyAndSender(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCurrencyAndSendAddressUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCurrencyAndSendAddressUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCurrencyAndReceiveAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCurrencyAndReceiveAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCurrencyAndOutputsTotal(value1 string, value2 float64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCurrencyAndOutputsTotal(value1 string, value2 float64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCurrencyAndSalt(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCurrencyAndSalt(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCurrencyAndTxId(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCurrencyAndTxId(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCurrencyAndReceiver(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCurrencyAndReceiver(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCurrencyAndAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCurrencyAndAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCurrencyAndReceiveAddressUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCurrencyAndReceiveAddressUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCurrencyAndTime(value1 string, value2 int64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCurrencyAndTime(value1 string, value2 int64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCurrencyAndCreated(value1 string, value2 time.Time) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCurrencyAndCreated(value1 string, value2 time.Time) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCurrencyAndProject(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCurrencyAndProject(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCurrencyAndWallet(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCurrencyAndWallet(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Currency\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) CountTransactionBySize(value int) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Transaction\" WHERE \"Size\" = $1;", value)
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

func (db *DB) GetTransactionBySize(value int) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySize(value int) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySizeAndSendAddress(value1 int, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySizeAndSendAddress(value1 int, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySizeAndCurrency(value1 int, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySizeAndCurrency(value1 int, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySizeAndUID(value1 int, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySizeAndUID(value1 int, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySizeAndUser(value1 int, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySizeAndUser(value1 int, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySizeAndSender(value1 int, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySizeAndSender(value1 int, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySizeAndSendAddressUID(value1 int, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySizeAndSendAddressUID(value1 int, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySizeAndReceiveAddress(value1 int, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySizeAndReceiveAddress(value1 int, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySizeAndOutputsTotal(value1 int, value2 float64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySizeAndOutputsTotal(value1 int, value2 float64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySizeAndSalt(value1 int, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySizeAndSalt(value1 int, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySizeAndTxId(value1 int, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySizeAndTxId(value1 int, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySizeAndReceiver(value1 int, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySizeAndReceiver(value1 int, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySizeAndAddress(value1 int, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySizeAndAddress(value1 int, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySizeAndReceiveAddressUID(value1 int, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySizeAndReceiveAddressUID(value1 int, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySizeAndTime(value1 int, value2 int64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySizeAndTime(value1 int, value2 int64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySizeAndCreated(value1 int, value2 time.Time) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySizeAndCreated(value1 int, value2 time.Time) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySizeAndProject(value1 int, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySizeAndProject(value1 int, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySizeAndWallet(value1 int, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySizeAndWallet(value1 int, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Size\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) CountTransactionByUID(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Transaction\" WHERE \"UID\" = $1;", value)
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

func (db *DB) GetTransactionByUID(value string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"UID\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByUID(value string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"UID\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) CountTransactionByUser(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Transaction\" WHERE \"User\" = $1;", value)
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

func (db *DB) GetTransactionByUser(value string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByUser(value string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByUserAndSendAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByUserAndSendAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByUserAndCurrency(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByUserAndCurrency(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByUserAndSize(value1 string, value2 int) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByUserAndSize(value1 string, value2 int) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByUserAndUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByUserAndUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByUserAndSender(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByUserAndSender(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByUserAndSendAddressUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByUserAndSendAddressUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByUserAndReceiveAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByUserAndReceiveAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByUserAndOutputsTotal(value1 string, value2 float64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByUserAndOutputsTotal(value1 string, value2 float64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByUserAndSalt(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByUserAndSalt(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByUserAndTxId(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByUserAndTxId(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByUserAndReceiver(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByUserAndReceiver(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByUserAndAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByUserAndAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByUserAndReceiveAddressUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByUserAndReceiveAddressUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByUserAndTime(value1 string, value2 int64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByUserAndTime(value1 string, value2 int64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByUserAndCreated(value1 string, value2 time.Time) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByUserAndCreated(value1 string, value2 time.Time) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByUserAndProject(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByUserAndProject(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByUserAndWallet(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByUserAndWallet(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"User\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) CountTransactionBySender(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Transaction\" WHERE \"Sender\" = $1;", value)
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

func (db *DB) GetTransactionBySender(value string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySender(value string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySenderAndSendAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySenderAndSendAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySenderAndCurrency(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySenderAndCurrency(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySenderAndSize(value1 string, value2 int) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySenderAndSize(value1 string, value2 int) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySenderAndUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySenderAndUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySenderAndUser(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySenderAndUser(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySenderAndSendAddressUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySenderAndSendAddressUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySenderAndReceiveAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySenderAndReceiveAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySenderAndOutputsTotal(value1 string, value2 float64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySenderAndOutputsTotal(value1 string, value2 float64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySenderAndSalt(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySenderAndSalt(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySenderAndTxId(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySenderAndTxId(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySenderAndReceiver(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySenderAndReceiver(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySenderAndAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySenderAndAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySenderAndReceiveAddressUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySenderAndReceiveAddressUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySenderAndTime(value1 string, value2 int64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySenderAndTime(value1 string, value2 int64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySenderAndCreated(value1 string, value2 time.Time) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySenderAndCreated(value1 string, value2 time.Time) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySenderAndProject(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySenderAndProject(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySenderAndWallet(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySenderAndWallet(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Sender\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) CountTransactionBySendAddressUID(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Transaction\" WHERE \"SendAddressUID\" = $1;", value)
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

func (db *DB) GetTransactionBySendAddressUID(value string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressUID(value string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressUIDAndSendAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressUIDAndSendAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressUIDAndCurrency(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressUIDAndCurrency(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressUIDAndSize(value1 string, value2 int) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressUIDAndSize(value1 string, value2 int) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressUIDAndUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressUIDAndUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressUIDAndUser(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressUIDAndUser(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressUIDAndSender(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressUIDAndSender(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressUIDAndReceiveAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressUIDAndReceiveAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressUIDAndOutputsTotal(value1 string, value2 float64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressUIDAndOutputsTotal(value1 string, value2 float64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressUIDAndSalt(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressUIDAndSalt(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressUIDAndTxId(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressUIDAndTxId(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressUIDAndReceiver(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressUIDAndReceiver(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressUIDAndAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressUIDAndAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressUIDAndReceiveAddressUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressUIDAndReceiveAddressUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressUIDAndTime(value1 string, value2 int64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressUIDAndTime(value1 string, value2 int64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressUIDAndCreated(value1 string, value2 time.Time) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressUIDAndCreated(value1 string, value2 time.Time) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressUIDAndProject(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressUIDAndProject(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySendAddressUIDAndWallet(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySendAddressUIDAndWallet(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"SendAddressUID\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) CountTransactionByReceiveAddress(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Transaction\" WHERE \"ReceiveAddress\" = $1;", value)
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

func (db *DB) GetTransactionByReceiveAddress(value string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddress(value string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressAndSendAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressAndSendAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressAndCurrency(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressAndCurrency(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressAndSize(value1 string, value2 int) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressAndSize(value1 string, value2 int) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressAndUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressAndUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressAndUser(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressAndUser(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressAndSender(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressAndSender(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressAndSendAddressUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressAndSendAddressUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressAndOutputsTotal(value1 string, value2 float64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressAndOutputsTotal(value1 string, value2 float64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressAndSalt(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressAndSalt(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressAndTxId(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressAndTxId(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressAndReceiver(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressAndReceiver(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressAndAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressAndAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressAndReceiveAddressUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressAndReceiveAddressUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressAndTime(value1 string, value2 int64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressAndTime(value1 string, value2 int64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressAndCreated(value1 string, value2 time.Time) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressAndCreated(value1 string, value2 time.Time) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressAndProject(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressAndProject(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressAndWallet(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressAndWallet(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddress\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) CountTransactionByOutputsTotal(value float64) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Transaction\" WHERE \"OutputsTotal\" = $1;", value)
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

func (db *DB) GetTransactionByOutputsTotal(value float64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByOutputsTotal(value float64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByOutputsTotalAndSendAddress(value1 float64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByOutputsTotalAndSendAddress(value1 float64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByOutputsTotalAndCurrency(value1 float64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByOutputsTotalAndCurrency(value1 float64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByOutputsTotalAndSize(value1 float64, value2 int) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByOutputsTotalAndSize(value1 float64, value2 int) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByOutputsTotalAndUID(value1 float64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByOutputsTotalAndUID(value1 float64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByOutputsTotalAndUser(value1 float64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByOutputsTotalAndUser(value1 float64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByOutputsTotalAndSender(value1 float64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByOutputsTotalAndSender(value1 float64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByOutputsTotalAndSendAddressUID(value1 float64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByOutputsTotalAndSendAddressUID(value1 float64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByOutputsTotalAndReceiveAddress(value1 float64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByOutputsTotalAndReceiveAddress(value1 float64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByOutputsTotalAndSalt(value1 float64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByOutputsTotalAndSalt(value1 float64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByOutputsTotalAndTxId(value1 float64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByOutputsTotalAndTxId(value1 float64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByOutputsTotalAndReceiver(value1 float64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByOutputsTotalAndReceiver(value1 float64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByOutputsTotalAndAddress(value1 float64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByOutputsTotalAndAddress(value1 float64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByOutputsTotalAndReceiveAddressUID(value1 float64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByOutputsTotalAndReceiveAddressUID(value1 float64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByOutputsTotalAndTime(value1 float64, value2 int64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByOutputsTotalAndTime(value1 float64, value2 int64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByOutputsTotalAndCreated(value1 float64, value2 time.Time) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByOutputsTotalAndCreated(value1 float64, value2 time.Time) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByOutputsTotalAndProject(value1 float64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByOutputsTotalAndProject(value1 float64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByOutputsTotalAndWallet(value1 float64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByOutputsTotalAndWallet(value1 float64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"OutputsTotal\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) CountTransactionBySalt(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Transaction\" WHERE \"Salt\" = $1;", value)
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

func (db *DB) GetTransactionBySalt(value string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySalt(value string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySaltAndSendAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySaltAndSendAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySaltAndCurrency(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySaltAndCurrency(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySaltAndSize(value1 string, value2 int) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySaltAndSize(value1 string, value2 int) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySaltAndUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySaltAndUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySaltAndUser(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySaltAndUser(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySaltAndSender(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySaltAndSender(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySaltAndSendAddressUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySaltAndSendAddressUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySaltAndReceiveAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySaltAndReceiveAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySaltAndOutputsTotal(value1 string, value2 float64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySaltAndOutputsTotal(value1 string, value2 float64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySaltAndTxId(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySaltAndTxId(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySaltAndReceiver(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySaltAndReceiver(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySaltAndAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySaltAndAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySaltAndReceiveAddressUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySaltAndReceiveAddressUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySaltAndTime(value1 string, value2 int64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySaltAndTime(value1 string, value2 int64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySaltAndCreated(value1 string, value2 time.Time) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySaltAndCreated(value1 string, value2 time.Time) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySaltAndProject(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySaltAndProject(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionBySaltAndWallet(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionBySaltAndWallet(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Salt\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) CountTransactionByTxId(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Transaction\" WHERE \"TxId\" = $1;", value)
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

func (db *DB) GetTransactionByTxId(value string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTxId(value string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTxIdAndSendAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTxIdAndSendAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTxIdAndCurrency(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTxIdAndCurrency(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTxIdAndSize(value1 string, value2 int) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTxIdAndSize(value1 string, value2 int) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTxIdAndUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTxIdAndUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTxIdAndUser(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTxIdAndUser(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTxIdAndSender(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTxIdAndSender(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTxIdAndSendAddressUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTxIdAndSendAddressUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTxIdAndReceiveAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTxIdAndReceiveAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTxIdAndOutputsTotal(value1 string, value2 float64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTxIdAndOutputsTotal(value1 string, value2 float64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTxIdAndSalt(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTxIdAndSalt(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTxIdAndReceiver(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTxIdAndReceiver(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTxIdAndAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTxIdAndAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTxIdAndReceiveAddressUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTxIdAndReceiveAddressUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTxIdAndTime(value1 string, value2 int64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTxIdAndTime(value1 string, value2 int64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTxIdAndCreated(value1 string, value2 time.Time) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTxIdAndCreated(value1 string, value2 time.Time) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTxIdAndProject(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTxIdAndProject(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTxIdAndWallet(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTxIdAndWallet(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"TxId\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) CountTransactionByReceiver(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Transaction\" WHERE \"Receiver\" = $1;", value)
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

func (db *DB) GetTransactionByReceiver(value string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiver(value string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiverAndSendAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiverAndSendAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiverAndCurrency(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiverAndCurrency(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiverAndSize(value1 string, value2 int) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiverAndSize(value1 string, value2 int) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiverAndUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiverAndUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiverAndUser(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiverAndUser(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiverAndSender(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiverAndSender(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiverAndSendAddressUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiverAndSendAddressUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiverAndReceiveAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiverAndReceiveAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiverAndOutputsTotal(value1 string, value2 float64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiverAndOutputsTotal(value1 string, value2 float64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiverAndSalt(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiverAndSalt(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiverAndTxId(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiverAndTxId(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiverAndAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiverAndAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiverAndReceiveAddressUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiverAndReceiveAddressUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiverAndTime(value1 string, value2 int64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiverAndTime(value1 string, value2 int64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiverAndCreated(value1 string, value2 time.Time) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiverAndCreated(value1 string, value2 time.Time) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiverAndProject(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiverAndProject(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiverAndWallet(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiverAndWallet(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Receiver\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) CountTransactionByAddress(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Transaction\" WHERE \"Address\" = $1;", value)
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

func (db *DB) GetTransactionByAddress(value string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByAddress(value string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByAddressAndSendAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByAddressAndSendAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByAddressAndCurrency(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByAddressAndCurrency(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByAddressAndSize(value1 string, value2 int) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByAddressAndSize(value1 string, value2 int) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByAddressAndUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByAddressAndUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByAddressAndUser(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByAddressAndUser(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByAddressAndSender(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByAddressAndSender(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByAddressAndSendAddressUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByAddressAndSendAddressUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByAddressAndReceiveAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByAddressAndReceiveAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByAddressAndOutputsTotal(value1 string, value2 float64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByAddressAndOutputsTotal(value1 string, value2 float64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByAddressAndSalt(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByAddressAndSalt(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByAddressAndTxId(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByAddressAndTxId(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByAddressAndReceiver(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByAddressAndReceiver(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByAddressAndReceiveAddressUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByAddressAndReceiveAddressUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByAddressAndTime(value1 string, value2 int64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByAddressAndTime(value1 string, value2 int64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByAddressAndCreated(value1 string, value2 time.Time) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByAddressAndCreated(value1 string, value2 time.Time) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByAddressAndProject(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByAddressAndProject(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByAddressAndWallet(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByAddressAndWallet(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Address\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) CountTransactionByReceiveAddressUID(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Transaction\" WHERE \"ReceiveAddressUID\" = $1;", value)
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

func (db *DB) GetTransactionByReceiveAddressUID(value string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressUID(value string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressUIDAndSendAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressUIDAndSendAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressUIDAndCurrency(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressUIDAndCurrency(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressUIDAndSize(value1 string, value2 int) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressUIDAndSize(value1 string, value2 int) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressUIDAndUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressUIDAndUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressUIDAndUser(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressUIDAndUser(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressUIDAndSender(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressUIDAndSender(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressUIDAndSendAddressUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressUIDAndSendAddressUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressUIDAndReceiveAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressUIDAndReceiveAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressUIDAndOutputsTotal(value1 string, value2 float64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressUIDAndOutputsTotal(value1 string, value2 float64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressUIDAndSalt(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressUIDAndSalt(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressUIDAndTxId(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressUIDAndTxId(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressUIDAndReceiver(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressUIDAndReceiver(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressUIDAndAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressUIDAndAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressUIDAndTime(value1 string, value2 int64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressUIDAndTime(value1 string, value2 int64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressUIDAndCreated(value1 string, value2 time.Time) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressUIDAndCreated(value1 string, value2 time.Time) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressUIDAndProject(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressUIDAndProject(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByReceiveAddressUIDAndWallet(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByReceiveAddressUIDAndWallet(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"ReceiveAddressUID\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) CountTransactionByTime(value int64) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Transaction\" WHERE \"Time\" = $1;", value)
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

func (db *DB) GetTransactionByTime(value int64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTime(value int64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTimeAndSendAddress(value1 int64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTimeAndSendAddress(value1 int64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTimeAndCurrency(value1 int64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTimeAndCurrency(value1 int64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTimeAndSize(value1 int64, value2 int) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTimeAndSize(value1 int64, value2 int) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTimeAndUID(value1 int64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTimeAndUID(value1 int64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTimeAndUser(value1 int64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTimeAndUser(value1 int64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTimeAndSender(value1 int64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTimeAndSender(value1 int64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTimeAndSendAddressUID(value1 int64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTimeAndSendAddressUID(value1 int64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTimeAndReceiveAddress(value1 int64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTimeAndReceiveAddress(value1 int64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTimeAndOutputsTotal(value1 int64, value2 float64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTimeAndOutputsTotal(value1 int64, value2 float64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTimeAndSalt(value1 int64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTimeAndSalt(value1 int64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTimeAndTxId(value1 int64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTimeAndTxId(value1 int64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTimeAndReceiver(value1 int64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTimeAndReceiver(value1 int64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTimeAndAddress(value1 int64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTimeAndAddress(value1 int64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTimeAndReceiveAddressUID(value1 int64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTimeAndReceiveAddressUID(value1 int64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTimeAndCreated(value1 int64, value2 time.Time) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTimeAndCreated(value1 int64, value2 time.Time) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTimeAndProject(value1 int64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTimeAndProject(value1 int64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByTimeAndWallet(value1 int64, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByTimeAndWallet(value1 int64, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Time\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) CountTransactionByCreated(value time.Time) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Transaction\" WHERE \"Created\" = $1;", value)
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

func (db *DB) GetTransactionByCreated(value time.Time) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCreated(value time.Time) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCreatedAndSendAddress(value1 time.Time, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCreatedAndSendAddress(value1 time.Time, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCreatedAndCurrency(value1 time.Time, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCreatedAndCurrency(value1 time.Time, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCreatedAndSize(value1 time.Time, value2 int) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCreatedAndSize(value1 time.Time, value2 int) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCreatedAndUID(value1 time.Time, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCreatedAndUID(value1 time.Time, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCreatedAndUser(value1 time.Time, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCreatedAndUser(value1 time.Time, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCreatedAndSender(value1 time.Time, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCreatedAndSender(value1 time.Time, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCreatedAndSendAddressUID(value1 time.Time, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCreatedAndSendAddressUID(value1 time.Time, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCreatedAndReceiveAddress(value1 time.Time, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCreatedAndReceiveAddress(value1 time.Time, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCreatedAndOutputsTotal(value1 time.Time, value2 float64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCreatedAndOutputsTotal(value1 time.Time, value2 float64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCreatedAndSalt(value1 time.Time, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCreatedAndSalt(value1 time.Time, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCreatedAndTxId(value1 time.Time, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCreatedAndTxId(value1 time.Time, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCreatedAndReceiver(value1 time.Time, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCreatedAndReceiver(value1 time.Time, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCreatedAndAddress(value1 time.Time, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCreatedAndAddress(value1 time.Time, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCreatedAndReceiveAddressUID(value1 time.Time, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCreatedAndReceiveAddressUID(value1 time.Time, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCreatedAndTime(value1 time.Time, value2 int64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCreatedAndTime(value1 time.Time, value2 int64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCreatedAndProject(value1 time.Time, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCreatedAndProject(value1 time.Time, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByCreatedAndWallet(value1 time.Time, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByCreatedAndWallet(value1 time.Time, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Created\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) CountTransactionByProject(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Transaction\" WHERE \"Project\" = $1;", value)
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

func (db *DB) GetTransactionByProject(value string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByProject(value string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByProjectAndSendAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByProjectAndSendAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByProjectAndCurrency(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByProjectAndCurrency(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByProjectAndSize(value1 string, value2 int) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByProjectAndSize(value1 string, value2 int) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByProjectAndUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByProjectAndUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByProjectAndUser(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByProjectAndUser(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByProjectAndSender(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByProjectAndSender(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByProjectAndSendAddressUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByProjectAndSendAddressUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByProjectAndReceiveAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByProjectAndReceiveAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByProjectAndOutputsTotal(value1 string, value2 float64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByProjectAndOutputsTotal(value1 string, value2 float64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByProjectAndSalt(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByProjectAndSalt(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByProjectAndTxId(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByProjectAndTxId(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByProjectAndReceiver(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByProjectAndReceiver(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByProjectAndAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByProjectAndAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByProjectAndReceiveAddressUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByProjectAndReceiveAddressUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByProjectAndTime(value1 string, value2 int64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByProjectAndTime(value1 string, value2 int64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByProjectAndCreated(value1 string, value2 time.Time) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByProjectAndCreated(value1 string, value2 time.Time) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByProjectAndWallet(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"Wallet\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByProjectAndWallet(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Project\" = $1 AND \"Wallet\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) CountTransactionByWallet(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Transaction\" WHERE \"Wallet\" = $1;", value)
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

func (db *DB) GetTransactionByWallet(value string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByWallet(value string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByWalletAndSendAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"SendAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByWalletAndSendAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"SendAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByWalletAndCurrency(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByWalletAndCurrency(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByWalletAndSize(value1 string, value2 int) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"Size\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByWalletAndSize(value1 string, value2 int) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"Size\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByWalletAndUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByWalletAndUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByWalletAndUser(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByWalletAndUser(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByWalletAndSender(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"Sender\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByWalletAndSender(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"Sender\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByWalletAndSendAddressUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"SendAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByWalletAndSendAddressUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"SendAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByWalletAndReceiveAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"ReceiveAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByWalletAndReceiveAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"ReceiveAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByWalletAndOutputsTotal(value1 string, value2 float64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"OutputsTotal\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByWalletAndOutputsTotal(value1 string, value2 float64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"OutputsTotal\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByWalletAndSalt(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByWalletAndSalt(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByWalletAndTxId(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"TxId\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByWalletAndTxId(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"TxId\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByWalletAndReceiver(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"Receiver\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByWalletAndReceiver(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"Receiver\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByWalletAndAddress(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"Address\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByWalletAndAddress(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"Address\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByWalletAndReceiveAddressUID(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"ReceiveAddressUID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByWalletAndReceiveAddressUID(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"ReceiveAddressUID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByWalletAndTime(value1 string, value2 int64) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"Time\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByWalletAndTime(value1 string, value2 int64) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"Time\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByWalletAndCreated(value1 string, value2 time.Time) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByWalletAndCreated(value1 string, value2 time.Time) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}

func (db *DB) GetTransactionByWalletAndProject(value1 string, value2 string) (bool, *models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToTransaction(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryTransactionByWalletAndProject(value1 string, value2 string) ([]*models.Transaction, error) {
	rows, err := db.DoQuery("SELECT \"Address\", \"Created\", \"Currency\", \"OutputsTotal\", \"Project\", \"ReceiveAddress\", \"ReceiveAddressUID\", \"Receiver\", \"Salt\", \"SendAddress\", \"SendAddressUID\", \"Sender\", \"Size\", \"Time\", \"TxId\", \"UID\", \"User\", \"Wallet\" FROM " + db.dbName + ".\"Transaction\" WHERE \"Wallet\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToTransaction(rows)
}


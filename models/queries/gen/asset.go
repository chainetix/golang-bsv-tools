package database

import (
	"fmt"
	"time"
	"database/sql"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/project"
)

func (db *DB) CreateTableAsset() error {
	q := "CREATE TABLE IF NOT EXISTS " + db.dbName + ".\"Asset\" (\"UID\" UUID PRIMARY KEY DEFAULT gen_random_uuid(), \"Salt\" UUID DEFAULT gen_random_uuid(), \"Created\" TIMESTAMP DEFAULT current_timestamp(), \"Project\" STRING NOT NULL, \"Currency\" STRING NOT NULL, \"Recipient\" STRING NOT NULL, \"Quantity\" FLOAT NOT NULL);"
	fmt.Println(q)
	_, err := db.DoExec(q)
	return err
}

func (db *DB) InsertAsset(Currency string, Project string, Quantity float64, Recipient string) (*models.Asset, error) {
	row := &models.Asset{}
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Asset\" (\"Currency\", \"Project\", \"Quantity\", \"Recipient\") VALUES ($1, $2, $3, $4) RETURNING Currency, Project, Quantity, Recipient;", Currency, Project, Quantity, Recipient).Scan(&row.Created, &row.Currency, &row.Project, &row.Quantity, &row.Recipient, &row.Salt, &row.UID); err != nil {
		return nil, err
	}
	return row, nil
}

func (db *DB) DeleteAsset(uid string) error {
	_, err := db.DoQuery("DELETE FROM " + db.dbName + ".\"Asset\" WHERE \"UID\" = $1;", uid)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) PutAsset(model *models.Asset) error {
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Asset\" (\"Currency\", \"Project\", \"Quantity\", \"Recipient\") VALUES ($1, $2, $3, $4) RETURNING \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\";", model.Currency, model.Project, model.Quantity, model.Recipient).Scan(&model.Created, &model.Currency, &model.Project, &model.Quantity, &model.Recipient, &model.Salt, &model.UID); err != nil {
		return err
	}
	return nil
}

func (db *DB) ScanToAsset(rows *sql.Rows) ([]*models.Asset, error) {
	defer rows.Close()
	results := []*models.Asset{}
	for rows.Next() {
		result := &models.Asset{}
		err := rows.Scan(&result.Created, &result.Currency, &result.Project, &result.Quantity, &result.Recipient, &result.Salt, &result.UID)
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

func (db *DB) CountAsset() (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM " + db.dbName + ".\"Asset\";")
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

func (db *DB) QueryAsset() ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\";")
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) CountAssetByUID(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Asset\" WHERE \"UID\" = $1;", value)
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

func (db *DB) GetAssetByUID(value string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"UID\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByUID(value string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"UID\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) CountAssetBySalt(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Asset\" WHERE \"Salt\" = $1;", value)
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

func (db *DB) GetAssetBySalt(value string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Salt\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetBySalt(value string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Salt\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetBySaltAndUID(value1 string, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Salt\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetBySaltAndUID(value1 string, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Salt\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetBySaltAndCreated(value1 string, value2 time.Time) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Salt\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetBySaltAndCreated(value1 string, value2 time.Time) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Salt\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetBySaltAndProject(value1 string, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Salt\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetBySaltAndProject(value1 string, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Salt\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetBySaltAndCurrency(value1 string, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Salt\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetBySaltAndCurrency(value1 string, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Salt\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetBySaltAndRecipient(value1 string, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Salt\" = $1 AND \"Recipient\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetBySaltAndRecipient(value1 string, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Salt\" = $1 AND \"Recipient\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetBySaltAndQuantity(value1 string, value2 float64) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Salt\" = $1 AND \"Quantity\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetBySaltAndQuantity(value1 string, value2 float64) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Salt\" = $1 AND \"Quantity\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) CountAssetByCreated(value time.Time) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Asset\" WHERE \"Created\" = $1;", value)
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

func (db *DB) GetAssetByCreated(value time.Time) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Created\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByCreated(value time.Time) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Created\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByCreatedAndUID(value1 time.Time, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Created\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByCreatedAndUID(value1 time.Time, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Created\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByCreatedAndSalt(value1 time.Time, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Created\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByCreatedAndSalt(value1 time.Time, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Created\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByCreatedAndProject(value1 time.Time, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Created\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByCreatedAndProject(value1 time.Time, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Created\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByCreatedAndCurrency(value1 time.Time, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Created\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByCreatedAndCurrency(value1 time.Time, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Created\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByCreatedAndRecipient(value1 time.Time, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Created\" = $1 AND \"Recipient\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByCreatedAndRecipient(value1 time.Time, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Created\" = $1 AND \"Recipient\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByCreatedAndQuantity(value1 time.Time, value2 float64) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Created\" = $1 AND \"Quantity\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByCreatedAndQuantity(value1 time.Time, value2 float64) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Created\" = $1 AND \"Quantity\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) CountAssetByProject(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Asset\" WHERE \"Project\" = $1;", value)
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

func (db *DB) GetAssetByProject(value string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Project\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByProject(value string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Project\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByProjectAndUID(value1 string, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Project\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByProjectAndUID(value1 string, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Project\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByProjectAndSalt(value1 string, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Project\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByProjectAndSalt(value1 string, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Project\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByProjectAndCreated(value1 string, value2 time.Time) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Project\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByProjectAndCreated(value1 string, value2 time.Time) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Project\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByProjectAndCurrency(value1 string, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Project\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByProjectAndCurrency(value1 string, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Project\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByProjectAndRecipient(value1 string, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Project\" = $1 AND \"Recipient\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByProjectAndRecipient(value1 string, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Project\" = $1 AND \"Recipient\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByProjectAndQuantity(value1 string, value2 float64) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Project\" = $1 AND \"Quantity\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByProjectAndQuantity(value1 string, value2 float64) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Project\" = $1 AND \"Quantity\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) CountAssetByCurrency(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Asset\" WHERE \"Currency\" = $1;", value)
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

func (db *DB) GetAssetByCurrency(value string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Currency\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByCurrency(value string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Currency\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByCurrencyAndUID(value1 string, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Currency\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByCurrencyAndUID(value1 string, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Currency\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByCurrencyAndSalt(value1 string, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Currency\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByCurrencyAndSalt(value1 string, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Currency\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByCurrencyAndCreated(value1 string, value2 time.Time) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Currency\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByCurrencyAndCreated(value1 string, value2 time.Time) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Currency\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByCurrencyAndProject(value1 string, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Currency\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByCurrencyAndProject(value1 string, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Currency\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByCurrencyAndRecipient(value1 string, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Currency\" = $1 AND \"Recipient\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByCurrencyAndRecipient(value1 string, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Currency\" = $1 AND \"Recipient\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByCurrencyAndQuantity(value1 string, value2 float64) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Currency\" = $1 AND \"Quantity\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByCurrencyAndQuantity(value1 string, value2 float64) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Currency\" = $1 AND \"Quantity\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) CountAssetByRecipient(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Asset\" WHERE \"Recipient\" = $1;", value)
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

func (db *DB) GetAssetByRecipient(value string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Recipient\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByRecipient(value string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Recipient\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByRecipientAndUID(value1 string, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Recipient\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByRecipientAndUID(value1 string, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Recipient\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByRecipientAndSalt(value1 string, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Recipient\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByRecipientAndSalt(value1 string, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Recipient\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByRecipientAndCreated(value1 string, value2 time.Time) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Recipient\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByRecipientAndCreated(value1 string, value2 time.Time) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Recipient\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByRecipientAndProject(value1 string, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Recipient\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByRecipientAndProject(value1 string, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Recipient\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByRecipientAndCurrency(value1 string, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Recipient\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByRecipientAndCurrency(value1 string, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Recipient\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByRecipientAndQuantity(value1 string, value2 float64) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Recipient\" = $1 AND \"Quantity\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByRecipientAndQuantity(value1 string, value2 float64) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Recipient\" = $1 AND \"Quantity\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) CountAssetByQuantity(value float64) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Asset\" WHERE \"Quantity\" = $1;", value)
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

func (db *DB) GetAssetByQuantity(value float64) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Quantity\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByQuantity(value float64) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Quantity\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByQuantityAndUID(value1 float64, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Quantity\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByQuantityAndUID(value1 float64, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Quantity\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByQuantityAndSalt(value1 float64, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Quantity\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByQuantityAndSalt(value1 float64, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Quantity\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByQuantityAndCreated(value1 float64, value2 time.Time) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Quantity\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByQuantityAndCreated(value1 float64, value2 time.Time) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Quantity\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByQuantityAndProject(value1 float64, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Quantity\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByQuantityAndProject(value1 float64, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Quantity\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByQuantityAndCurrency(value1 float64, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Quantity\" = $1 AND \"Currency\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByQuantityAndCurrency(value1 float64, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Quantity\" = $1 AND \"Currency\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}

func (db *DB) GetAssetByQuantityAndRecipient(value1 float64, value2 string) (bool, *models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Quantity\" = $1 AND \"Recipient\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAsset(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAssetByQuantityAndRecipient(value1 float64, value2 string) ([]*models.Asset, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Currency\", \"Project\", \"Quantity\", \"Recipient\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Asset\" WHERE \"Quantity\" = $1 AND \"Recipient\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAsset(rows)
}


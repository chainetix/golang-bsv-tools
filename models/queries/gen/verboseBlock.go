package database

import (
	"fmt"
	"time"
	"database/sql"
	"encoding/json"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/project"
)

func (db *DB) CreateTableVerboseBlock() error {
	q := "CREATE TABLE IF NOT EXISTS " + db.dbName + ".\"VerboseBlock\" (\"Hash\" STRING NOT NULL, \"Info\" JSON NOT NULL, \"UID\" UUID PRIMARY KEY DEFAULT gen_random_uuid(), \"Salt\" UUID DEFAULT gen_random_uuid(), \"Created\" TIMESTAMP DEFAULT current_timestamp());"
	fmt.Println(q)
	_, err := db.DoExec(q)
	return err
}

func (db *DB) InsertVerboseBlock(Hash string, Info map[string]interface {}) (*models.VerboseBlock, error) {
	row := &models.VerboseBlock{}
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"VerboseBlock\" (\"Hash\", \"Info\") VALUES ($1, $2) RETURNING Hash, Info;", Hash, Info).Scan(&row.Created, &row.Hash, &row.Info, &row.Salt, &row.UID); err != nil {
		return nil, err
	}
	return row, nil
}

func (db *DB) DeleteVerboseBlock(uid string) error {
	_, err := db.DoQuery("DELETE FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"UID\" = $1;", uid)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) PutVerboseBlock(model *models.VerboseBlock) error {
	info, _ := json.Marshal(model.Info)
	var s string
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"VerboseBlock\" (\"Hash\", \"Info\") VALUES ($1, $2) RETURNING \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\";", model.Hash, string(info)).Scan(&model.Created, &model.Hash, &s, &model.Salt, &model.UID); err != nil {
		return err
	}
	err := json.Unmarshal([]byte(s), &model.Info)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) ScanToVerboseBlock(rows *sql.Rows) ([]*models.VerboseBlock, error) {
	defer rows.Close()
	results := []*models.VerboseBlock{}
	for rows.Next() {
		result := &models.VerboseBlock{}
		var s string
		err := rows.Scan(&result.Created, &result.Hash, &s, &result.Salt, &result.UID)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(s), &result.Info)
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

func (db *DB) CountVerboseBlock() (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM " + db.dbName + ".\"VerboseBlock\";")
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

func (db *DB) QueryVerboseBlock() ([]*models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\";")
	if err != nil {
		return nil, err
	}
	return db.ScanToVerboseBlock(rows)
}

func (db *DB) CountVerboseBlockByHash(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"VerboseBlock\" WHERE \"Hash\" = $1;", value)
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

func (db *DB) GetVerboseBlockByHash(value string) (bool, *models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Hash\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToVerboseBlock(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryVerboseBlockByHash(value string) ([]*models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Hash\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToVerboseBlock(rows)
}

func (db *DB) GetVerboseBlockByHashAndInfo(value1 string, value2 map[string]interface {}) (bool, *models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Hash\" = $1 AND \"Info\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToVerboseBlock(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryVerboseBlockByHashAndInfo(value1 string, value2 map[string]interface {}) ([]*models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Hash\" = $1 AND \"Info\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToVerboseBlock(rows)
}

func (db *DB) GetVerboseBlockByHashAndUID(value1 string, value2 string) (bool, *models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Hash\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToVerboseBlock(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryVerboseBlockByHashAndUID(value1 string, value2 string) ([]*models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Hash\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToVerboseBlock(rows)
}

func (db *DB) GetVerboseBlockByHashAndSalt(value1 string, value2 string) (bool, *models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Hash\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToVerboseBlock(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryVerboseBlockByHashAndSalt(value1 string, value2 string) ([]*models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Hash\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToVerboseBlock(rows)
}

func (db *DB) GetVerboseBlockByHashAndCreated(value1 string, value2 time.Time) (bool, *models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Hash\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToVerboseBlock(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryVerboseBlockByHashAndCreated(value1 string, value2 time.Time) ([]*models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Hash\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToVerboseBlock(rows)
}

func (db *DB) CountVerboseBlockByInfo(value map[string]interface {}) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"VerboseBlock\" WHERE \"Info\" = $1;", value)
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

func (db *DB) GetVerboseBlockByInfo(value map[string]interface {}) (bool, *models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Info\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToVerboseBlock(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryVerboseBlockByInfo(value map[string]interface {}) ([]*models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Info\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToVerboseBlock(rows)
}

func (db *DB) GetVerboseBlockByInfoAndHash(value1 map[string]interface {}, value2 string) (bool, *models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Info\" = $1 AND \"Hash\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToVerboseBlock(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryVerboseBlockByInfoAndHash(value1 map[string]interface {}, value2 string) ([]*models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Info\" = $1 AND \"Hash\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToVerboseBlock(rows)
}

func (db *DB) GetVerboseBlockByInfoAndUID(value1 map[string]interface {}, value2 string) (bool, *models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Info\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToVerboseBlock(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryVerboseBlockByInfoAndUID(value1 map[string]interface {}, value2 string) ([]*models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Info\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToVerboseBlock(rows)
}

func (db *DB) GetVerboseBlockByInfoAndSalt(value1 map[string]interface {}, value2 string) (bool, *models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Info\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToVerboseBlock(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryVerboseBlockByInfoAndSalt(value1 map[string]interface {}, value2 string) ([]*models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Info\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToVerboseBlock(rows)
}

func (db *DB) GetVerboseBlockByInfoAndCreated(value1 map[string]interface {}, value2 time.Time) (bool, *models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Info\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToVerboseBlock(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryVerboseBlockByInfoAndCreated(value1 map[string]interface {}, value2 time.Time) ([]*models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Info\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToVerboseBlock(rows)
}

func (db *DB) CountVerboseBlockByUID(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"VerboseBlock\" WHERE \"UID\" = $1;", value)
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

func (db *DB) GetVerboseBlockByUID(value string) (bool, *models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"UID\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToVerboseBlock(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryVerboseBlockByUID(value string) ([]*models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"UID\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToVerboseBlock(rows)
}

func (db *DB) CountVerboseBlockBySalt(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"VerboseBlock\" WHERE \"Salt\" = $1;", value)
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

func (db *DB) GetVerboseBlockBySalt(value string) (bool, *models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Salt\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToVerboseBlock(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryVerboseBlockBySalt(value string) ([]*models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Salt\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToVerboseBlock(rows)
}

func (db *DB) GetVerboseBlockBySaltAndHash(value1 string, value2 string) (bool, *models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Salt\" = $1 AND \"Hash\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToVerboseBlock(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryVerboseBlockBySaltAndHash(value1 string, value2 string) ([]*models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Salt\" = $1 AND \"Hash\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToVerboseBlock(rows)
}

func (db *DB) GetVerboseBlockBySaltAndInfo(value1 string, value2 map[string]interface {}) (bool, *models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Salt\" = $1 AND \"Info\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToVerboseBlock(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryVerboseBlockBySaltAndInfo(value1 string, value2 map[string]interface {}) ([]*models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Salt\" = $1 AND \"Info\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToVerboseBlock(rows)
}

func (db *DB) GetVerboseBlockBySaltAndUID(value1 string, value2 string) (bool, *models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Salt\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToVerboseBlock(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryVerboseBlockBySaltAndUID(value1 string, value2 string) ([]*models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Salt\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToVerboseBlock(rows)
}

func (db *DB) GetVerboseBlockBySaltAndCreated(value1 string, value2 time.Time) (bool, *models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Salt\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToVerboseBlock(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryVerboseBlockBySaltAndCreated(value1 string, value2 time.Time) ([]*models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Salt\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToVerboseBlock(rows)
}

func (db *DB) CountVerboseBlockByCreated(value time.Time) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"VerboseBlock\" WHERE \"Created\" = $1;", value)
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

func (db *DB) GetVerboseBlockByCreated(value time.Time) (bool, *models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Created\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToVerboseBlock(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryVerboseBlockByCreated(value time.Time) ([]*models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Created\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToVerboseBlock(rows)
}

func (db *DB) GetVerboseBlockByCreatedAndHash(value1 time.Time, value2 string) (bool, *models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Created\" = $1 AND \"Hash\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToVerboseBlock(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryVerboseBlockByCreatedAndHash(value1 time.Time, value2 string) ([]*models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Created\" = $1 AND \"Hash\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToVerboseBlock(rows)
}

func (db *DB) GetVerboseBlockByCreatedAndInfo(value1 time.Time, value2 map[string]interface {}) (bool, *models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Created\" = $1 AND \"Info\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToVerboseBlock(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryVerboseBlockByCreatedAndInfo(value1 time.Time, value2 map[string]interface {}) ([]*models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Created\" = $1 AND \"Info\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToVerboseBlock(rows)
}

func (db *DB) GetVerboseBlockByCreatedAndUID(value1 time.Time, value2 string) (bool, *models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Created\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToVerboseBlock(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryVerboseBlockByCreatedAndUID(value1 time.Time, value2 string) ([]*models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Created\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToVerboseBlock(rows)
}

func (db *DB) GetVerboseBlockByCreatedAndSalt(value1 time.Time, value2 string) (bool, *models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Created\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToVerboseBlock(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryVerboseBlockByCreatedAndSalt(value1 time.Time, value2 string) ([]*models.VerboseBlock, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Hash\", \"Info\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"VerboseBlock\" WHERE \"Created\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToVerboseBlock(rows)
}

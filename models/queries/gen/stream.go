package database

import (
	"fmt"
	"time"
	"database/sql"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/project"
)

func (db *DB) CreateTableStream() error {
	q := "CREATE TABLE IF NOT EXISTS " + db.dbName + ".\"Stream\" (\"Project\" STRING NOT NULL, \"PublicKey\" STRING NOT NULL, \"Label\" STRING NOT NULL, \"UID\" UUID PRIMARY KEY DEFAULT gen_random_uuid(), \"Salt\" UUID DEFAULT gen_random_uuid(), \"Created\" TIMESTAMP DEFAULT current_timestamp());"
	fmt.Println(q)
	_, err := db.DoExec(q)
	return err
}

func (db *DB) InsertStream(Label string, Project string, PublicKey string) (*models.Stream, error) {
	row := &models.Stream{}
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Stream\" (\"Label\", \"Project\", \"PublicKey\") VALUES ($1, $2, $3) RETURNING Label, Project, PublicKey;", Label, Project, PublicKey).Scan(&row.Created, &row.Label, &row.Project, &row.PublicKey, &row.Salt, &row.UID); err != nil {
		return nil, err
	}
	return row, nil
}

func (db *DB) DeleteStream(uid string) error {
	_, err := db.DoQuery("DELETE FROM " + db.dbName + ".\"Stream\" WHERE \"UID\" = $1;", uid)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) PutStream(model *models.Stream) error {
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Stream\" (\"Label\", \"Project\", \"PublicKey\") VALUES ($1, $2, $3) RETURNING \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\";", model.Label, model.Project, model.PublicKey).Scan(&model.Created, &model.Label, &model.Project, &model.PublicKey, &model.Salt, &model.UID); err != nil {
		return err
	}
	return nil
}

func (db *DB) ScanToStream(rows *sql.Rows) ([]*models.Stream, error) {
	defer rows.Close()
	results := []*models.Stream{}
	for rows.Next() {
		result := &models.Stream{}
		err := rows.Scan(&result.Created, &result.Label, &result.Project, &result.PublicKey, &result.Salt, &result.UID)
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

func (db *DB) CountStream() (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM " + db.dbName + ".\"Stream\";")
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

func (db *DB) QueryStream() ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\";")
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) CountStreamByProject(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Stream\" WHERE \"Project\" = $1;", value)
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

func (db *DB) GetStreamByProject(value string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Project\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByProject(value string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Project\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamByProjectAndPublicKey(value1 string, value2 string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Project\" = $1 AND \"PublicKey\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByProjectAndPublicKey(value1 string, value2 string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Project\" = $1 AND \"PublicKey\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamByProjectAndLabel(value1 string, value2 string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Project\" = $1 AND \"Label\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByProjectAndLabel(value1 string, value2 string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Project\" = $1 AND \"Label\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamByProjectAndUID(value1 string, value2 string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Project\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByProjectAndUID(value1 string, value2 string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Project\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamByProjectAndSalt(value1 string, value2 string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Project\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByProjectAndSalt(value1 string, value2 string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Project\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamByProjectAndCreated(value1 string, value2 time.Time) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Project\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByProjectAndCreated(value1 string, value2 time.Time) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Project\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) CountStreamByPublicKey(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Stream\" WHERE \"PublicKey\" = $1;", value)
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

func (db *DB) GetStreamByPublicKey(value string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"PublicKey\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByPublicKey(value string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"PublicKey\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamByPublicKeyAndProject(value1 string, value2 string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"PublicKey\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByPublicKeyAndProject(value1 string, value2 string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"PublicKey\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamByPublicKeyAndLabel(value1 string, value2 string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"PublicKey\" = $1 AND \"Label\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByPublicKeyAndLabel(value1 string, value2 string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"PublicKey\" = $1 AND \"Label\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamByPublicKeyAndUID(value1 string, value2 string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"PublicKey\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByPublicKeyAndUID(value1 string, value2 string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"PublicKey\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamByPublicKeyAndSalt(value1 string, value2 string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"PublicKey\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByPublicKeyAndSalt(value1 string, value2 string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"PublicKey\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamByPublicKeyAndCreated(value1 string, value2 time.Time) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"PublicKey\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByPublicKeyAndCreated(value1 string, value2 time.Time) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"PublicKey\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) CountStreamByLabel(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Stream\" WHERE \"Label\" = $1;", value)
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

func (db *DB) GetStreamByLabel(value string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Label\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByLabel(value string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Label\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamByLabelAndProject(value1 string, value2 string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Label\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByLabelAndProject(value1 string, value2 string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Label\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamByLabelAndPublicKey(value1 string, value2 string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Label\" = $1 AND \"PublicKey\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByLabelAndPublicKey(value1 string, value2 string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Label\" = $1 AND \"PublicKey\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamByLabelAndUID(value1 string, value2 string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Label\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByLabelAndUID(value1 string, value2 string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Label\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamByLabelAndSalt(value1 string, value2 string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Label\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByLabelAndSalt(value1 string, value2 string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Label\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamByLabelAndCreated(value1 string, value2 time.Time) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Label\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByLabelAndCreated(value1 string, value2 time.Time) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Label\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) CountStreamByUID(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Stream\" WHERE \"UID\" = $1;", value)
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

func (db *DB) GetStreamByUID(value string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"UID\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByUID(value string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"UID\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) CountStreamBySalt(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Stream\" WHERE \"Salt\" = $1;", value)
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

func (db *DB) GetStreamBySalt(value string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Salt\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamBySalt(value string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Salt\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamBySaltAndProject(value1 string, value2 string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Salt\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamBySaltAndProject(value1 string, value2 string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Salt\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamBySaltAndPublicKey(value1 string, value2 string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Salt\" = $1 AND \"PublicKey\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamBySaltAndPublicKey(value1 string, value2 string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Salt\" = $1 AND \"PublicKey\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamBySaltAndLabel(value1 string, value2 string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Salt\" = $1 AND \"Label\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamBySaltAndLabel(value1 string, value2 string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Salt\" = $1 AND \"Label\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamBySaltAndUID(value1 string, value2 string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Salt\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamBySaltAndUID(value1 string, value2 string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Salt\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamBySaltAndCreated(value1 string, value2 time.Time) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Salt\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamBySaltAndCreated(value1 string, value2 time.Time) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Salt\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) CountStreamByCreated(value time.Time) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Stream\" WHERE \"Created\" = $1;", value)
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

func (db *DB) GetStreamByCreated(value time.Time) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Created\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByCreated(value time.Time) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Created\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamByCreatedAndProject(value1 time.Time, value2 string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Created\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByCreatedAndProject(value1 time.Time, value2 string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Created\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamByCreatedAndPublicKey(value1 time.Time, value2 string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Created\" = $1 AND \"PublicKey\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByCreatedAndPublicKey(value1 time.Time, value2 string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Created\" = $1 AND \"PublicKey\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamByCreatedAndLabel(value1 time.Time, value2 string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Created\" = $1 AND \"Label\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByCreatedAndLabel(value1 time.Time, value2 string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Created\" = $1 AND \"Label\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamByCreatedAndUID(value1 time.Time, value2 string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Created\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByCreatedAndUID(value1 time.Time, value2 string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Created\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}

func (db *DB) GetStreamByCreatedAndSalt(value1 time.Time, value2 string) (bool, *models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Created\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToStream(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryStreamByCreatedAndSalt(value1 time.Time, value2 string) ([]*models.Stream, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"PublicKey\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Stream\" WHERE \"Created\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToStream(rows)
}


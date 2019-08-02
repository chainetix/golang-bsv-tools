package database

import (
	"fmt"
	"time"
	"database/sql"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/project"
)

func (db *DB) CreateTableApiToken() error {
	q := "CREATE TABLE IF NOT EXISTS " + db.dbName + ".\"ApiToken\" (\"UID\" UUID PRIMARY KEY DEFAULT gen_random_uuid(), \"Salt\" UUID DEFAULT gen_random_uuid(), \"Created\" TIMESTAMP DEFAULT current_timestamp(), \"Project\" STRING NOT NULL, \"Token\" STRING NOT NULL, \"Digest\" STRING NOT NULL, \"Resources\" STRING NOT NULL);"
	fmt.Println(q)
	_, err := db.DoExec(q)
	return err
}

func (db *DB) InsertApiToken(Digest string, Project string, Resources string, Token string) (*models.ApiToken, error) {
	row := &models.ApiToken{}
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"ApiToken\" (\"Digest\", \"Project\", \"Resources\", \"Token\") VALUES ($1, $2, $3, $4) RETURNING Digest, Project, Resources, Token;", Digest, Project, Resources, Token).Scan(&row.Created, &row.Digest, &row.Project, &row.Resources, &row.Salt, &row.Token, &row.UID); err != nil {
		return nil, err
	}
	return row, nil
}

func (db *DB) DeleteApiToken(uid string) error {
	_, err := db.DoQuery("DELETE FROM " + db.dbName + ".\"ApiToken\" WHERE \"UID\" = $1;", uid)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) PutApiToken(model *models.ApiToken) error {
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"ApiToken\" (\"Digest\", \"Project\", \"Resources\", \"Token\") VALUES ($1, $2, $3, $4) RETURNING \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\";", model.Digest, model.Project, model.Resources, model.Token).Scan(&model.Created, &model.Digest, &model.Project, &model.Resources, &model.Salt, &model.Token, &model.UID); err != nil {
		return err
	}
	return nil
}

func (db *DB) ScanToApiToken(rows *sql.Rows) ([]*models.ApiToken, error) {
	defer rows.Close()
	results := []*models.ApiToken{}
	for rows.Next() {
		result := &models.ApiToken{}
		err := rows.Scan(&result.Created, &result.Digest, &result.Project, &result.Resources, &result.Salt, &result.Token, &result.UID)
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

func (db *DB) CountApiToken() (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM " + db.dbName + ".\"ApiToken\";")
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

func (db *DB) QueryApiToken() ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\";")
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) CountApiTokenByUID(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"ApiToken\" WHERE \"UID\" = $1;", value)
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

func (db *DB) GetApiTokenByUID(value string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"UID\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByUID(value string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"UID\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) CountApiTokenBySalt(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"ApiToken\" WHERE \"Salt\" = $1;", value)
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

func (db *DB) GetApiTokenBySalt(value string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Salt\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenBySalt(value string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Salt\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenBySaltAndUID(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Salt\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenBySaltAndUID(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Salt\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenBySaltAndCreated(value1 string, value2 time.Time) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Salt\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenBySaltAndCreated(value1 string, value2 time.Time) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Salt\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenBySaltAndProject(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Salt\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenBySaltAndProject(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Salt\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenBySaltAndToken(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Salt\" = $1 AND \"Token\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenBySaltAndToken(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Salt\" = $1 AND \"Token\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenBySaltAndDigest(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Salt\" = $1 AND \"Digest\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenBySaltAndDigest(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Salt\" = $1 AND \"Digest\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenBySaltAndResources(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Salt\" = $1 AND \"Resources\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenBySaltAndResources(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Salt\" = $1 AND \"Resources\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) CountApiTokenByCreated(value time.Time) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"ApiToken\" WHERE \"Created\" = $1;", value)
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

func (db *DB) GetApiTokenByCreated(value time.Time) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Created\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByCreated(value time.Time) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Created\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByCreatedAndUID(value1 time.Time, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Created\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByCreatedAndUID(value1 time.Time, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Created\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByCreatedAndSalt(value1 time.Time, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Created\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByCreatedAndSalt(value1 time.Time, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Created\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByCreatedAndProject(value1 time.Time, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Created\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByCreatedAndProject(value1 time.Time, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Created\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByCreatedAndToken(value1 time.Time, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Created\" = $1 AND \"Token\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByCreatedAndToken(value1 time.Time, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Created\" = $1 AND \"Token\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByCreatedAndDigest(value1 time.Time, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Created\" = $1 AND \"Digest\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByCreatedAndDigest(value1 time.Time, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Created\" = $1 AND \"Digest\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByCreatedAndResources(value1 time.Time, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Created\" = $1 AND \"Resources\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByCreatedAndResources(value1 time.Time, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Created\" = $1 AND \"Resources\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) CountApiTokenByProject(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"ApiToken\" WHERE \"Project\" = $1;", value)
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

func (db *DB) GetApiTokenByProject(value string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Project\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByProject(value string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Project\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByProjectAndUID(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Project\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByProjectAndUID(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Project\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByProjectAndSalt(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Project\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByProjectAndSalt(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Project\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByProjectAndCreated(value1 string, value2 time.Time) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Project\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByProjectAndCreated(value1 string, value2 time.Time) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Project\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByProjectAndToken(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Project\" = $1 AND \"Token\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByProjectAndToken(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Project\" = $1 AND \"Token\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByProjectAndDigest(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Project\" = $1 AND \"Digest\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByProjectAndDigest(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Project\" = $1 AND \"Digest\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByProjectAndResources(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Project\" = $1 AND \"Resources\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByProjectAndResources(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Project\" = $1 AND \"Resources\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) CountApiTokenByToken(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"ApiToken\" WHERE \"Token\" = $1;", value)
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

func (db *DB) GetApiTokenByToken(value string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Token\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByToken(value string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Token\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByTokenAndUID(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Token\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByTokenAndUID(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Token\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByTokenAndSalt(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Token\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByTokenAndSalt(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Token\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByTokenAndCreated(value1 string, value2 time.Time) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Token\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByTokenAndCreated(value1 string, value2 time.Time) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Token\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByTokenAndProject(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Token\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByTokenAndProject(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Token\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByTokenAndDigest(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Token\" = $1 AND \"Digest\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByTokenAndDigest(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Token\" = $1 AND \"Digest\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByTokenAndResources(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Token\" = $1 AND \"Resources\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByTokenAndResources(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Token\" = $1 AND \"Resources\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) CountApiTokenByDigest(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"ApiToken\" WHERE \"Digest\" = $1;", value)
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

func (db *DB) GetApiTokenByDigest(value string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Digest\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByDigest(value string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Digest\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByDigestAndUID(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Digest\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByDigestAndUID(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Digest\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByDigestAndSalt(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Digest\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByDigestAndSalt(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Digest\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByDigestAndCreated(value1 string, value2 time.Time) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Digest\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByDigestAndCreated(value1 string, value2 time.Time) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Digest\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByDigestAndProject(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Digest\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByDigestAndProject(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Digest\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByDigestAndToken(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Digest\" = $1 AND \"Token\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByDigestAndToken(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Digest\" = $1 AND \"Token\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByDigestAndResources(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Digest\" = $1 AND \"Resources\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByDigestAndResources(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Digest\" = $1 AND \"Resources\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) CountApiTokenByResources(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"ApiToken\" WHERE \"Resources\" = $1;", value)
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

func (db *DB) GetApiTokenByResources(value string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Resources\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByResources(value string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Resources\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByResourcesAndUID(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Resources\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByResourcesAndUID(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Resources\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByResourcesAndSalt(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Resources\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByResourcesAndSalt(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Resources\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByResourcesAndCreated(value1 string, value2 time.Time) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Resources\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByResourcesAndCreated(value1 string, value2 time.Time) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Resources\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByResourcesAndProject(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Resources\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByResourcesAndProject(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Resources\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByResourcesAndToken(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Resources\" = $1 AND \"Token\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByResourcesAndToken(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Resources\" = $1 AND \"Token\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}

func (db *DB) GetApiTokenByResourcesAndDigest(value1 string, value2 string) (bool, *models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Resources\" = $1 AND \"Digest\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToApiToken(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryApiTokenByResourcesAndDigest(value1 string, value2 string) ([]*models.ApiToken, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Digest\", \"Project\", \"Resources\", \"Salt\", \"Token\", \"UID\" FROM " + db.dbName + ".\"ApiToken\" WHERE \"Resources\" = $1 AND \"Digest\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToApiToken(rows)
}


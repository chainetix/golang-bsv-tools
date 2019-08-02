package database

import (
	"fmt"
	"time"
	"database/sql"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/project"
)

func (db *DB) CreateTableUser() error {
	q := "CREATE TABLE IF NOT EXISTS " + db.dbName + ".\"User\" (\"Salt\" UUID DEFAULT gen_random_uuid(), \"Created\" TIMESTAMP DEFAULT current_timestamp(), \"Project\" STRING NOT NULL, \"Agent\" BOOL NOT NULL, \"Label\" STRING NOT NULL, \"UID\" UUID PRIMARY KEY DEFAULT gen_random_uuid());"
	fmt.Println(q)
	_, err := db.DoExec(q)
	return err
}

func (db *DB) InsertUser(Agent bool, Label string, Project string) (*models.User, error) {
	row := &models.User{}
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"User\" (\"Agent\", \"Label\", \"Project\") VALUES ($1, $2, $3) RETURNING Agent, Label, Project;", Agent, Label, Project).Scan(&row.Agent, &row.Created, &row.Label, &row.Project, &row.Salt, &row.UID); err != nil {
		return nil, err
	}
	return row, nil
}

func (db *DB) DeleteUser(uid string) error {
	_, err := db.DoQuery("DELETE FROM " + db.dbName + ".\"User\" WHERE \"UID\" = $1;", uid)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) PutUser(model *models.User) error {
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"User\" (\"Agent\", \"Label\", \"Project\") VALUES ($1, $2, $3) RETURNING \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\";", model.Agent, model.Label, model.Project).Scan(&model.Agent, &model.Created, &model.Label, &model.Project, &model.Salt, &model.UID); err != nil {
		return err
	}
	return nil
}

func (db *DB) ScanToUser(rows *sql.Rows) ([]*models.User, error) {
	defer rows.Close()
	results := []*models.User{}
	for rows.Next() {
		result := &models.User{}
		err := rows.Scan(&result.Agent, &result.Created, &result.Label, &result.Project, &result.Salt, &result.UID)
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

func (db *DB) CountUser() (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM " + db.dbName + ".\"User\";")
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

func (db *DB) QueryUser() ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\";")
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) CountUserBySalt(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"User\" WHERE \"Salt\" = $1;", value)
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

func (db *DB) GetUserBySalt(value string) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Salt\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserBySalt(value string) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Salt\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserBySaltAndCreated(value1 string, value2 time.Time) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Salt\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserBySaltAndCreated(value1 string, value2 time.Time) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Salt\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserBySaltAndProject(value1 string, value2 string) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Salt\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserBySaltAndProject(value1 string, value2 string) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Salt\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserBySaltAndAgent(value1 string, value2 bool) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Salt\" = $1 AND \"Agent\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserBySaltAndAgent(value1 string, value2 bool) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Salt\" = $1 AND \"Agent\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserBySaltAndLabel(value1 string, value2 string) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Salt\" = $1 AND \"Label\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserBySaltAndLabel(value1 string, value2 string) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Salt\" = $1 AND \"Label\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserBySaltAndUID(value1 string, value2 string) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Salt\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserBySaltAndUID(value1 string, value2 string) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Salt\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) CountUserByCreated(value time.Time) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"User\" WHERE \"Created\" = $1;", value)
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

func (db *DB) GetUserByCreated(value time.Time) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Created\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByCreated(value time.Time) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Created\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserByCreatedAndSalt(value1 time.Time, value2 string) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Created\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByCreatedAndSalt(value1 time.Time, value2 string) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Created\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserByCreatedAndProject(value1 time.Time, value2 string) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Created\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByCreatedAndProject(value1 time.Time, value2 string) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Created\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserByCreatedAndAgent(value1 time.Time, value2 bool) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Created\" = $1 AND \"Agent\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByCreatedAndAgent(value1 time.Time, value2 bool) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Created\" = $1 AND \"Agent\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserByCreatedAndLabel(value1 time.Time, value2 string) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Created\" = $1 AND \"Label\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByCreatedAndLabel(value1 time.Time, value2 string) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Created\" = $1 AND \"Label\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserByCreatedAndUID(value1 time.Time, value2 string) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Created\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByCreatedAndUID(value1 time.Time, value2 string) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Created\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) CountUserByProject(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"User\" WHERE \"Project\" = $1;", value)
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

func (db *DB) GetUserByProject(value string) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Project\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByProject(value string) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Project\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserByProjectAndSalt(value1 string, value2 string) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Project\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByProjectAndSalt(value1 string, value2 string) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Project\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserByProjectAndCreated(value1 string, value2 time.Time) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Project\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByProjectAndCreated(value1 string, value2 time.Time) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Project\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserByProjectAndAgent(value1 string, value2 bool) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Project\" = $1 AND \"Agent\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByProjectAndAgent(value1 string, value2 bool) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Project\" = $1 AND \"Agent\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserByProjectAndLabel(value1 string, value2 string) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Project\" = $1 AND \"Label\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByProjectAndLabel(value1 string, value2 string) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Project\" = $1 AND \"Label\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserByProjectAndUID(value1 string, value2 string) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Project\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByProjectAndUID(value1 string, value2 string) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Project\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) CountUserByAgent(value bool) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"User\" WHERE \"Agent\" = $1;", value)
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

func (db *DB) GetUserByAgent(value bool) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Agent\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByAgent(value bool) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Agent\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserByAgentAndSalt(value1 bool, value2 string) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Agent\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByAgentAndSalt(value1 bool, value2 string) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Agent\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserByAgentAndCreated(value1 bool, value2 time.Time) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Agent\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByAgentAndCreated(value1 bool, value2 time.Time) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Agent\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserByAgentAndProject(value1 bool, value2 string) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Agent\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByAgentAndProject(value1 bool, value2 string) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Agent\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserByAgentAndLabel(value1 bool, value2 string) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Agent\" = $1 AND \"Label\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByAgentAndLabel(value1 bool, value2 string) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Agent\" = $1 AND \"Label\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserByAgentAndUID(value1 bool, value2 string) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Agent\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByAgentAndUID(value1 bool, value2 string) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Agent\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) CountUserByLabel(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"User\" WHERE \"Label\" = $1;", value)
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

func (db *DB) GetUserByLabel(value string) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Label\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByLabel(value string) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Label\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserByLabelAndSalt(value1 string, value2 string) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Label\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByLabelAndSalt(value1 string, value2 string) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Label\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserByLabelAndCreated(value1 string, value2 time.Time) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Label\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByLabelAndCreated(value1 string, value2 time.Time) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Label\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserByLabelAndProject(value1 string, value2 string) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Label\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByLabelAndProject(value1 string, value2 string) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Label\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserByLabelAndAgent(value1 string, value2 bool) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Label\" = $1 AND \"Agent\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByLabelAndAgent(value1 string, value2 bool) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Label\" = $1 AND \"Agent\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) GetUserByLabelAndUID(value1 string, value2 string) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Label\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByLabelAndUID(value1 string, value2 string) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"Label\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}

func (db *DB) CountUserByUID(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"User\" WHERE \"UID\" = $1;", value)
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

func (db *DB) GetUserByUID(value string) (bool, *models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"UID\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUser(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserByUID(value string) ([]*models.User, error) {
	rows, err := db.DoQuery("SELECT \"Agent\", \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"User\" WHERE \"UID\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToUser(rows)
}


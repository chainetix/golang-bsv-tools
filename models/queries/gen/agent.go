package database

import (
	"fmt"
	"time"
	"database/sql"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/project"
)

func (db *DB) CreateTableAgent() error {
	q := "CREATE TABLE IF NOT EXISTS " + db.dbName + ".\"Agent\" (\"UID\" UUID PRIMARY KEY DEFAULT gen_random_uuid(), \"Salt\" UUID DEFAULT gen_random_uuid(), \"Created\" TIMESTAMP DEFAULT current_timestamp(), \"Project\" STRING NOT NULL, \"Label\" STRING NOT NULL);"
	fmt.Println(q)
	_, err := db.DoExec(q)
	return err
}

func (db *DB) InsertAgent(Label string, Project string) (*models.Agent, error) {
	row := &models.Agent{}
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Agent\" (\"Label\", \"Project\") VALUES ($1, $2) RETURNING Label, Project;", Label, Project).Scan(&row.Created, &row.Label, &row.Project, &row.Salt, &row.UID); err != nil {
		return nil, err
	}
	return row, nil
}

func (db *DB) DeleteAgent(uid string) error {
	_, err := db.DoQuery("DELETE FROM " + db.dbName + ".\"Agent\" WHERE \"UID\" = $1;", uid)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) PutAgent(model *models.Agent) error {
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Agent\" (\"Label\", \"Project\") VALUES ($1, $2) RETURNING \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\";", model.Label, model.Project).Scan(&model.Created, &model.Label, &model.Project, &model.Salt, &model.UID); err != nil {
		return err
	}
	return nil
}

func (db *DB) ScanToAgent(rows *sql.Rows) ([]*models.Agent, error) {
	defer rows.Close()
	results := []*models.Agent{}
	for rows.Next() {
		result := &models.Agent{}
		err := rows.Scan(&result.Created, &result.Label, &result.Project, &result.Salt, &result.UID)
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

func (db *DB) CountAgent() (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM " + db.dbName + ".\"Agent\";")
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

func (db *DB) QueryAgent() ([]*models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\";")
	if err != nil {
		return nil, err
	}
	return db.ScanToAgent(rows)
}

func (db *DB) CountAgentByUID(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Agent\" WHERE \"UID\" = $1;", value)
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

func (db *DB) GetAgentByUID(value string) (bool, *models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"UID\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAgent(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAgentByUID(value string) ([]*models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"UID\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToAgent(rows)
}

func (db *DB) CountAgentBySalt(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Agent\" WHERE \"Salt\" = $1;", value)
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

func (db *DB) GetAgentBySalt(value string) (bool, *models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Salt\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAgent(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAgentBySalt(value string) ([]*models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Salt\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToAgent(rows)
}

func (db *DB) GetAgentBySaltAndUID(value1 string, value2 string) (bool, *models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Salt\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAgent(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAgentBySaltAndUID(value1 string, value2 string) ([]*models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Salt\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAgent(rows)
}

func (db *DB) GetAgentBySaltAndCreated(value1 string, value2 time.Time) (bool, *models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Salt\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAgent(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAgentBySaltAndCreated(value1 string, value2 time.Time) ([]*models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Salt\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAgent(rows)
}

func (db *DB) GetAgentBySaltAndProject(value1 string, value2 string) (bool, *models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Salt\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAgent(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAgentBySaltAndProject(value1 string, value2 string) ([]*models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Salt\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAgent(rows)
}

func (db *DB) GetAgentBySaltAndLabel(value1 string, value2 string) (bool, *models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Salt\" = $1 AND \"Label\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAgent(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAgentBySaltAndLabel(value1 string, value2 string) ([]*models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Salt\" = $1 AND \"Label\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAgent(rows)
}

func (db *DB) CountAgentByCreated(value time.Time) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Agent\" WHERE \"Created\" = $1;", value)
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

func (db *DB) GetAgentByCreated(value time.Time) (bool, *models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Created\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAgent(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAgentByCreated(value time.Time) ([]*models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Created\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToAgent(rows)
}

func (db *DB) GetAgentByCreatedAndUID(value1 time.Time, value2 string) (bool, *models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Created\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAgent(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAgentByCreatedAndUID(value1 time.Time, value2 string) ([]*models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Created\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAgent(rows)
}

func (db *DB) GetAgentByCreatedAndSalt(value1 time.Time, value2 string) (bool, *models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Created\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAgent(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAgentByCreatedAndSalt(value1 time.Time, value2 string) ([]*models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Created\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAgent(rows)
}

func (db *DB) GetAgentByCreatedAndProject(value1 time.Time, value2 string) (bool, *models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Created\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAgent(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAgentByCreatedAndProject(value1 time.Time, value2 string) ([]*models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Created\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAgent(rows)
}

func (db *DB) GetAgentByCreatedAndLabel(value1 time.Time, value2 string) (bool, *models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Created\" = $1 AND \"Label\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAgent(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAgentByCreatedAndLabel(value1 time.Time, value2 string) ([]*models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Created\" = $1 AND \"Label\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAgent(rows)
}

func (db *DB) CountAgentByProject(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Agent\" WHERE \"Project\" = $1;", value)
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

func (db *DB) GetAgentByProject(value string) (bool, *models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Project\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAgent(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAgentByProject(value string) ([]*models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Project\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToAgent(rows)
}

func (db *DB) GetAgentByProjectAndUID(value1 string, value2 string) (bool, *models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Project\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAgent(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAgentByProjectAndUID(value1 string, value2 string) ([]*models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Project\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAgent(rows)
}

func (db *DB) GetAgentByProjectAndSalt(value1 string, value2 string) (bool, *models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Project\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAgent(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAgentByProjectAndSalt(value1 string, value2 string) ([]*models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Project\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAgent(rows)
}

func (db *DB) GetAgentByProjectAndCreated(value1 string, value2 time.Time) (bool, *models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Project\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAgent(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAgentByProjectAndCreated(value1 string, value2 time.Time) ([]*models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Project\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAgent(rows)
}

func (db *DB) GetAgentByProjectAndLabel(value1 string, value2 string) (bool, *models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Project\" = $1 AND \"Label\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAgent(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAgentByProjectAndLabel(value1 string, value2 string) ([]*models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Project\" = $1 AND \"Label\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAgent(rows)
}

func (db *DB) CountAgentByLabel(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Agent\" WHERE \"Label\" = $1;", value)
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

func (db *DB) GetAgentByLabel(value string) (bool, *models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Label\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAgent(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAgentByLabel(value string) ([]*models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Label\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToAgent(rows)
}

func (db *DB) GetAgentByLabelAndUID(value1 string, value2 string) (bool, *models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Label\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAgent(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAgentByLabelAndUID(value1 string, value2 string) ([]*models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Label\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAgent(rows)
}

func (db *DB) GetAgentByLabelAndSalt(value1 string, value2 string) (bool, *models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Label\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAgent(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAgentByLabelAndSalt(value1 string, value2 string) ([]*models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Label\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAgent(rows)
}

func (db *DB) GetAgentByLabelAndCreated(value1 string, value2 time.Time) (bool, *models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Label\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAgent(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAgentByLabelAndCreated(value1 string, value2 time.Time) ([]*models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Label\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAgent(rows)
}

func (db *DB) GetAgentByLabelAndProject(value1 string, value2 string) (bool, *models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Label\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToAgent(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryAgentByLabelAndProject(value1 string, value2 string) ([]*models.Agent, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Label\", \"Project\", \"Salt\", \"UID\" FROM " + db.dbName + ".\"Agent\" WHERE \"Label\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToAgent(rows)
}


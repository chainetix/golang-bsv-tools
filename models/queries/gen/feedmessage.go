package database

import (
	"fmt"
	"time"
	"database/sql"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/project"
)

func (db *DB) CreateTableFeedMessage() error {
	q := "CREATE TABLE IF NOT EXISTS " + db.dbName + ".\"FeedMessage\" (\"Parents\" STRING NOT NULL, \"UID\" UUID PRIMARY KEY DEFAULT gen_random_uuid(), \"Salt\" UUID DEFAULT gen_random_uuid(), \"Created\" TIMESTAMP DEFAULT current_timestamp(), \"Project\" STRING NOT NULL, \"Target\" STRING NOT NULL, \"Type\" STRING NOT NULL, \"Subject\" STRING NOT NULL);"
	fmt.Println(q)
	_, err := db.DoExec(q)
	return err
}

func (db *DB) InsertFeedMessage(Parents string, Project string, Subject string, Target string, Type string) (*models.FeedMessage, error) {
	row := &models.FeedMessage{}
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"FeedMessage\" (\"Parents\", \"Project\", \"Subject\", \"Target\", \"Type\") VALUES ($1, $2, $3, $4, $5) RETURNING Parents, Project, Subject, Target, Type;", Parents, Project, Subject, Target, Type).Scan(&row.Created, &row.Parents, &row.Project, &row.Salt, &row.Subject, &row.Target, &row.Type, &row.UID); err != nil {
		return nil, err
	}
	return row, nil
}

func (db *DB) DeleteFeedMessage(uid string) error {
	_, err := db.DoQuery("DELETE FROM " + db.dbName + ".\"FeedMessage\" WHERE \"UID\" = $1;", uid)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) PutFeedMessage(model *models.FeedMessage) error {
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"FeedMessage\" (\"Parents\", \"Project\", \"Subject\", \"Target\", \"Type\") VALUES ($1, $2, $3, $4, $5) RETURNING \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\";", model.Parents, model.Project, model.Subject, model.Target, model.Type).Scan(&model.Created, &model.Parents, &model.Project, &model.Salt, &model.Subject, &model.Target, &model.Type, &model.UID); err != nil {
		return err
	}
	return nil
}

func (db *DB) ScanToFeedMessage(rows *sql.Rows) ([]*models.FeedMessage, error) {
	defer rows.Close()
	results := []*models.FeedMessage{}
	for rows.Next() {
		result := &models.FeedMessage{}
		err := rows.Scan(&result.Created, &result.Parents, &result.Project, &result.Salt, &result.Subject, &result.Target, &result.Type, &result.UID)
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

func (db *DB) CountFeedMessage() (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM " + db.dbName + ".\"FeedMessage\";")
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

func (db *DB) QueryFeedMessage() ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\";")
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) CountFeedMessageByParents(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"FeedMessage\" WHERE \"Parents\" = $1;", value)
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

func (db *DB) GetFeedMessageByParents(value string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Parents\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByParents(value string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Parents\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByParentsAndUID(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Parents\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByParentsAndUID(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Parents\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByParentsAndSalt(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Parents\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByParentsAndSalt(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Parents\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByParentsAndCreated(value1 string, value2 time.Time) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Parents\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByParentsAndCreated(value1 string, value2 time.Time) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Parents\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByParentsAndProject(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Parents\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByParentsAndProject(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Parents\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByParentsAndTarget(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Parents\" = $1 AND \"Target\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByParentsAndTarget(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Parents\" = $1 AND \"Target\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByParentsAndType(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Parents\" = $1 AND \"Type\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByParentsAndType(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Parents\" = $1 AND \"Type\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByParentsAndSubject(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Parents\" = $1 AND \"Subject\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByParentsAndSubject(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Parents\" = $1 AND \"Subject\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) CountFeedMessageByUID(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"FeedMessage\" WHERE \"UID\" = $1;", value)
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

func (db *DB) GetFeedMessageByUID(value string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"UID\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByUID(value string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"UID\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) CountFeedMessageBySalt(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"FeedMessage\" WHERE \"Salt\" = $1;", value)
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

func (db *DB) GetFeedMessageBySalt(value string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Salt\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageBySalt(value string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Salt\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageBySaltAndParents(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Salt\" = $1 AND \"Parents\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageBySaltAndParents(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Salt\" = $1 AND \"Parents\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageBySaltAndUID(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Salt\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageBySaltAndUID(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Salt\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageBySaltAndCreated(value1 string, value2 time.Time) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Salt\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageBySaltAndCreated(value1 string, value2 time.Time) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Salt\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageBySaltAndProject(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Salt\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageBySaltAndProject(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Salt\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageBySaltAndTarget(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Salt\" = $1 AND \"Target\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageBySaltAndTarget(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Salt\" = $1 AND \"Target\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageBySaltAndType(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Salt\" = $1 AND \"Type\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageBySaltAndType(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Salt\" = $1 AND \"Type\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageBySaltAndSubject(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Salt\" = $1 AND \"Subject\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageBySaltAndSubject(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Salt\" = $1 AND \"Subject\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) CountFeedMessageByCreated(value time.Time) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"FeedMessage\" WHERE \"Created\" = $1;", value)
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

func (db *DB) GetFeedMessageByCreated(value time.Time) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Created\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByCreated(value time.Time) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Created\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByCreatedAndParents(value1 time.Time, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Created\" = $1 AND \"Parents\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByCreatedAndParents(value1 time.Time, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Created\" = $1 AND \"Parents\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByCreatedAndUID(value1 time.Time, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Created\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByCreatedAndUID(value1 time.Time, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Created\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByCreatedAndSalt(value1 time.Time, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Created\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByCreatedAndSalt(value1 time.Time, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Created\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByCreatedAndProject(value1 time.Time, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Created\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByCreatedAndProject(value1 time.Time, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Created\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByCreatedAndTarget(value1 time.Time, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Created\" = $1 AND \"Target\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByCreatedAndTarget(value1 time.Time, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Created\" = $1 AND \"Target\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByCreatedAndType(value1 time.Time, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Created\" = $1 AND \"Type\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByCreatedAndType(value1 time.Time, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Created\" = $1 AND \"Type\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByCreatedAndSubject(value1 time.Time, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Created\" = $1 AND \"Subject\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByCreatedAndSubject(value1 time.Time, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Created\" = $1 AND \"Subject\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) CountFeedMessageByProject(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"FeedMessage\" WHERE \"Project\" = $1;", value)
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

func (db *DB) GetFeedMessageByProject(value string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Project\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByProject(value string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Project\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByProjectAndParents(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Project\" = $1 AND \"Parents\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByProjectAndParents(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Project\" = $1 AND \"Parents\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByProjectAndUID(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Project\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByProjectAndUID(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Project\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByProjectAndSalt(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Project\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByProjectAndSalt(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Project\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByProjectAndCreated(value1 string, value2 time.Time) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Project\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByProjectAndCreated(value1 string, value2 time.Time) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Project\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByProjectAndTarget(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Project\" = $1 AND \"Target\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByProjectAndTarget(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Project\" = $1 AND \"Target\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByProjectAndType(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Project\" = $1 AND \"Type\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByProjectAndType(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Project\" = $1 AND \"Type\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByProjectAndSubject(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Project\" = $1 AND \"Subject\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByProjectAndSubject(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Project\" = $1 AND \"Subject\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) CountFeedMessageByTarget(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"FeedMessage\" WHERE \"Target\" = $1;", value)
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

func (db *DB) GetFeedMessageByTarget(value string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Target\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByTarget(value string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Target\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByTargetAndParents(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Target\" = $1 AND \"Parents\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByTargetAndParents(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Target\" = $1 AND \"Parents\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByTargetAndUID(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Target\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByTargetAndUID(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Target\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByTargetAndSalt(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Target\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByTargetAndSalt(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Target\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByTargetAndCreated(value1 string, value2 time.Time) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Target\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByTargetAndCreated(value1 string, value2 time.Time) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Target\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByTargetAndProject(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Target\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByTargetAndProject(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Target\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByTargetAndType(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Target\" = $1 AND \"Type\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByTargetAndType(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Target\" = $1 AND \"Type\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByTargetAndSubject(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Target\" = $1 AND \"Subject\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByTargetAndSubject(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Target\" = $1 AND \"Subject\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) CountFeedMessageByType(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"FeedMessage\" WHERE \"Type\" = $1;", value)
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

func (db *DB) GetFeedMessageByType(value string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Type\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByType(value string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Type\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByTypeAndParents(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Type\" = $1 AND \"Parents\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByTypeAndParents(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Type\" = $1 AND \"Parents\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByTypeAndUID(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Type\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByTypeAndUID(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Type\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByTypeAndSalt(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Type\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByTypeAndSalt(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Type\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByTypeAndCreated(value1 string, value2 time.Time) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Type\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByTypeAndCreated(value1 string, value2 time.Time) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Type\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByTypeAndProject(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Type\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByTypeAndProject(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Type\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByTypeAndTarget(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Type\" = $1 AND \"Target\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByTypeAndTarget(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Type\" = $1 AND \"Target\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageByTypeAndSubject(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Type\" = $1 AND \"Subject\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageByTypeAndSubject(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Type\" = $1 AND \"Subject\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) CountFeedMessageBySubject(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"FeedMessage\" WHERE \"Subject\" = $1;", value)
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

func (db *DB) GetFeedMessageBySubject(value string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Subject\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageBySubject(value string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Subject\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageBySubjectAndParents(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Subject\" = $1 AND \"Parents\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageBySubjectAndParents(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Subject\" = $1 AND \"Parents\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageBySubjectAndUID(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Subject\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageBySubjectAndUID(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Subject\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageBySubjectAndSalt(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Subject\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageBySubjectAndSalt(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Subject\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageBySubjectAndCreated(value1 string, value2 time.Time) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Subject\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageBySubjectAndCreated(value1 string, value2 time.Time) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Subject\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageBySubjectAndProject(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Subject\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageBySubjectAndProject(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Subject\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageBySubjectAndTarget(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Subject\" = $1 AND \"Target\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageBySubjectAndTarget(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Subject\" = $1 AND \"Target\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}

func (db *DB) GetFeedMessageBySubjectAndType(value1 string, value2 string) (bool, *models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Subject\" = $1 AND \"Type\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToFeedMessage(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryFeedMessageBySubjectAndType(value1 string, value2 string) ([]*models.FeedMessage, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Parents\", \"Project\", \"Salt\", \"Subject\", \"Target\", \"Type\", \"UID\" FROM " + db.dbName + ".\"FeedMessage\" WHERE \"Subject\" = $1 AND \"Type\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToFeedMessage(rows)
}


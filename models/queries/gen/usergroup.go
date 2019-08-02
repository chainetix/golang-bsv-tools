package database

import (
	"fmt"
	"time"
	"database/sql"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/project"
)

func (db *DB) CreateTableUserGroup() error {
	q := "CREATE TABLE IF NOT EXISTS " + db.dbName + ".\"UserGroup\" (\"Salt\" UUID DEFAULT gen_random_uuid(), \"Created\" TIMESTAMP DEFAULT current_timestamp(), \"Project\" STRING NOT NULL, \"User\" STRING NOT NULL, \"Name\" STRING NOT NULL, \"UID\" UUID PRIMARY KEY DEFAULT gen_random_uuid());"
	fmt.Println(q)
	_, err := db.DoExec(q)
	return err
}

func (db *DB) InsertUserGroup(Name string, Project string, User string) (*models.UserGroup, error) {
	row := &models.UserGroup{}
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"UserGroup\" (\"Name\", \"Project\", \"User\") VALUES ($1, $2, $3) RETURNING Name, Project, User;", Name, Project, User).Scan(&row.Created, &row.Name, &row.Project, &row.Salt, &row.UID, &row.User); err != nil {
		return nil, err
	}
	return row, nil
}

func (db *DB) DeleteUserGroup(uid string) error {
	_, err := db.DoQuery("DELETE FROM " + db.dbName + ".\"UserGroup\" WHERE \"UID\" = $1;", uid)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) PutUserGroup(model *models.UserGroup) error {
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"UserGroup\" (\"Name\", \"Project\", \"User\") VALUES ($1, $2, $3) RETURNING \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\";", model.Name, model.Project, model.User).Scan(&model.Created, &model.Name, &model.Project, &model.Salt, &model.UID, &model.User); err != nil {
		return err
	}
	return nil
}

func (db *DB) ScanToUserGroup(rows *sql.Rows) ([]*models.UserGroup, error) {
	defer rows.Close()
	results := []*models.UserGroup{}
	for rows.Next() {
		result := &models.UserGroup{}
		err := rows.Scan(&result.Created, &result.Name, &result.Project, &result.Salt, &result.UID, &result.User)
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

func (db *DB) CountUserGroup() (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM " + db.dbName + ".\"UserGroup\";")
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

func (db *DB) QueryUserGroup() ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\";")
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) CountUserGroupBySalt(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"UserGroup\" WHERE \"Salt\" = $1;", value)
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

func (db *DB) GetUserGroupBySalt(value string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Salt\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupBySalt(value string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Salt\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupBySaltAndCreated(value1 string, value2 time.Time) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Salt\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupBySaltAndCreated(value1 string, value2 time.Time) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Salt\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupBySaltAndProject(value1 string, value2 string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Salt\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupBySaltAndProject(value1 string, value2 string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Salt\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupBySaltAndUser(value1 string, value2 string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Salt\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupBySaltAndUser(value1 string, value2 string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Salt\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupBySaltAndName(value1 string, value2 string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Salt\" = $1 AND \"Name\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupBySaltAndName(value1 string, value2 string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Salt\" = $1 AND \"Name\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupBySaltAndUID(value1 string, value2 string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Salt\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupBySaltAndUID(value1 string, value2 string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Salt\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) CountUserGroupByCreated(value time.Time) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"UserGroup\" WHERE \"Created\" = $1;", value)
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

func (db *DB) GetUserGroupByCreated(value time.Time) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Created\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByCreated(value time.Time) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Created\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupByCreatedAndSalt(value1 time.Time, value2 string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Created\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByCreatedAndSalt(value1 time.Time, value2 string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Created\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupByCreatedAndProject(value1 time.Time, value2 string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Created\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByCreatedAndProject(value1 time.Time, value2 string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Created\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupByCreatedAndUser(value1 time.Time, value2 string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Created\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByCreatedAndUser(value1 time.Time, value2 string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Created\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupByCreatedAndName(value1 time.Time, value2 string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Created\" = $1 AND \"Name\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByCreatedAndName(value1 time.Time, value2 string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Created\" = $1 AND \"Name\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupByCreatedAndUID(value1 time.Time, value2 string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Created\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByCreatedAndUID(value1 time.Time, value2 string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Created\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) CountUserGroupByProject(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"UserGroup\" WHERE \"Project\" = $1;", value)
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

func (db *DB) GetUserGroupByProject(value string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Project\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByProject(value string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Project\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupByProjectAndSalt(value1 string, value2 string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Project\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByProjectAndSalt(value1 string, value2 string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Project\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupByProjectAndCreated(value1 string, value2 time.Time) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Project\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByProjectAndCreated(value1 string, value2 time.Time) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Project\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupByProjectAndUser(value1 string, value2 string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Project\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByProjectAndUser(value1 string, value2 string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Project\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupByProjectAndName(value1 string, value2 string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Project\" = $1 AND \"Name\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByProjectAndName(value1 string, value2 string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Project\" = $1 AND \"Name\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupByProjectAndUID(value1 string, value2 string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Project\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByProjectAndUID(value1 string, value2 string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Project\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) CountUserGroupByUser(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"UserGroup\" WHERE \"User\" = $1;", value)
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

func (db *DB) GetUserGroupByUser(value string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"User\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByUser(value string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"User\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupByUserAndSalt(value1 string, value2 string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"User\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByUserAndSalt(value1 string, value2 string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"User\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupByUserAndCreated(value1 string, value2 time.Time) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"User\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByUserAndCreated(value1 string, value2 time.Time) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"User\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupByUserAndProject(value1 string, value2 string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"User\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByUserAndProject(value1 string, value2 string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"User\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupByUserAndName(value1 string, value2 string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"User\" = $1 AND \"Name\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByUserAndName(value1 string, value2 string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"User\" = $1 AND \"Name\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupByUserAndUID(value1 string, value2 string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"User\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByUserAndUID(value1 string, value2 string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"User\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) CountUserGroupByName(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"UserGroup\" WHERE \"Name\" = $1;", value)
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

func (db *DB) GetUserGroupByName(value string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Name\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByName(value string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Name\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupByNameAndSalt(value1 string, value2 string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Name\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByNameAndSalt(value1 string, value2 string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Name\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupByNameAndCreated(value1 string, value2 time.Time) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Name\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByNameAndCreated(value1 string, value2 time.Time) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Name\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupByNameAndProject(value1 string, value2 string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Name\" = $1 AND \"Project\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByNameAndProject(value1 string, value2 string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Name\" = $1 AND \"Project\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupByNameAndUser(value1 string, value2 string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Name\" = $1 AND \"User\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByNameAndUser(value1 string, value2 string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Name\" = $1 AND \"User\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) GetUserGroupByNameAndUID(value1 string, value2 string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Name\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByNameAndUID(value1 string, value2 string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"Name\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}

func (db *DB) CountUserGroupByUID(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"UserGroup\" WHERE \"UID\" = $1;", value)
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

func (db *DB) GetUserGroupByUID(value string) (bool, *models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"UID\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToUserGroup(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryUserGroupByUID(value string) ([]*models.UserGroup, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"Name\", \"Project\", \"Salt\", \"UID\", \"User\" FROM " + db.dbName + ".\"UserGroup\" WHERE \"UID\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToUserGroup(rows)
}


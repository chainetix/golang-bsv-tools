package database

import (
	"fmt"
	"time"
	"database/sql"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/project"
)

func (db *DB) CreateTableProject() error {
	q := "CREATE TABLE IF NOT EXISTS " + db.dbName + ".\"Project\" (\"Created\" TIMESTAMP DEFAULT current_timestamp(), \"BurnAddress\" STRING NOT NULL, \"UID\" UUID PRIMARY KEY DEFAULT gen_random_uuid(), \"Salt\" UUID DEFAULT gen_random_uuid(), \"StartBlock\" INT NOT NULL, \"Description\" STRING NOT NULL, \"DefaultStream\" STRING NOT NULL, \"Public\" BOOL NOT NULL, \"Title\" STRING NOT NULL);"
	fmt.Println(q)
	_, err := db.DoExec(q)
	return err
}

func (db *DB) InsertProject(BurnAddress string, DefaultStream string, Description string, Public bool, StartBlock int, Title string) (*models.Project, error) {
	row := &models.Project{}
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Project\" (\"BurnAddress\", \"DefaultStream\", \"Description\", \"Public\", \"StartBlock\", \"Title\") VALUES ($1, $2, $3, $4, $5, $6) RETURNING BurnAddress, DefaultStream, Description, Public, StartBlock, Title;", BurnAddress, DefaultStream, Description, Public, StartBlock, Title).Scan(&row.BurnAddress, &row.Created, &row.DefaultStream, &row.Description, &row.Public, &row.Salt, &row.StartBlock, &row.Title, &row.UID); err != nil {
		return nil, err
	}
	return row, nil
}

func (db *DB) DeleteProject(uid string) error {
	_, err := db.DoQuery("DELETE FROM " + db.dbName + ".\"Project\" WHERE \"UID\" = $1;", uid)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) PutProject(model *models.Project) error {
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Project\" (\"BurnAddress\", \"DefaultStream\", \"Description\", \"Public\", \"StartBlock\", \"Title\") VALUES ($1, $2, $3, $4, $5, $6) RETURNING \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\";", model.BurnAddress, model.DefaultStream, model.Description, model.Public, model.StartBlock, model.Title).Scan(&model.BurnAddress, &model.Created, &model.DefaultStream, &model.Description, &model.Public, &model.Salt, &model.StartBlock, &model.Title, &model.UID); err != nil {
		return err
	}
	return nil
}

func (db *DB) ScanToProject(rows *sql.Rows) ([]*models.Project, error) {
	defer rows.Close()
	results := []*models.Project{}
	for rows.Next() {
		result := &models.Project{}
		err := rows.Scan(&result.BurnAddress, &result.Created, &result.DefaultStream, &result.Description, &result.Public, &result.Salt, &result.StartBlock, &result.Title, &result.UID)
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

func (db *DB) CountProject() (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM " + db.dbName + ".\"Project\";")
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

func (db *DB) QueryProject() ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\";")
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) CountProjectByCreated(value time.Time) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Project\" WHERE \"Created\" = $1;", value)
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

func (db *DB) GetProjectByCreated(value time.Time) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Created\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByCreated(value time.Time) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Created\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByCreatedAndBurnAddress(value1 time.Time, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Created\" = $1 AND \"BurnAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByCreatedAndBurnAddress(value1 time.Time, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Created\" = $1 AND \"BurnAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByCreatedAndUID(value1 time.Time, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Created\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByCreatedAndUID(value1 time.Time, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Created\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByCreatedAndSalt(value1 time.Time, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Created\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByCreatedAndSalt(value1 time.Time, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Created\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByCreatedAndStartBlock(value1 time.Time, value2 int) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Created\" = $1 AND \"StartBlock\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByCreatedAndStartBlock(value1 time.Time, value2 int) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Created\" = $1 AND \"StartBlock\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByCreatedAndDescription(value1 time.Time, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Created\" = $1 AND \"Description\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByCreatedAndDescription(value1 time.Time, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Created\" = $1 AND \"Description\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByCreatedAndDefaultStream(value1 time.Time, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Created\" = $1 AND \"DefaultStream\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByCreatedAndDefaultStream(value1 time.Time, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Created\" = $1 AND \"DefaultStream\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByCreatedAndPublic(value1 time.Time, value2 bool) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Created\" = $1 AND \"Public\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByCreatedAndPublic(value1 time.Time, value2 bool) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Created\" = $1 AND \"Public\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByCreatedAndTitle(value1 time.Time, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Created\" = $1 AND \"Title\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByCreatedAndTitle(value1 time.Time, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Created\" = $1 AND \"Title\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) CountProjectByBurnAddress(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Project\" WHERE \"BurnAddress\" = $1;", value)
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

func (db *DB) GetProjectByBurnAddress(value string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"BurnAddress\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByBurnAddress(value string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"BurnAddress\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByBurnAddressAndCreated(value1 string, value2 time.Time) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"BurnAddress\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByBurnAddressAndCreated(value1 string, value2 time.Time) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"BurnAddress\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByBurnAddressAndUID(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"BurnAddress\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByBurnAddressAndUID(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"BurnAddress\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByBurnAddressAndSalt(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"BurnAddress\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByBurnAddressAndSalt(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"BurnAddress\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByBurnAddressAndStartBlock(value1 string, value2 int) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"BurnAddress\" = $1 AND \"StartBlock\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByBurnAddressAndStartBlock(value1 string, value2 int) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"BurnAddress\" = $1 AND \"StartBlock\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByBurnAddressAndDescription(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"BurnAddress\" = $1 AND \"Description\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByBurnAddressAndDescription(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"BurnAddress\" = $1 AND \"Description\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByBurnAddressAndDefaultStream(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"BurnAddress\" = $1 AND \"DefaultStream\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByBurnAddressAndDefaultStream(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"BurnAddress\" = $1 AND \"DefaultStream\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByBurnAddressAndPublic(value1 string, value2 bool) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"BurnAddress\" = $1 AND \"Public\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByBurnAddressAndPublic(value1 string, value2 bool) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"BurnAddress\" = $1 AND \"Public\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByBurnAddressAndTitle(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"BurnAddress\" = $1 AND \"Title\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByBurnAddressAndTitle(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"BurnAddress\" = $1 AND \"Title\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) CountProjectByUID(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Project\" WHERE \"UID\" = $1;", value)
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

func (db *DB) GetProjectByUID(value string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"UID\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByUID(value string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"UID\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) CountProjectBySalt(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Project\" WHERE \"Salt\" = $1;", value)
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

func (db *DB) GetProjectBySalt(value string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Salt\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectBySalt(value string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Salt\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectBySaltAndCreated(value1 string, value2 time.Time) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Salt\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectBySaltAndCreated(value1 string, value2 time.Time) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Salt\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectBySaltAndBurnAddress(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Salt\" = $1 AND \"BurnAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectBySaltAndBurnAddress(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Salt\" = $1 AND \"BurnAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectBySaltAndUID(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Salt\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectBySaltAndUID(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Salt\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectBySaltAndStartBlock(value1 string, value2 int) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Salt\" = $1 AND \"StartBlock\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectBySaltAndStartBlock(value1 string, value2 int) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Salt\" = $1 AND \"StartBlock\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectBySaltAndDescription(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Salt\" = $1 AND \"Description\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectBySaltAndDescription(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Salt\" = $1 AND \"Description\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectBySaltAndDefaultStream(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Salt\" = $1 AND \"DefaultStream\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectBySaltAndDefaultStream(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Salt\" = $1 AND \"DefaultStream\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectBySaltAndPublic(value1 string, value2 bool) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Salt\" = $1 AND \"Public\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectBySaltAndPublic(value1 string, value2 bool) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Salt\" = $1 AND \"Public\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectBySaltAndTitle(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Salt\" = $1 AND \"Title\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectBySaltAndTitle(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Salt\" = $1 AND \"Title\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) CountProjectByStartBlock(value int) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Project\" WHERE \"StartBlock\" = $1;", value)
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

func (db *DB) GetProjectByStartBlock(value int) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"StartBlock\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByStartBlock(value int) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"StartBlock\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByStartBlockAndCreated(value1 int, value2 time.Time) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"StartBlock\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByStartBlockAndCreated(value1 int, value2 time.Time) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"StartBlock\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByStartBlockAndBurnAddress(value1 int, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"StartBlock\" = $1 AND \"BurnAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByStartBlockAndBurnAddress(value1 int, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"StartBlock\" = $1 AND \"BurnAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByStartBlockAndUID(value1 int, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"StartBlock\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByStartBlockAndUID(value1 int, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"StartBlock\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByStartBlockAndSalt(value1 int, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"StartBlock\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByStartBlockAndSalt(value1 int, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"StartBlock\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByStartBlockAndDescription(value1 int, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"StartBlock\" = $1 AND \"Description\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByStartBlockAndDescription(value1 int, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"StartBlock\" = $1 AND \"Description\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByStartBlockAndDefaultStream(value1 int, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"StartBlock\" = $1 AND \"DefaultStream\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByStartBlockAndDefaultStream(value1 int, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"StartBlock\" = $1 AND \"DefaultStream\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByStartBlockAndPublic(value1 int, value2 bool) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"StartBlock\" = $1 AND \"Public\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByStartBlockAndPublic(value1 int, value2 bool) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"StartBlock\" = $1 AND \"Public\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByStartBlockAndTitle(value1 int, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"StartBlock\" = $1 AND \"Title\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByStartBlockAndTitle(value1 int, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"StartBlock\" = $1 AND \"Title\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) CountProjectByDescription(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Project\" WHERE \"Description\" = $1;", value)
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

func (db *DB) GetProjectByDescription(value string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Description\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByDescription(value string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Description\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByDescriptionAndCreated(value1 string, value2 time.Time) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Description\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByDescriptionAndCreated(value1 string, value2 time.Time) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Description\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByDescriptionAndBurnAddress(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Description\" = $1 AND \"BurnAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByDescriptionAndBurnAddress(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Description\" = $1 AND \"BurnAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByDescriptionAndUID(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Description\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByDescriptionAndUID(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Description\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByDescriptionAndSalt(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Description\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByDescriptionAndSalt(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Description\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByDescriptionAndStartBlock(value1 string, value2 int) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Description\" = $1 AND \"StartBlock\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByDescriptionAndStartBlock(value1 string, value2 int) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Description\" = $1 AND \"StartBlock\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByDescriptionAndDefaultStream(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Description\" = $1 AND \"DefaultStream\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByDescriptionAndDefaultStream(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Description\" = $1 AND \"DefaultStream\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByDescriptionAndPublic(value1 string, value2 bool) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Description\" = $1 AND \"Public\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByDescriptionAndPublic(value1 string, value2 bool) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Description\" = $1 AND \"Public\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByDescriptionAndTitle(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Description\" = $1 AND \"Title\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByDescriptionAndTitle(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Description\" = $1 AND \"Title\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) CountProjectByDefaultStream(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Project\" WHERE \"DefaultStream\" = $1;", value)
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

func (db *DB) GetProjectByDefaultStream(value string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"DefaultStream\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByDefaultStream(value string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"DefaultStream\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByDefaultStreamAndCreated(value1 string, value2 time.Time) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"DefaultStream\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByDefaultStreamAndCreated(value1 string, value2 time.Time) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"DefaultStream\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByDefaultStreamAndBurnAddress(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"DefaultStream\" = $1 AND \"BurnAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByDefaultStreamAndBurnAddress(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"DefaultStream\" = $1 AND \"BurnAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByDefaultStreamAndUID(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"DefaultStream\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByDefaultStreamAndUID(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"DefaultStream\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByDefaultStreamAndSalt(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"DefaultStream\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByDefaultStreamAndSalt(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"DefaultStream\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByDefaultStreamAndStartBlock(value1 string, value2 int) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"DefaultStream\" = $1 AND \"StartBlock\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByDefaultStreamAndStartBlock(value1 string, value2 int) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"DefaultStream\" = $1 AND \"StartBlock\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByDefaultStreamAndDescription(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"DefaultStream\" = $1 AND \"Description\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByDefaultStreamAndDescription(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"DefaultStream\" = $1 AND \"Description\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByDefaultStreamAndPublic(value1 string, value2 bool) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"DefaultStream\" = $1 AND \"Public\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByDefaultStreamAndPublic(value1 string, value2 bool) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"DefaultStream\" = $1 AND \"Public\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByDefaultStreamAndTitle(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"DefaultStream\" = $1 AND \"Title\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByDefaultStreamAndTitle(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"DefaultStream\" = $1 AND \"Title\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) CountProjectByPublic(value bool) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Project\" WHERE \"Public\" = $1;", value)
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

func (db *DB) GetProjectByPublic(value bool) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Public\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByPublic(value bool) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Public\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByPublicAndCreated(value1 bool, value2 time.Time) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Public\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByPublicAndCreated(value1 bool, value2 time.Time) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Public\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByPublicAndBurnAddress(value1 bool, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Public\" = $1 AND \"BurnAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByPublicAndBurnAddress(value1 bool, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Public\" = $1 AND \"BurnAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByPublicAndUID(value1 bool, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Public\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByPublicAndUID(value1 bool, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Public\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByPublicAndSalt(value1 bool, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Public\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByPublicAndSalt(value1 bool, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Public\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByPublicAndStartBlock(value1 bool, value2 int) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Public\" = $1 AND \"StartBlock\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByPublicAndStartBlock(value1 bool, value2 int) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Public\" = $1 AND \"StartBlock\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByPublicAndDescription(value1 bool, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Public\" = $1 AND \"Description\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByPublicAndDescription(value1 bool, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Public\" = $1 AND \"Description\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByPublicAndDefaultStream(value1 bool, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Public\" = $1 AND \"DefaultStream\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByPublicAndDefaultStream(value1 bool, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Public\" = $1 AND \"DefaultStream\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByPublicAndTitle(value1 bool, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Public\" = $1 AND \"Title\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByPublicAndTitle(value1 bool, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Public\" = $1 AND \"Title\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) CountProjectByTitle(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Project\" WHERE \"Title\" = $1;", value)
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

func (db *DB) GetProjectByTitle(value string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Title\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByTitle(value string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Title\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByTitleAndCreated(value1 string, value2 time.Time) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Title\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByTitleAndCreated(value1 string, value2 time.Time) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Title\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByTitleAndBurnAddress(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Title\" = $1 AND \"BurnAddress\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByTitleAndBurnAddress(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Title\" = $1 AND \"BurnAddress\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByTitleAndUID(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Title\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByTitleAndUID(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Title\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByTitleAndSalt(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Title\" = $1 AND \"Salt\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByTitleAndSalt(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Title\" = $1 AND \"Salt\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByTitleAndStartBlock(value1 string, value2 int) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Title\" = $1 AND \"StartBlock\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByTitleAndStartBlock(value1 string, value2 int) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Title\" = $1 AND \"StartBlock\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByTitleAndDescription(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Title\" = $1 AND \"Description\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByTitleAndDescription(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Title\" = $1 AND \"Description\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByTitleAndDefaultStream(value1 string, value2 string) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Title\" = $1 AND \"DefaultStream\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByTitleAndDefaultStream(value1 string, value2 string) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Title\" = $1 AND \"DefaultStream\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}

func (db *DB) GetProjectByTitleAndPublic(value1 string, value2 bool) (bool, *models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Title\" = $1 AND \"Public\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToProject(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryProjectByTitleAndPublic(value1 string, value2 bool) ([]*models.Project, error) {
	rows, err := db.DoQuery("SELECT \"BurnAddress\", \"Created\", \"DefaultStream\", \"Description\", \"Public\", \"Salt\", \"StartBlock\", \"Title\", \"UID\" FROM " + db.dbName + ".\"Project\" WHERE \"Title\" = $1 AND \"Public\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToProject(rows)
}


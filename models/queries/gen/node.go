package database

import (
	"fmt"
	"database/sql"
	//
	"gitlab.com/chainetix-groups/chains/multichain/omega/codebase/libraries/models/network"
)

func (db *DB) CreateTableNode() error {
	q := "CREATE TABLE IF NOT EXISTS " + db.dbName + ".\"Node\" (\"UID\" UUID PRIMARY KEY DEFAULT gen_random_uuid(), \"Created\" TIMESTAMP DEFAULT current_timestamp(), \"IPv4\" STRING NOT NULL, \"RPCUser\" STRING NOT NULL, \"RPCPassword\" STRING NOT NULL, \"Model\" INT NOT NULL);"
	fmt.Println(q)
	_, err := db.DoExec(q)
	return err
}

func (db *DB) InsertNode(IPv4 string, Model int, RPCPassword string, RPCUser string) (*models.Node, error) {
	row := &models.Node{}
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Node\" (\"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\") VALUES ($1, $2, $3, $4) RETURNING IPv4, Model, RPCPassword, RPCUser;", IPv4, Model, RPCPassword, RPCUser).Scan(&row.Created, &row.IPv4, &row.Model, &row.RPCPassword, &row.RPCUser, &row.UID); err != nil {
		return nil, err
	}
	return row, nil
}

func (db *DB) DeleteNode(uid string) error {
	_, err := db.DoQuery("DELETE FROM " + db.dbName + ".\"Node\" WHERE \"UID\" = $1;", uid)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) PutNode(model *models.Node) error {
	if err := db.DoQueryRow("INSERT INTO " + db.dbName + ".\"Node\" (\"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\") VALUES ($1, $2, $3, $4) RETURNING \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\";", model.IPv4, model.Model, model.RPCPassword, model.RPCUser).Scan(&model.Created, &model.IPv4, &model.Model, &model.RPCPassword, &model.RPCUser, &model.UID); err != nil {
		return err
	}
	return nil
}

func (db *DB) ScanToNode(rows *sql.Rows) ([]*models.Node, error) {
	defer rows.Close()
	results := []*models.Node{}
	for rows.Next() {
		result := &models.Node{}
		err := rows.Scan(&result.Created, &result.IPv4, &result.Model, &result.RPCPassword, &result.RPCUser, &result.UID)
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

func (db *DB) CountNode() (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM " + db.dbName + ".\"Node\";")
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

func (db *DB) QueryNode() ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\";")
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) CountNodeByUID(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Node\" WHERE \"UID\" = $1;", value)
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

func (db *DB) GetNodeByUID(value string) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"UID\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByUID(value string) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"UID\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) CountNodeByCreated(value int64) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Node\" WHERE \"Created\" = $1;", value)
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

func (db *DB) GetNodeByCreated(value int64) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"Created\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByCreated(value int64) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"Created\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByCreatedAndUID(value1 int64, value2 string) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"Created\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByCreatedAndUID(value1 int64, value2 string) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"Created\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByCreatedAndIPv4(value1 int64, value2 string) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"Created\" = $1 AND \"IPv4\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByCreatedAndIPv4(value1 int64, value2 string) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"Created\" = $1 AND \"IPv4\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByCreatedAndRPCUser(value1 int64, value2 string) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"Created\" = $1 AND \"RPCUser\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByCreatedAndRPCUser(value1 int64, value2 string) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"Created\" = $1 AND \"RPCUser\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByCreatedAndRPCPassword(value1 int64, value2 string) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"Created\" = $1 AND \"RPCPassword\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByCreatedAndRPCPassword(value1 int64, value2 string) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"Created\" = $1 AND \"RPCPassword\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByCreatedAndModel(value1 int64, value2 int) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"Created\" = $1 AND \"Model\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByCreatedAndModel(value1 int64, value2 int) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"Created\" = $1 AND \"Model\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) CountNodeByIPv4(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Node\" WHERE \"IPv4\" = $1;", value)
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

func (db *DB) GetNodeByIPv4(value string) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"IPv4\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByIPv4(value string) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"IPv4\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByIPv4AndUID(value1 string, value2 string) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"IPv4\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByIPv4AndUID(value1 string, value2 string) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"IPv4\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByIPv4AndCreated(value1 string, value2 int64) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"IPv4\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByIPv4AndCreated(value1 string, value2 int64) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"IPv4\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByIPv4AndRPCUser(value1 string, value2 string) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"IPv4\" = $1 AND \"RPCUser\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByIPv4AndRPCUser(value1 string, value2 string) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"IPv4\" = $1 AND \"RPCUser\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByIPv4AndRPCPassword(value1 string, value2 string) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"IPv4\" = $1 AND \"RPCPassword\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByIPv4AndRPCPassword(value1 string, value2 string) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"IPv4\" = $1 AND \"RPCPassword\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByIPv4AndModel(value1 string, value2 int) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"IPv4\" = $1 AND \"Model\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByIPv4AndModel(value1 string, value2 int) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"IPv4\" = $1 AND \"Model\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) CountNodeByRPCUser(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Node\" WHERE \"RPCUser\" = $1;", value)
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

func (db *DB) GetNodeByRPCUser(value string) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"RPCUser\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByRPCUser(value string) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"RPCUser\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByRPCUserAndUID(value1 string, value2 string) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"RPCUser\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByRPCUserAndUID(value1 string, value2 string) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"RPCUser\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByRPCUserAndCreated(value1 string, value2 int64) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"RPCUser\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByRPCUserAndCreated(value1 string, value2 int64) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"RPCUser\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByRPCUserAndIPv4(value1 string, value2 string) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"RPCUser\" = $1 AND \"IPv4\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByRPCUserAndIPv4(value1 string, value2 string) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"RPCUser\" = $1 AND \"IPv4\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByRPCUserAndRPCPassword(value1 string, value2 string) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"RPCUser\" = $1 AND \"RPCPassword\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByRPCUserAndRPCPassword(value1 string, value2 string) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"RPCUser\" = $1 AND \"RPCPassword\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByRPCUserAndModel(value1 string, value2 int) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"RPCUser\" = $1 AND \"Model\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByRPCUserAndModel(value1 string, value2 int) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"RPCUser\" = $1 AND \"Model\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) CountNodeByRPCPassword(value string) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Node\" WHERE \"RPCPassword\" = $1;", value)
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

func (db *DB) GetNodeByRPCPassword(value string) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"RPCPassword\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByRPCPassword(value string) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"RPCPassword\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByRPCPasswordAndUID(value1 string, value2 string) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"RPCPassword\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByRPCPasswordAndUID(value1 string, value2 string) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"RPCPassword\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByRPCPasswordAndCreated(value1 string, value2 int64) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"RPCPassword\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByRPCPasswordAndCreated(value1 string, value2 int64) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"RPCPassword\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByRPCPasswordAndIPv4(value1 string, value2 string) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"RPCPassword\" = $1 AND \"IPv4\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByRPCPasswordAndIPv4(value1 string, value2 string) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"RPCPassword\" = $1 AND \"IPv4\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByRPCPasswordAndRPCUser(value1 string, value2 string) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"RPCPassword\" = $1 AND \"RPCUser\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByRPCPasswordAndRPCUser(value1 string, value2 string) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"RPCPassword\" = $1 AND \"RPCUser\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByRPCPasswordAndModel(value1 string, value2 int) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"RPCPassword\" = $1 AND \"Model\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByRPCPasswordAndModel(value1 string, value2 int) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"RPCPassword\" = $1 AND \"Model\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) CountNodeByModel(value int) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM \"Node\" WHERE \"Model\" = $1;", value)
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

func (db *DB) GetNodeByModel(value int) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"Model\" = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByModel(value int) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"Model\" = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByModelAndUID(value1 int, value2 string) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"Model\" = $1 AND \"UID\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByModelAndUID(value1 int, value2 string) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"Model\" = $1 AND \"UID\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByModelAndCreated(value1 int, value2 int64) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"Model\" = $1 AND \"Created\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByModelAndCreated(value1 int, value2 int64) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"Model\" = $1 AND \"Created\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByModelAndIPv4(value1 int, value2 string) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"Model\" = $1 AND \"IPv4\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByModelAndIPv4(value1 int, value2 string) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"Model\" = $1 AND \"IPv4\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByModelAndRPCUser(value1 int, value2 string) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"Model\" = $1 AND \"RPCUser\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByModelAndRPCUser(value1 int, value2 string) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"Model\" = $1 AND \"RPCUser\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}

func (db *DB) GetNodeByModelAndRPCPassword(value1 int, value2 string) (bool, *models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"Model\" = $1 AND \"RPCPassword\" = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanToNode(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

func (db *DB) QueryNodeByModelAndRPCPassword(value1 int, value2 string) ([]*models.Node, error) {
	rows, err := db.DoQuery("SELECT \"Created\", \"IPv4\", \"Model\", \"RPCPassword\", \"RPCUser\", \"UID\" FROM " + db.dbName + ".\"Node\" WHERE \"Model\" = $1 AND \"RPCPassword\" = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanToNode(rows)
}


package database

const (
CONST_PACKAGE_IMPORTS_PROJECT = `import (
	"fmt"
	"time"
	"database/sql"
	//
	"%s"
)

`
CONST_PACKAGE_IMPORTS_NETWORK = `import (
	"fmt"
	"database/sql"
	//
	"%s"
)

`
CONST_PACKAGE_DECLARATION = `package database

`
CONST_TEMPLATE_CREATETABLE = `
	if err := db.CreateTable%s(); err != nil {
		return err
	}

`
CONST_TEMPLATE_SQL_CREATETABLE = `func (db *DB) CreateTable%s() error {
	q := "CREATE TABLE IF NOT EXISTS %s (%s);"
	fmt.Println(q)
	_, err := db.DoExec(q)
	return err
}

`
CONST_TEMPLATE_SQL_DELETE = `func (db *DB) Delete%s(uid string) error {
	_, err := db.DoQuery("DELETE FROM %s WHERE %s = $1;", uid)
	if err != nil {
		return err
	}
	return nil
}

`
CONST_TEMPLATE_SQL_PUT = `func (db *DB) Put%s(%s) error {
	if err := db.DoQueryRow("INSERT INTO %s (%s) VALUES (%s) RETURNING %s;", %s).Scan(%s); err != nil {
		return err
	}
	return nil
}

`
CONST_TEMPLATE_SQL_INSERT = `func (db *DB) Insert%s(%s) (*models.%s, error) {
	row := &models.%s{}
	if err := db.DoQueryRow("INSERT INTO %s (%s) VALUES (%s) RETURNING %s;", %s).Scan(%s); err != nil {
		return nil, err
	}
	return row, nil
}

`
CONST_TEMPLATE_SQL_GET = `func (db *DB) Get%sBy%s(value %s) (bool, *models.%s, error) {
	rows, err := db.DoQuery("SELECT %s FROM %s WHERE %s = $1 LIMIT 1;", value)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanTo%s(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

`
CONST_TEMPLATE_SQL_GETAND = `func (db *DB) Get%sBy%sAnd%s(value1 %s, value2 %s) (bool, *models.%s, error) {
	rows, err := db.DoQuery("SELECT %s FROM %s WHERE %s = $1 AND %s = $2 LIMIT 1;", value1, value2)
	if err != nil {
		return false, nil, err
	}
	results, err := db.ScanTo%s(rows)
	if err != nil {
		return false, nil, err
	}
	if len(results) == 0 {
		return false, nil, nil
	}
	return true, results[0], nil
}

`
CONST_TEMPLATE_SQL_COUNT = `func (db *DB) Count%sBy%s(value %s) (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM %s WHERE %s = $1;", value)
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

`
CONST_TEMPLATE_SQL_COUNTALL = `func (db *DB) Count%s() (int, error) {
	rows, err := db.DoQuery("SELECT count(*) FROM %s;")
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

`
CONST_TEMPLATE_SQL_QUERY = `func (db *DB) Query%sBy%s(value %s) ([]*models.%s, error) {
	rows, err := db.DoQuery("SELECT %s FROM %s WHERE %s = $1;", value)
	if err != nil {
		return nil, err
	}
	return db.ScanTo%s(rows)
}

`
CONST_TEMPLATE_SQL_QUERYAND = `func (db *DB) Query%sBy%sAnd%s(value1 %s, value2 %s) ([]*models.%s, error) {
	rows, err := db.DoQuery("SELECT %s FROM %s WHERE %s = $1 AND %s = $2;", value1, value2)
	if err != nil {
		return nil, err
	}
	return db.ScanTo%s(rows)
}

`
CONST_TEMPLATE_SQL_ALL = `func (db *DB) Query%s() ([]*models.%s, error) {
	rows, err := db.DoQuery("SELECT %s FROM %s;")
	if err != nil {
		return nil, err
	}
	return db.ScanTo%s(rows)
}

`
CONST_TEMPLATE_SQL_ROWSCAN = `func (db *DB) ScanTo%s(rows *sql.Rows) ([]*models.%s, error) {
	defer rows.Close()
	results := []*models.%s{}
	for rows.Next() {
		result := &models.%s{}
		err := rows.Scan(%s)
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

`
)

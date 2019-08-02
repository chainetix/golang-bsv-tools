package database

import (
	"fmt"
	"time"
	"database/sql"
	//
//	"golang.org/x/net/context"
	//
	_ "github.com/lib/pq"
)

const (
	CONST_PORT = 80
	CONST_DBURL = "postgresql://root@cockroachdb:26257?sslmode=disable"
)

type DB struct {
	dbName string
	Debug bool
}

func NewClient(dbName string) *DB {
	return &DB{
		dbName: dbName,
	}
}

func refreshClient() (sqlclient *sql.DB) {

	var err error
	for {
		sqlclient, err = sql.Open("postgres", CONST_DBURL)
		if err == nil {
			break
		}
		fmt.Println("Waiting for database to be ready:", err)
		time.Sleep(time.Second / 3)
	}

	return
}

func (db *DB) DoQuery(s string, args ...interface{}) (*sql.Rows, error) {

	c := refreshClient()
	defer c.Close()

	if db.Debug {
		fmt.Println("QUERY", s, args)
	}
	return c.Query(s, args...)
}

func (db *DB) DoQueryRow(s string, args ...interface{}) *sql.Row {

	c := refreshClient()
	defer c.Close()

	if db.Debug {
		fmt.Println("QUERYROW", s, args)
	}
	return c.QueryRow(s, args...)
}

func (db *DB) DoExec(s string, args ...interface{}) (sql.Result, error) {

	c := refreshClient()
	defer c.Close()

	if db.Debug {
		fmt.Println("EXEC", s, args)
	}
	return c.Exec(s, args...)
}

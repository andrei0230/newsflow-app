package storage

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func StartMySql() (*sql.DB, error) {
	var err error
	database, err := sql.Open("mysql", "root:1130@tcp(127.0.0.1:3306)/newsflow")
	if err != nil {
		return nil, err
	}
	return database, nil
}

func StopSql(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}

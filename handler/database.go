package handler

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB(dataSource string) error {
	var err error
	db, err = sql.Open("mysql", dataSource)
	if err != nil {
		return err
	}
	return nil
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func GetDB() *sql.DB {
	return db
}

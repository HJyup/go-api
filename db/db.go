package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

func NewMySQLStorage(config mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

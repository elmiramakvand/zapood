package config

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

func GetDB() (db *sql.DB, err error) {
	db, err = sql.Open("mssql", "server=185.94.98.228;user id=rose;password=z6prT9@5;database=iranvar;")
	return
}

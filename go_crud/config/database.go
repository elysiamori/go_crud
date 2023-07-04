package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//Konfigurasi Database

func DBConnection() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "go_crud"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return db, err
}

package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() *sql.DB {
	dsn := "root:rootpassword@tcp(localhost:3307)/crud_employee"

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err.Error())
	}

	// Connect ke database
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db
}

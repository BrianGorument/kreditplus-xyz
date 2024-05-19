package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(connectionString string) *sql.DB {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Could not connect to database: %s\n", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Could not ping database: %s\n", err)
	}
	return db
}

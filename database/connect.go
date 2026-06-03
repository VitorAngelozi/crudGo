package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() {
	connectStr := "host=localhost port=5432 user=postgres password=admin dbname=crudgo sslmode=disable"

	db, err := sql.Open("postgres", connectStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	log.Println("Database connection established")
}

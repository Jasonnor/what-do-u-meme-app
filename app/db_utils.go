package app

import (
	"database/sql"
	"log"
)

func connectDB(driver string, login string) (*sql.DB, error) {
	var db *sql.DB
	db, err := sql.Open(
		driver,
		login,
	)
	if err != nil {
		log.Fatal(err)
		return db, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return db, err
	}

	return db, nil
}

func mockConnectDB(driver string, login string) (*sql.DB, error) {
	var db *sql.DB
	return db, nil
}

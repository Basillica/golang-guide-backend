package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func FetchUser(email string) {
	db, err := sql.Open("mysql", "user:password@/database")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECE * FROM users WHERE email = ?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(email)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// marshalling to a user struct
}

func Form() {
	email := "test.user@gmail.com; DROP TABLE users"
	FetchUser(email)
}

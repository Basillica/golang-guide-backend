package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func GetClient() *sql.DB {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s",
		"", "", "", "", "",
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if dErr := db.Ping(); dErr != nil {
		panic(dErr)
	}

	// pooling configurations
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(50)
	db.SetConnMaxLifetime(time.Hour)
	return db
}

package mock_test

import (
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, sqlmock.Sqlmock, *sql.DB) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("[sqlmock new] %s", err)
	}

	dialectgor := mysql.New(mysql.Config{
		Conn:       sqlDB,
		DriverName: "mysql",
	})

	//open database conn
	db, err := gorm.Open(dialectgor, &gorm.Config{})
	if err != nil {
		log.Fatalf("[sqlmock new] %s", err)
	}

	return db, mock, sqlDB
}

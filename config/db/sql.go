package db

import (
	"fmt"
	"time"

	"github.com/Basillica/golang-guide/config/helpers"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func GetClientGorm() *gorm.DB {
	var db *gorm.DB
	var err error
	server := helpers.GetEnv("SQL_SERVER", "")
	user := helpers.GetEnv("SQL_USER", "")
	password := helpers.GetEnv("SQL_PASSWORD", "")
	port := helpers.GetEnvAsInt("SQL_PORT", 0000)
	database := helpers.GetEnv("SQL_DATABASE", "")
	dns := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database,
	)
	if db, err = gorm.Open(sqlserver.Open(dns), &gorm.Config{
		SkipDefaultTransaction: true,
	}); err != nil {
		panic(err)
	}

	if sql, err := db.DB(); err == nil {
		sql.SetMaxIdleConns(100)
		sql.SetMaxOpenConns(50)
		sql.SetConnMaxLifetime(time.Hour)
	}
	return db
}

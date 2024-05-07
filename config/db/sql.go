package db

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func GetClientGorm() *gorm.DB {
	var db *gorm.DB
	var err error
	dns := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		"", "", "", 1433, "",
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

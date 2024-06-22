package middleware

import (
	"fmt"
	"time"

	"github.com/Basillica/golang-guide/config/helpers"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func SQLMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var db *gorm.DB
		var err error
		appenv := c.MustGet("appenv").(*helpers.EnvConfig)
		dns := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
			appenv.SqlServer, appenv.SqlUser, appenv.SqlPassword, appenv.SqlPort, appenv.SqlDatabase,
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

		c.Set("sql_client", db)
	}
}

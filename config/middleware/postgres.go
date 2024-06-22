package middleware

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func PostgresMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
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

		c.Set("postgres_client", db)
	}
}

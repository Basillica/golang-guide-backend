package api

import (
	"fmt"

	"github.com/Basillica/golang-guide/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(version string) {
	r := gin.Default()
	db, err := gorm.Open(sqlite.Open("session.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	if err := db.AutoMigrate(&models.Session{}); err != nil {
		panic("unable to make migrations to sessions table")
	}

	AddAuthRoutes(r, version)
	AddUserRoutes(r, version)
	AddMiscRoutes(r, version)
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	fmt.Println("listening at port 8080 ....")
	r.Run(":8080")
}

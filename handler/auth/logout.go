package auth

import (
	"encoding/json"

	"github.com/Basillica/golang-guide/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func LogoutHandler(c *gin.Context) {
	sessionID, err := c.Cookie("session_id")
	if err != nil {
		panic("cooking missing in request")
	}

	var session models.Session
	db, err := gorm.Open(sqlite.Open("session.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.Where("id = ?", sessionID).First(&session)
	var data map[string]string
	json.Unmarshal(session.Data, &data)

	c.JSON(200, gin.H{
		"sesison": session,
	})
}

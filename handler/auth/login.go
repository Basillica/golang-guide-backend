package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Basillica/golang-guide/config/db"
	"github.com/Basillica/golang-guide/config/helpers"
	"github.com/Basillica/golang-guide/models"
	"github.com/Basillica/golang-guide/types/requests"
	"github.com/Basillica/golang-guide/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func LoginHandler(c *gin.Context) {
	var user models.User
	var req requests.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.FormatError(c, err)
		return
	}

	user.Email = req.Username

	db := db.GetClientGorm()
	if err := user.GetByAttr(db).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := utils.VerifyPassword(user.Password, req.Password); err != nil {
		c.JSON(401, gin.H{
			"error": err,
		})
		return
	}

	data := map[string]string{"username": user.Email, "user-id": user.ID}
	data_bytes, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	session := models.Session{
		ID:     uuid.NewString(),
		Data:   data_bytes,
		Expiry: time.Now().Add(24 * time.Hour).Unix(),
	}

	gormDB, err := gorm.Open(sqlite.Open("session.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	if err := gormDB.Create((&session)).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	secret := helpers.GetEnv("JWT_SECRET", "")
	if token, err := utils.GenerateJWT(user, secret); err != nil {
		c.JSON(401, gin.H{
			"error": "you are not authorized",
		})
	} else {
		utils.PersistCookie(c, session.ID, token)
		c.JSON(200, gin.H{
			"token": token,
		})
	}
}

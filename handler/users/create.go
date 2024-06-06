package users

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Basillica/golang-guide/config/db"
	"github.com/Basillica/golang-guide/config/helpers"
	"github.com/Basillica/golang-guide/handler"
	"github.com/Basillica/golang-guide/models"
	"github.com/Basillica/golang-guide/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateHandler(c *gin.Context) {
	var req models.User
	var access_token string
	var err error

	if access_token, err = handler.GetToken(c); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token provided or token missing in header",
		})
		return
	}

	secret := helpers.GetEnv("JWT_SECRET", "")
	claims, err := utils.ValidateJWT(access_token, secret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token provided",
		})
		return
	}

	fmt.Println(claims, "the user decoced jwt")

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"reason": "invalid payload provided",
		})
		return
	}

	req.CreatedAt = time.Now().Local().String()
	req.ID = uuid.New().String()
	if pass, err := utils.HashPassword(req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"reason": "inter server password",
		})
		return
	} else {
		req.Password = pass
	}

	db := db.GetClientGorm()
	if err := req.Create(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
	})
}

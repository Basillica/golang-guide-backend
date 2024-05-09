package users

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Basillica/golang-guide/config/db"
	"github.com/Basillica/golang-guide/models"
	"github.com/Basillica/golang-guide/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateHandler(c *gin.Context) {
	var req models.User

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
		fmt.Println(pass, "*****++++++++++++++")
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

package users

import (
	"net/http"

	"github.com/Basillica/golang-guide/config/db"
	"github.com/Basillica/golang-guide/models"
	"github.com/gin-gonic/gin"
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

	db := db.GetClientGorm()
	if err := req.Create(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
	})
}

package users

import (
	"fmt"
	"net/http"

	"github.com/Basillica/golang-guide/config/db"
	"github.com/Basillica/golang-guide/handler"
	"github.com/Basillica/golang-guide/models"
	"github.com/Basillica/golang-guide/utils"
	"github.com/gin-gonic/gin"
)

func DeleteHandler(c *gin.Context) {
	var req models.User
	var access_token string
	var err error
	if access_token, err = handler.GetToken(c); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token provided or token missing in header",
		})
		return
	}

	claims, err := utils.ValidateJWT(access_token, "mysecretkey")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token provided",
		})
		return
	}

	fmt.Println(claims, "the user decoced jwt")

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "request missing user id",
		})
		return
	}

	req.ID = id
	db := db.GetClientGorm()
	if err := req.Delete(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
	})
}

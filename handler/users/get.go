package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Basillica/golang-guide/config/helpers"
	"github.com/Basillica/golang-guide/datamanager"
	"github.com/Basillica/golang-guide/handler"
	"github.com/Basillica/golang-guide/models"
	"github.com/Basillica/golang-guide/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetHandler(c *gin.Context) {
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

	fmt.Println(claims, "the usre decoced jwt")
	role, _ := c.Cookie("userrole")
	scopestring, _ := c.Cookie("userscopes")
	var scopes []string
	json.Unmarshal([]byte(scopestring), &scopes)

	db := c.MustGet("sql_client").(*gorm.DB)
	manager, err := datamanager.NewUserQueryManager(db, scopes, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var users []models.User
	res := manager.Get(db).Find(&users)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": res.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success", "users": users,
	})
}

package auth

import (
	"fmt"
	"net/http"

	"github.com/Basillica/golang-guide/config/db"
	"github.com/Basillica/golang-guide/models"
	"github.com/Basillica/golang-guide/types/requests"
	"github.com/Basillica/golang-guide/utils"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var user models.User
	var req requests.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"error":  err.Error(),
		// 	"reason": "invalid payload provided",
		// })
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

	fmt.Println(user.Email, ">>>>>>>>>>>>>>>>")
	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(401, gin.H{
			"error": err,
		})
		return
	}

	fmt.Println("the hashed password: ", hash)
	if err := utils.VerifyPassword(hash, req.Password); err != nil {
		c.JSON(401, gin.H{
			"error": err,
		})
		return
	}

	secret := "mysecretkey"
	if token, err := utils.GenerateJWT(user, secret); err != nil {
		c.JSON(401, gin.H{
			"error": "you are not authorized",
		})
	} else {
		c.JSON(200, gin.H{
			"token": token,
		})
	}

}

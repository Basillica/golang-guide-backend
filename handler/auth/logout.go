package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LogoutHandler(c *gin.Context) {
	fmt.Println("you visited the login endpoint")
	c.JSON(200, "you visited thet logout page")
}

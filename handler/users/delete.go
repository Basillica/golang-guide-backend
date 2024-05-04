package users

import "github.com/gin-gonic/gin"

func DeleteHandler(c *gin.Context) {
	c.JSON(200, "you have successfully deleted the user")
}

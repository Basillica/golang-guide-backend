package misc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FormData struct {
	Name    string `form:"name"`
	Email   string `form:"email"`
	Message string `form:"message"`
}

func FormHandler(c *gin.Context) {
	var req FormData
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"name":    req.Name,
		"email":   req.Email,
		"message": req.Message,
	})
}

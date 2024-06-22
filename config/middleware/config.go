package middleware

import (
	"net/http"

	"github.com/Basillica/golang-guide/config/helpers"
	"github.com/gin-gonic/gin"
)

func ConfigMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		env, err := helpers.LoadConfig(".")
		if err != nil {
			c.Abort()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Set("appenv", env)
	}
}

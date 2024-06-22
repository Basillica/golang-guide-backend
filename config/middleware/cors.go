package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	allowedHeaders := "Content-Type, Content-Length, Accept-Encoding, User-Scope, User-Role, Authorization, accept, origin, Cache-Control, X-Requested-With, content-disposition"
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "localhost")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, HEAD, DELETE")

		if c.Request.Method == "options" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

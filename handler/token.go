package handler

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) (string, error) {
	token := c.Request.Header.Get("Authorization")
	access_token := strings.TrimPrefix(token, "Bearer ")
	if access_token == "" || access_token == token {
		return "", errors.New("token missing in header of request")
	}
	return access_token, nil
}

package api

import (
	"github.com/Basillica/golang-guide/handler/auth"
	"github.com/gin-gonic/gin"
)

func AddAuthRoutes(r *gin.Engine, version string) {
	g := r.Group(version + "/auth")
	g.POST("/login", auth.LoginHandler)
	g.GET("/logout", auth.LogoutHandler)
}

package api

import (
	"github.com/Basillica/golang-guide/handler/users"
	"github.com/gin-gonic/gin"
)

func AddUserRoutes(r *gin.Engine, version string) {
	g := r.Group(version + "/user")
	g.POST("/create", users.CreateHandler)
	g.DELETE("/delete", users.DeleteHandler)
}

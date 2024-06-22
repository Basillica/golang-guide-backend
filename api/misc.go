package api

import (
	"github.com/Basillica/golang-guide/handler/misc"
	"github.com/gin-gonic/gin"
)

func AddMiscRoutes(r *gin.Engine, version string) {
	g := r.Group(version + "/misc")
	g.POST("/health", misc.HealthHandler)
}

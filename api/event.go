package api

import (
	"github.com/Basillica/golang-guide/handler/event"
	"github.com/gin-gonic/gin"
)

func AddEventRoutes(r *gin.Engine, version string) {
	g := r.Group(version + "/events")
	g.POST("/form", event.FormHandler)
	g.GET("/ws", event.EventsHandler)
}

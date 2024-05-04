package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func New(version string) {
	r := gin.Default()
	AddAuthRoutes(r, version)
	AddUserRoutes(r, version)
	fmt.Println("listening at port 8080 ....")
	r.Run(":8080")
}

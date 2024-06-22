package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Basillica/golang-guide/config/middleware"
	"github.com/Basillica/golang-guide/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type API struct {
	Engine *gin.Engine
	svr    *http.Server
}

func (a *API) Start(r *gin.Engine) {
	go func() {
		if err := a.svr.ListenAndServe(); err != nil {
			log.Panicf("listen: %s\n", err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := a.svr.Shutdown(ctx); err != nil {
		log.Panicf("Server shut down: .... %s\n", err)
	}
	log.Println("Server is exiting ..........")
}

func (a *API) Init(version string) (*gin.Engine, error) {
	r := gin.New()
	r.Use(a.WithMiddleware()...)
	a.Engine = r
	// add routes
	a.AddRoutes(version)
	a.WithServer()

	return r, nil
}

func (a *API) WithServer() API {
	ch := make(chan API)
	go func() {
		a.svr = &http.Server{
			Addr:         fmt.Sprintf(":%s", "80"),
			Handler:      a.Engine,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		ch <- *a
	}()

	return <-ch
}

func (a *API) AddRoutes(version string) {
	AddAuthRoutes(a.Engine, version)
	AddUserRoutes(a.Engine, version)
	AddEventRoutes(a.Engine, version)
	AddMiscRoutes(a.Engine, version)
}

func (a *API) WithMiddleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		gin.Recovery(), gin.Logger(),
		middleware.ConfigMiddleware(),
		middleware.CorsMiddleware(),
	}
}

func New(version string) API {
	api := API{}
	r, err := api.Init(version)
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("session.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	if err := db.AutoMigrate(&models.Session{}); err != nil {
		panic("unable to make migrations to sessions table")
	}

	api.Engine = r
	return api
}

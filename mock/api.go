package mock_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Basillica/golang-guide/api"
	"github.com/Basillica/golang-guide/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type ApiTest struct {
	router *gin.Engine
	t      *testing.T
}

func ApiTestHandler(t *testing.T) {
	router := api.New("v0").Engine
	api := ApiTest{
		router: router,
		t:      t,
	}

	// test misc endpoints
	api.TestMiscRoute()
	// non existent endpoints
	api.Test404Route()
	// test auth endpoints
	api.TestAuthRoute()

}

func (a *ApiTest) TestMiscRoute() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v0/misc/health", http.NoBody)
	a.router.ServeHTTP(w, req)

	assert.Equal(a.t, http.StatusOK, w.Code)
	assert.Equal(a.t, "{\"message\": \"pong\"}", w.Body.String())
}

func (a *ApiTest) Test404Route() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v0/misc/nonexsistent", http.NoBody)
	a.router.ServeHTTP(w, req)

	assert.Equal(a.t, http.StatusOK, w.Code)
}

func (a *ApiTest) TestAuthRoute() {
	db, mock, sqlDB := NewDatabase()
	defer sqlDB.Close()

	db.Migrator().CreateTable(&models.User{})

	mock.ExpectBegin()
	mock.ExpectExec("CREATE TABLE users(.*)")
	mock.ExpectCommit()

}

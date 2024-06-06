package utils

import (
	"net/http"

	"github.com/Basillica/golang-guide/models"
	"github.com/gin-gonic/gin"
)

func PersistCookie(c *gin.Context, user models.User, session_id, token string) {
	c.SetCookie("scope", user.Scope, 86400, "/", "localhost", false, true)
	c.SetCookie("userrole", user.UserRole, 86400, "/", "localhost", false, true)
	c.SetCookie("userscopes", user.UserScopes.String(), 86400, "/", "localhost", false, true)
	c.SetCookie("access_token", token, 86400, "/", "localhost", false, true)
	c.SetCookie("cookie_exp", "86400", 86400, "/", "localhost", false, true)
	c.SetCookie("session_id", session_id, 3600, "/", "", false, true)
	c.SetSameSite(http.SameSiteStrictMode)
}

func RemoveCookie(c *gin.Context) {
	c.SetCookie("access_token", "", 86400, "/", "localhost", false, true)
	c.SetCookie("cookie_exp", "", 86400, "/", "localhost", false, true)
	c.SetCookie("session_id", "", 86400, "/", "", false, true)
	c.SetSameSite(http.SameSiteStrictMode)
}

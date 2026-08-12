package login

import (
	"github.com/barbar17/finance-app/internal/auth"
	"github.com/gin-gonic/gin"
)

func LoginRoutes(r *gin.Engine, sessionStore *auth.SessionStore) {
	r.GET("/login", loginPageHandler)
	r.POST("/login", loginPostHandler(sessionStore))
}

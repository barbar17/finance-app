package login

import (
	"net/http"

	"github.com/barbar17/finance-app/internal/auth"
	"github.com/barbar17/finance-app/internal/db"
	"github.com/barbar17/finance-app/internal/user"
	"github.com/gin-gonic/gin"
)

func loginPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login/login.html", nil)
}

func loginPostHandler(sessionStore *auth.SessionStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		uname := c.PostForm("username")
		pass := c.PostForm("password")

		userRepo := user.NewRepo(db.DB)

		u, err := userRepo.FindByUsername(c, uname)
		if err != nil || pass != u.Password {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"msg":     "username or password is wrong",
			})
			return
		}

		sessionID, err := sessionStore.Create(u.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"msg":     "failed to create session",
			})
		}

		c.SetCookie("sesid", sessionID, 86400, "/", "", false, true)

		c.JSON(http.StatusOK, gin.H{
			"success":  true,
			"redirect": "/",
		})
		return
	}
}

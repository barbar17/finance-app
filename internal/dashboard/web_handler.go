package dashboard

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func dashboardWebHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"Title": "Go Finance App",
	})
}

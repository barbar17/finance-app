package dashboard

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func dashboardPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard/dashboard.html", gin.H{
		"Title": "Go Finance App",
	})
}

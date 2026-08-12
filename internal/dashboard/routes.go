package dashboard

import (
	"github.com/gin-gonic/gin"
)

func DashboardRoutes(r *gin.Engine) {
	r.GET("/", dashboardWebHandler)
}

package dashboard

import (
	"github.com/barbar17/finance-app/internal/utils"
	"github.com/gin-gonic/gin"
)

func dashboardPageHandler(c *gin.Context) {
	utils.Render(c, "dashboard", gin.H{"Title": "Dashboard"})
}

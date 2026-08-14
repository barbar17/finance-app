package errorpage

import (
	"github.com/barbar17/finance-app/internal/utils"
	"github.com/gin-gonic/gin"
)

func NotFoundPageHandler(c *gin.Context) {
	utils.Render(c, "404", gin.H{"Title": "404"})
}

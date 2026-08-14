package errorpage

import "github.com/gin-gonic/gin"

func ErrorPageRoutes(r *gin.Engine) {
	r.NoRoute(NotFoundPageHandler)
}

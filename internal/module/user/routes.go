package user

import "github.com/gin-gonic/gin"

func UserRoutes(r *gin.Engine) {
	r.GET("/user", UserPageHandler)

	api := r.Group("/api")
	{
		api.GET("/users", GetUserTable)
	}
}

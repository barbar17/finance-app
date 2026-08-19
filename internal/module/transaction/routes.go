package transaction

import "github.com/gin-gonic/gin"

func TransactionRoutes(r *gin.Engine) {
	r.GET("/transaction", TransactionPageHandler)

	api := r.Group("/api")
	{
		api.GET("/transactions", GetTransactionTable)
		api.POST("/transaction", Create)
	}
}

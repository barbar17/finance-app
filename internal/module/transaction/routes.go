package transaction

import "github.com/gin-gonic/gin"

func TransactionRoutes(r *gin.Engine) {
	r.GET("/transaction", transactionPageHandler)
}

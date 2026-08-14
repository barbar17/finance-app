package transaction

import (
	"github.com/barbar17/finance-app/internal/utils"
	"github.com/gin-gonic/gin"
)

func transactionPageHandler(c *gin.Context) {
	utils.Render(c, "transaction", gin.H{"Title": "Transaction"})
}

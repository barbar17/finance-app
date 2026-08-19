package transaction

import (
	"math"
	"net/http"
	"strconv"

	"github.com/barbar17/finance-app/internal/db"
	"github.com/barbar17/finance-app/internal/types"
	"github.com/barbar17/finance-app/internal/utils"
	"github.com/gin-gonic/gin"
)

func TransactionPageHandler(c *gin.Context) {
	utils.Render(c, "transaction", gin.H{"Title": "Transaction"})
}

func GetTransactionTable(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "page query is not a number",
			"err":     err,
		})
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "limit query is not a number",
			"err":     err,
		})
	}

	tableParams := types.TableParams{
		Search: c.Query("search"),
		Sort:   c.Query("sort"),
		Order:  c.Query("order"),
		Limit:  limit,
		Offset: (pageInt - 1) * limitInt,
	}

	transactionRepo := NewRepo(db.DB)

	transactions, total, err := transactionRepo.GetTransactionTable(c, tableParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"msg":     "failed to get transactions",
			"err":     err,
		})
	}

	totalPages := int(math.Ceil(float64(total) / float64(limitInt)))
	c.JSON(http.StatusOK, gin.H{
		"data": transactions,
		"pagination": gin.H{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"totalPages": totalPages,
		},
	})
}

func Create(c *gin.Context) {
	var body struct {
		Name     string `form:"name" binding:"required"`
		Type     string `form:"type" binding:"required"`
		Amount   int    `form:"amount" binding:"required"`
		Desc     string `form:"desc" binding:"required"`
		Category string `form:"category" binding:"required"`
	}

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "invalid request body",
			"err":     err,
		})
	}

	var transaction CreateTransaction
	transaction.Name = body.Name
	transaction.Desc = body.Desc
	transaction.Category = body.Category
	if body.Type == "income" {
		transaction.Amount = body.Amount
	} else {
		transaction.Amount = -body.Amount
	}

	transactionRepo := NewRepo(db.DB)

	if err := transactionRepo.Create(c, transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"msg":     "failed to create transaction",
			"err":     err,
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"msg":     "transaction created",
	})
}

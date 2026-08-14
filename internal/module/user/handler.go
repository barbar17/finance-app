package user

import (
	"math"
	"net/http"
	"strconv"

	"github.com/barbar17/finance-app/internal/db"
	"github.com/barbar17/finance-app/internal/types"
	"github.com/barbar17/finance-app/internal/utils"
	"github.com/gin-gonic/gin"
)

func UserPageHandler(c *gin.Context) {
	utils.Render(c, "user", gin.H{"Title": "User"})
}

func GetUsersTable(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "page query is not a number",
		})
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "limit query is not a number",
		})
		return
	}

	search := c.Query("search")
	sort := c.Query("sort")
	order := c.Query("order")
	offset := (pageInt - 1) * limitInt
	tableParams := types.TableParams{Search: search, Sort: sort, Order: order, Limit: limit, Offset: offset}

	userRepo := NewRepo(db.DB)

	users, total, err := userRepo.GetUsersTable(c, tableParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"msg":     "failed to get users",
		})
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(limitInt)))

	c.JSON(http.StatusOK, gin.H{
		"data": users,
		"pagination": gin.H{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"totalPages": totalPages,
		},
	})
}

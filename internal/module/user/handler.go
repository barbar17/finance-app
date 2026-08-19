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

func GetUserTable(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "page query is not a number",
			"err":     err,
		})
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "limit query is not a number",
			"err":     err,
		})
		return
	}

	tableParams := types.TableParams{
		Search: c.Query("search"),
		Sort:   c.Query("sort"),
		Order:  c.Query("order"),
		Limit:  limit,
		Offset: (pageInt - 1) * limitInt,
	}

	userRepo := NewRepo(db.DB)

	users, total, err := userRepo.GetUserTable(c, tableParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"msg":     "failed to get users",
			"err":     err,
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

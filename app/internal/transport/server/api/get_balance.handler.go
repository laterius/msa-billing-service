package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	s "github.com/laterius/service_architecture_hw3/app/internal/service"
	"net/http"
)

// GetBalanceHandler handles request to get order by ID
func GetBalanceHandler(service s.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("userId")

		userId, err := uuid.Parse(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "user id can't parse",
				"data":    gin.H{},
			})
			return
		}

		account, err := service.GetAccount(userId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "account not found",
				"data":    gin.H{},
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "account found",
			"data": gin.H{
				"account": account.Balance,
			},
		})

	}
}

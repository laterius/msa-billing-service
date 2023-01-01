package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/laterius/service_architecture_hw3/app/internal/service"
	"net/http"
)

// ChangeBalanceHandler handles request to make payment
func ChangeBalanceHandler(service service.Service) func(c *gin.Context) {
	// Request body structure
	type Body struct {
		UserID uuid.UUID `json:"user_id"`
		Amount int       `json:"amount"`
	}

	return func(c *gin.Context) {
		body := Body{}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": fmt.Sprintf("Problem with parse body of request. Error = %s", err.Error()),
				"data":    gin.H{},
			})
			return
		}

		account, err := service.GetAccount(body.UserID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "payment not found",
				"data":    gin.H{},
			})
			return
		}

		account.Balance += body.Amount
		if err = service.ChangeBalance(account); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": fmt.Sprintf("Problem with change balance. Error = %s", err.Error()),
				"data":    gin.H{},
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
			"data": gin.H{
				"account": account.Balance,
			},
		})
	}
}

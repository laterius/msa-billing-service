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
		Amount int `json:"amount"`
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

		if account.Balance+body.Amount < 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "There are not enough funds on the balance sheet",
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
				"balance": account.Balance,
			},
		})
	}
}

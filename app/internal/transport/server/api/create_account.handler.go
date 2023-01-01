package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/laterius/service_architecture_hw3/app/internal/service"
	"net/http"
)

// CreateAccount handles request to create account
func CreateAccount(service service.Service) func(c *gin.Context) {
	// Request body structure
	type Body struct {
		UserId uuid.UUID `json:"user_id"`
	}

	return func(c *gin.Context) {
		body := Body{}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
				"data":    gin.H{},
			})
			return
		}

		if err := service.CreateAccount(body.UserId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
				"data":    gin.H{},
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
			"data":    gin.H{},
		})
	}
}

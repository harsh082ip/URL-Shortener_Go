package authcontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/URL-Shortener_Go/models"
)

func SignUp(c *gin.Context) {

	var jsonData models.User
	// Bind and validate JSON
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		// Return a bad request response if there's an error in binding/validation
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Error in binding/validation",
			"error":  err.Error(),
		})
		return
	}

}

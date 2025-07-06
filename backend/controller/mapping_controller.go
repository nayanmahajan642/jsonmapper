package controller

import (
	"net/http" // HTTP status codes like 400, 200

	"github.com/gin-gonic/gin"
	"github.com/nayanmahajan642/jsonmapper/models"
	"github.com/nayanmahajan642/jsonmapper/service"
)

func MapHandler(c *gin.Context) {
	// Create a variable to store parsed JSON request
	var req models.InputRequest

	// Bind incoming JSON to InputRequest struct (auto handles decoding)
	if err := c.ShouldBindJSON(&req); err != nil {
		// Return 400 Bad Request if JSON is not in expected format
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Call service layer to perform mapping logic
	output, err := service.MapService(req.RequestJSON, req.RequestMapping)

	// Handle error from service layer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return final mapped output as JSON response
	c.JSON(http.StatusOK, output)
}

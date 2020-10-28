package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSON serializes the api response properly to json
func JSON(c *gin.Context, message string, status int, data interface{}, errs error) {
	responsedata := gin.H{
		"message": message,
		"data":    data,
		"errors":  errs,
		"status":  http.StatusText(status),
	}

	c.JSON(status, responsedata)
}

package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vdamery/jdria/internal/pkg/services"
	"github.com/vdamery/jdria/pkg/api"
	"net/http"
)

func NoRoute(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": "Not Found"})
}

func Health() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK", "services": gin.H{}})
	}
}

func Test(internalBootService services.InternalBootService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Example of internal service usage
		var request api.TestRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "json decoding : " + err.Error(), "status": http.StatusBadRequest})
			return
		}
		resp, err := internalBootService.Test(request.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internalBootService.Test : " + err.Error(), "status": http.StatusInternalServerError})
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}

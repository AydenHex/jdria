package handlers

import (
	"github.com/figarocms/hr-go-utils/v2/logger"
	"github.com/gin-gonic/gin"
	"github.com/vdamery/jdria/internal/pkg/services"
	"github.com/vdamery/jdria/pkg/api"
	"go.uber.org/zap"
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

func Index(gameService services.GameService) gin.HandlerFunc {
	return func(c *gin.Context) {
		game, err := gameService.StartGame()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": http.StatusInternalServerError})
			return
		}
		logger.Log.Info("Starting game", zap.Any("game", game))
		c.HTML(http.StatusOK, "index.gohtml", game)
	}
}

func Send(gameService services.GameService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request api.SendRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": http.StatusBadRequest})
			return
		}

		game, err := gameService.Send(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": http.StatusInternalServerError})
			return
		}
		logger.Log.Info("Send", zap.Any("game", game))
		c.JSON(http.StatusOK, game)
	}
}

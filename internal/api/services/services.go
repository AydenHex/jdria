package services

import (
	"github.com/figarocms/hr-go-utils/v2/logger"
	"github.com/vdamery/jdria/configs"
	"github.com/vdamery/jdria/internal/pkg/services"
)

type Services struct {
	GamerService services.GameService
}

func InitServices(config configs.Api) *Services {
	logger.Log.Info("services loaded")

	// Service interne de test
	gameService := services.NewInternalBootService()

	return &Services{
		GamerService: gameService,
	}
}

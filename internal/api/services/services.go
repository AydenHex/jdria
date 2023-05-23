package services

import (
	"github.com/figarocms/hr-go-utils/v2/logger"
	"github.com/vdamery/jdria/configs"
	"github.com/vdamery/jdria/internal/pkg/services"
)

type Services struct {
	InternalBootService services.InternalBootService
}

func InitServices(config configs.Api) *Services {
	logger.Log.Info("services loaded")

	// Service interne de test
	internalBootService := services.NewInternalBootService()

	return &Services{
		InternalBootService: internalBootService,
	}
}

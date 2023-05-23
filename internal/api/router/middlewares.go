package router

import (
	"github.com/figarocms/hr-go-utils/v2/logger"
	fcmsrouter "github.com/figarocms/hr-go-utils/v2/logger/gin"
	"github.com/figarocms/hr-go-utils/v2/tracer"
	"github.com/gin-gonic/gin"
)

func initializeMiddlewares(router *gin.Engine) {

	tracer.Tracer.UseTraceMiddleware(router)

	// use recovery with zap
	router.Use(fcmsrouter.RecoveryWithZap(logger.Log, true))

}

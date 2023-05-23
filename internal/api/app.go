package api

import (
	"context"
	"fmt"
	"github.com/figarocms/hr-go-utils/v2/logger"
	"github.com/figarocms/hr-go-utils/v2/tracer"
	"github.com/gin-gonic/gin"
	"github.com/vdamery/jdria/configs"
	"github.com/vdamery/jdria/internal/api/router"
	"github.com/vdamery/jdria/internal/api/services"
	"go.uber.org/zap"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var Logger *logger.Logger

type App struct {
	config   configs.Api
	router   *gin.Engine
	services *services.Services
}

func New() *App {
	app := &App{}
	app.setup()
	return app
}

func (app *App) setup() {
	config := configs.LoadApi()

	Logger = logger.NewLogger()
	Logger.SetVersion(config.AppVersion)

	tracer.Tracer.Enabled = config.TracerEnabled
	tracer.Tracer.ApplicationName = config.Name

	svcs := services.InitServices(config)

	r := router.InitializeRouter(svcs)

	app.config = config
	app.router = r
	app.services = svcs

}

func (app *App) Run() {
	if app.config.TracerEnabled {
		if err := profiler.Start(); err != nil {
			logger.Log.Fatal("failed to start profiler", zap.Error(err))
		}
		defer profiler.Stop()
		tracer.Tracer.Start()
		defer tracer.Tracer.Stop()
	}

	logger.Log.Info(fmt.Sprintf("RUN APP on PORT %d", app.config.Port))

	// https://gin-gonic.com/docs/examples/graceful-restart-or-stop/
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.Port),
		Handler: app.router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal("Run ListenAndServe", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatal("Run Shutdown", zap.Error(err))
	}

	logger.Log.Info("Server exiting")
}

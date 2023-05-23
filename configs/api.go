package configs

import (
	"fmt"
	"github.com/figarocms/hr-go-utils/v2/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/joho/godotenv"
	"time"
)

type Api struct {
	errors     []error
	AppVersion string `json:"appVersion"`

	Port          int    `json:"port"`
	Name          string `json:"name"`
	TracerEnabled bool   `json:"tracerEnabled"`
}

func LoadApi() Api {
	_ = godotenv.Load(".env")

	viper.AutomaticEnv()

	c := Api{}
	c.AppVersion = c.getStringWithDefault("APP_VERSION", "wip")
	c.Port = c.getIntWithDefault("API_PORT", 8088)
	c.Name = c.getStringWithDefault("NAME", "hr-bootapi-template")
	c.TracerEnabled = c.getBooleanWithDefault("API_TRACER_ENABLED", false)

	env := c.getStringWithDefault("ENV", "dev")
	if env == "dev" {
		logger.Log.Info("LoadApi dump config", zap.Any("config", c))
	}

	if len(c.errors) != 0 {
		logger.Log.Fatal("LoadApi config", zap.Errors("errors", c.errors))
	}
	return c

}

func (c *Api) getIntWithDefault(key string, defaultValue int) int {
	viper.SetDefault(key, defaultValue)
	return viper.GetInt(key)
}

func (c *Api) getBooleanWithDefault(key string, defaultValue bool) bool {
	viper.SetDefault(key, defaultValue)
	return viper.GetBool(key)
}

func (c *Api) getMandatoryString(key string) (value string) {
	if value = viper.GetString(key); value == "" {
		c.errors = append(c.errors, fmt.Errorf("cannot find configuration for key %s", key))
	}
	return value
}

func (c *Api) getStringWithDefault(key, defaultValue string) string {
	viper.SetDefault(key, defaultValue)
	return viper.GetString(key)
}

func (c *Api) getDurationWithDefault(key, defaultValue string) time.Duration {
	viper.SetDefault(key, defaultValue)
	return viper.GetDuration(key)
}

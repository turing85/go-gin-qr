package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mrsimonemms/gin-structured-logger"
	"github.com/rs/zerolog/log"
)

type EngineConfig interface {
	HealthPath() string
	MetricsPath() string
}

func SetupEngine(config EngineConfig) *gin.Engine {
	engine := initializeEngine(
		ginstructuredlogger.New(),
		gin.Recovery())
	initializeHealthPath(engine, config.HealthPath())
	initializeMetrics(engine, config.MetricsPath())
	return engine
}

func initializeEngine(middleware ...gin.HandlerFunc) *gin.Engine {
	engine := gin.New()
	engine.Use(middleware...)
	if err := engine.SetTrustedProxies(nil); err != nil {
		log.Error().Msgf(`Configuring trusted proxies failed. Terminating. Cause: %s`, err)
		panic(err)
	}
	return engine
}

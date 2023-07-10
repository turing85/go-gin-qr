package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/tavsec/gin-healthcheck"
	"github.com/tavsec/gin-healthcheck/checks"
	healthconfig "github.com/tavsec/gin-healthcheck/config"
)

func initializeHealthPath(engine *gin.Engine, healthPath string) {
	err := gin_healthcheck.New(
		engine,
		healthconfig.Config{
			HealthPath:  healthPath,
			Method:      http.MethodGet,
			StatusOK:    http.StatusOK,
			StatusNotOK: http.StatusServiceUnavailable,
		},
		[]checks.Check{})
	if err != nil {
		log.Warn().Msgf(`Failed to initialize health checks. Cause: %s`, err)
	}
}

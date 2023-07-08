package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/tavsec/gin-healthcheck"
	"github.com/tavsec/gin-healthcheck/checks"
	healthconfig "github.com/tavsec/gin-healthcheck/config"

	"go-gin-qr/internal/config"
)

func addHealthChecks(engine *gin.Engine) {
	err := gin_healthcheck.New(
		engine,
		healthconfig.Config{
			HealthPath:  config.GetConfig().Health().Path(),
			Method:      http.MethodGet,
			StatusOK:    http.StatusOK,
			StatusNotOK: http.StatusServiceUnavailable,
		},
		[]checks.Check{})
	if err != nil {
		log.Warn().Msgf(`Failed to start health checks. Cause: %s`, err)
	}
}

package middleware

import (
	"net/http"

	"go-gin-qr/config"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/tavsec/gin-healthcheck"
	"github.com/tavsec/gin-healthcheck/checks"
	ginhealthcheckconfig "github.com/tavsec/gin-healthcheck/config"
)

func addHealthChecks(engine *gin.Engine) {
	err := gin_healthcheck.New(
		engine,
		ginhealthcheckconfig.Config{
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

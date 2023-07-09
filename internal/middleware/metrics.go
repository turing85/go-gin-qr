package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"

	"go-gin-qr/internal/config"
)

func initializeMetrics(engine *gin.Engine) {
	metrics := ginmetrics.GetMonitor()
	metrics.SetMetricPath(config.GetConfig().Metrics().Path())
	metrics.Use(engine)
}

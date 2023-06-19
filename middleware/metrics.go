package middleware

import (
	"go-gin-qr/config"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

func initializeMetrics(engine *gin.Engine) *gin.Engine {
	metrics := ginmetrics.GetMonitor()
	metrics.SetMetricPath(config.GetConfig().Metrics().Path())
	metrics.Use(engine)
	return engine
}

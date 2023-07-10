package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

func initializeMetrics(engine *gin.Engine, metricsPath string) {
	metrics := ginmetrics.GetMonitor()
	metrics.SetMetricPath(metricsPath)
	metrics.Use(engine)
}

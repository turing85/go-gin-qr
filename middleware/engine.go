package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func InitializeEngine(middleware ...gin.HandlerFunc) *gin.Engine {
	engine := gin.New()
	engine.Use(middleware...)
	if err := engine.SetTrustedProxies(nil); err != nil {
		log.Error().Msg(fmt.Sprintf(`Configuring trusted proxies failed. Terminating. %s`, err))
		panic(err)
	}
	InitializeMetrics(engine)
	return engine
}

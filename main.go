package main

import (
	"bytes"
	"fmt"
	"go-gin-qr/internal/endpoints"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"

	"go-gin-qr/internal/appconfig"
	"go-gin-qr/internal/middleware"
)

func main() {
	appConfig := appconfig.GetConfig()
	printConfigInDebugMode(appConfig)
	engine := endpoints.AddQrEndpoint(middleware.SetupEngine(appConfig), appConfig.Qr().Path())
	port := appconfig.GetConfig().Http().Port()
	host := appconfig.GetConfig().Http().Host()
	log.Info().Msgf(`Starting HTTP Server on %s:%d`, host, port)
	if err := engine.Run(fmt.Sprintf(`%s:%d`, host, port)); err != nil {
		log.Panic().Msgf(`Failure during shutdown: %e`, err)
		panic(err)
	}
}

func printConfigInDebugMode(appConfig appconfig.Config) {
	if gin.Mode() == gin.DebugMode {
		buffer := new(bytes.Buffer)
		err := yaml.NewEncoder(buffer).Encode(appConfig)
		if err != nil {
			log.Error().Msgf(`Unable to convert config to yaml. Cause: %s`, err)
		}
		log.Info().Msgf(`Config: \n%s`, buffer.String())
	}
}

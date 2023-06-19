package main

import (
	"fmt"

	"go-gin-qr/config"
	"go-gin-qr/endpoint"
	"go-gin-qr/middleware"

	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Println(config.GetConfig())
	engine := middleware.SetupEngine()
	endpoint.AddQrEndpoint(engine)
	port := config.GetConfig().Http().Port()
	log.Info().Msgf(`Starting HTTP Server on port %d`, port)
	_ = engine.Run(fmt.Sprintf(":%d", port))
}

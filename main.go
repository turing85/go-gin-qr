package main

import (
	"fmt"
	"net/http"

	"go-gin-qr/config"
	"go-gin-qr/middleware"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/skip2/go-qrcode"
)

func main() {
	engine := SetupEngine()
	_ = engine.Run(fmt.Sprintf(":%d", config.GetConfig().Http.Port))
}

func SetupEngine() *gin.Engine {
	engine := middleware.InitializeEngine(
		middleware.DefaultStructuredLogger(),
		gin.Recovery())

	engine.GET(fmt.Sprintf(`%s/:data`, config.GetConfig().Qr.Path), getQrCode)
	return engine
}

func getQrCode(context *gin.Context) {
	var text = context.Param("data")
	log.Info().Msg(fmt.Sprintf(`Generating QR-code for string "%s"`, text))
	png, err := qrcode.Encode(text, qrcode.Medium, 250)
	if err == nil {
		context.Header(http.CanonicalHeaderKey("Content-Disposition"), "inline")
		context.Data(200, "image/png", png)
	} else {
		log.Error().Msg(fmt.Sprintf(`Error-Code: %d`, err))
		context.String(500, "Internal server error")
	}
}

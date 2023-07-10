package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrsimonemms/gin-structured-logger"
	"github.com/rs/zerolog/log"
	"github.com/skip2/go-qrcode"
)

func AddQrEndpoint(engine *gin.Engine, qrPath string) *gin.Engine {
	engine.GET(qrPath, getQrCode)
	return engine
}

func getQrCode(context *gin.Context) {
	var text = context.Query("data")
	log.Info().Msgf(`Generating QR-code for string "%s"`, text)
	png, err := qrcode.Encode(text, qrcode.Medium, 250)
	if err != nil {
		ginstructuredlogger.Get(context).Error().Msgf(`Error-Code: %d`, err)
		context.String(500, "Internal server error")
	} else {
		context.Header(http.CanonicalHeaderKey("Content-Disposition"), "inline")
		context.Data(200, "image/png", png)
	}
}

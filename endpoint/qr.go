package endpoint

import (
	"github.com/gin-gonic/gin"
	ginstructuredlogger "github.com/mrsimonemms/gin-structured-logger"
	"github.com/rs/zerolog/log"
	"net/http"

	"go-gin-qr/config"

	"github.com/skip2/go-qrcode"
)

func AddQrEndpoint(engine *gin.Engine) {
	engine.GET(config.GetConfig().Qr().Path(), getQrCode)
}

func getQrCode(context *gin.Context) {
	var text = context.Query("data")
	log.Info().Msgf(`Generating QR-code for string "%s"`, text)
	if png, err := qrcode.Encode(text, qrcode.Medium, 250); err == nil {
		context.Header(http.CanonicalHeaderKey("Content-Disposition"), "inline")
		context.Data(200, "image/png", png)
	} else {
		ginstructuredlogger.Get(context).Error().Msgf(`Error-Code: %d`, err)
		context.String(500, "Internal server error")
	}
}

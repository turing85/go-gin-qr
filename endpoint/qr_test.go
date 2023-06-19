package endpoint

import (
	"bytes"
	"fmt"
	"image"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-gin-qr/config"
	"go-gin-qr/middleware"

	"github.com/liyue201/goqr"
	"github.com/stretchr/testify/assert"
)

func TestQrRoute(t *testing.T) {
	// GIVEN
	expected := "Hello, world!"
	engine := middleware.SetupEngine()
	AddQrEndpoint(engine)
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(`%s?data=%s`, config.GetConfig().Qr().Path(), expected),
		nil)
	if err != nil {
		assert.Error(t, err, "Failed to create HTTP request")
	}

	// WHEN
	engine.ServeHTTP(recorder, req)

	// THEN
	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "inline", recorder.Header().Get(http.CanonicalHeaderKey("Content-Disposition")))
	assert.Equal(t, "image/png", recorder.Header().Get(http.CanonicalHeaderKey("Content-Type")))
	assert.NotEmpty(t, t, recorder.Body)
	qrCodeData := extractQrData(t, recorder, err)
	assert.Len(t, qrCodeData, 1)
	actual := string(qrCodeData[0].Payload)
	assert.Equal(t, expected, actual)
}

func extractQrData(t *testing.T, recorder *httptest.ResponseRecorder, err error) []*goqr.QRData {
	data := recorder.Body.Bytes()
	dataImage, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		assert.Error(t, err, "Error during image decoding")
	}
	qrCodeData, err := goqr.Recognize(dataImage)
	if err != nil {
		assert.Error(t, err, "Error during QR code recognition")
	}
	return qrCodeData
}

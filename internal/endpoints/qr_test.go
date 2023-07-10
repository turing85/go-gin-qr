package endpoints

import (
	"bytes"
	"fmt"
	"image"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/liyue201/goqr"
	"github.com/stretchr/testify/assert"

	"go-gin-qr/internal/appconfig"
	"go-gin-qr/internal/middleware"
)

func TestQrRoute(t *testing.T) {
	// GIVEN
	const expected = "Hello, world!"
	appConfig := appconfig.GetConfig()
	engine := AddQrEndpoint(middleware.SetupEngine(appConfig), appConfig.Qr().Path())
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(`%s?data=%s`, appConfig.Qr().Path(), expected),
		nil)
	if err != nil {
		t.Fatalf("Failure during HTTP request creation: %s", err)
	}

	// WHEN
	engine.ServeHTTP(recorder, req)

	// THEN
	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "inline", recorder.Header().Get(http.CanonicalHeaderKey("Content-Disposition")))
	assert.Equal(t, "image/png", recorder.Header().Get(http.CanonicalHeaderKey("Content-Type")))
	assert.NotEmpty(t, t, recorder.Body)
	qrCodeData := extractQrData(t, recorder.Body.Bytes())
	assert.Len(t, qrCodeData, 1)
	actual := string(qrCodeData[0].Payload)
	assert.Equal(t, expected, actual)
}

func extractQrData(t *testing.T, qrImageData []byte) []*goqr.QRData {
	dataImage, _, err := image.Decode(bytes.NewReader(qrImageData))
	if err != nil {
		t.Fatalf("Failure during image decoding: %s", err)
	}
	qrCodeData, err := goqr.Recognize(dataImage)
	if err != nil {
		t.Fatalf("Failure during QR code recognition: %s", err)
	}
	return qrCodeData
}

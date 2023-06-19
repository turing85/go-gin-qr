package main

import (
	"bytes"
	"fmt"
	"image"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-gin-qr/config"

	"github.com/liyue201/goqr"
	"github.com/stretchr/testify/assert"
)

func TestQrRoute(t *testing.T) {
	// GIVEN
	expected := "Hello, world!"
	router := SetupEngine()
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s?data=%s", config.GetConfig().Qr.Path, expected), nil)
	if err != nil {
		assert.FailNow(t, "Failed to create HTTP request", err)
	}

	// WHEN
	router.ServeHTTP(recorder, req)

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
		assert.FailNow(t, "Error during image decoding", err)
	}
	qrCodeData, err := goqr.Recognize(dataImage)
	if err != nil {
		assert.FailNow(t, "Error during QR code recognition", err)
	}
	return qrCodeData
}

func TestMetricsRoute(t *testing.T) {
	// GIVEN
	metricsPath := config.GetConfig().Metrics.Path
	router := SetupEngine()
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, metricsPath, nil)
	if err != nil {
		assert.FailNow(t, "Failed to create HTTP request", err)
	}

	// WHEN
	router.ServeHTTP(recorder, req)

	// THEN
	assert.Equal(t, 200, recorder.Code)
	assert.Contains(t, recorder.Header().Get(http.CanonicalHeaderKey("Content-Type")), "text/plain")
	assert.NotEmpty(t, recorder.Body)
}

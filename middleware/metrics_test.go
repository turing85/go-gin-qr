package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go-gin-qr/config"

	"github.com/stretchr/testify/assert"
)

func TestMetricsRoute(t *testing.T) {
	// GIVEN
	router := SetupEngine()
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, config.GetConfig().Metrics().Path(), nil)
	if err != nil {
		assert.Error(t, err, "Failed to create HTTP request")
	}

	// WHEN
	router.ServeHTTP(recorder, req)

	// THEN
	assert.Equal(t, 200, recorder.Code)
	assert.Contains(t, recorder.Header().Get(http.CanonicalHeaderKey("Content-Type")), "text/plain")
	assert.NotEmpty(t, recorder.Body)
}

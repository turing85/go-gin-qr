package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"go-gin-qr/internal/config"
)

func TestMetricsRoute(t *testing.T) {
	// GIVEN
	engine := SetupEngine()
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, config.GetConfig().Metrics().Path(), nil)
	if err != nil {
		t.Fatalf("Failure during HTTP request creation: %s", err)
	}

	// WHEN
	engine.ServeHTTP(recorder, req)

	// THEN
	assert.Equal(t, 200, recorder.Code)
	assert.Contains(t, recorder.Header().Get(http.CanonicalHeaderKey("Content-Type")), "text/plain")
	assert.NotEmpty(t, recorder.Body)
}

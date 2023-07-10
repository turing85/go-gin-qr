package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"go-gin-qr/internal/appconfig"
)

func TestMetricsRoute(t *testing.T) {
	// GIVEN
	appConfig := appconfig.GetConfig()
	engine := SetupEngine(appConfig)
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, appConfig.Metrics().Path(), nil)
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

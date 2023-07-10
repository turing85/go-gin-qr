package appconfig

type Config interface {
	Health() WithPath
	Http() Http
	Metrics() WithPath
	Qr() WithPath

	// HealthPath to implement middleware.EngineConfig
	HealthPath() string
	// MetricsPath to implement middleware.EngineConfig
	MetricsPath() string
}

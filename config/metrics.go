package config

type metrics struct {
	Path_ string `yaml:"path" envconfig:"METRICS_PATH"`
}

func (m *metrics) Path() string {
	return m.Path_
}

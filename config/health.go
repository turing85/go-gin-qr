package config

type health struct {
	Path_ string `yaml:"path" envconfig:"HEALTH_PATH"`
}

func (h *health) Path() string {
	return h.Path_
}

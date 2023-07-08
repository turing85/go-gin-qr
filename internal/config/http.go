package config

type Http interface {
	Host() string
	Port() int
}

type http struct {
	Host_ string `yaml:"host" envconfig:"HTTP_HOST"`
	Port_ int    `yaml:"port" envconfig:"HTTP_PORT"`
}

func (h *http) Host() string {
	return h.Host_
}

func (h *http) Port() int {
	return h.Port_
}

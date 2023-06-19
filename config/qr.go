package config

type qr struct {
	Path_ string `yaml:"path" envconfig:"QR_PATH"`
}

func (q *qr) Path() string {
	return q.Path_
}

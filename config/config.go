package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

const configFileName = "config.yml"

type Config interface {
	Health() WithPath
	Http() Http
	Metrics() WithPath
	Qr() WithPath
}

type config struct {
	Health_  health  `yaml:"health"`
	Http_    http    `yaml:"http"`
	Metrics_ metrics `yaml:"metrics"`
	Qr_      qr      `yaml:"qr"`
}

func (c *config) Health() WithPath {
	return &c.Health_
}

func (c *config) Http() Http {
	return &c.Http_
}

func (c *config) Metrics() WithPath {
	return &c.Metrics_
}

func (c *config) Qr() WithPath {
	return &c.Qr_
}

func newDefaultConfig() *config {
	return &config{
		health{
			Path_: "/health",
		},
		http{
			Host_: "0.0.0.0",
			Port_: 8080,
		},
		metrics{
			Path_: "/metrics",
		},
		qr{
			Path_: "/qr-code",
		},
	}
}

var config_ *config = nil

func GetConfig() Config {
	if config_ == nil {
		config_ = newDefaultConfig().
			enhanceFromFile().
			enhanceFromEnv()
	}
	return config_
}

func (c *config) enhanceFromFile() *config {
	if fallback := getFallbackIfConfigFileAbsent(); fallback != nil {
		return fallback
	}
	return c.openAndReadConfigFile()
}

func getFallbackIfConfigFileAbsent() *config {
	if _, err := os.Stat(configFileName); err != nil {
		log.Info().Msgf(
			`Config file "%s" not found; using default configuration. Cause: %s`,
			configFileName,
			err)
		return newDefaultConfig()
	}
	return nil
}

func (c *config) openAndReadConfigFile() *config {
	file, err := os.Open(configFileName)
	defer func() {
		_ = file.Close()
	}()
	if err != nil {
		log.Warn().Msgf(
			`Error opening config_ file "%s"; using default configuration. Cause: %s`,
			configFileName,
			err)
		return newDefaultConfig()
	}
	return c.parseYaml(file)
}

func (c *config) parseYaml(file *os.File) *config {
	enhanced := *c
	if err := yaml.NewDecoder(file).Decode(&enhanced); err != nil {
		log.Warn().Msgf(
			`Error reading config_ file "%s"; using default configuration. Cause: %s`,
			configFileName,
			err)
		return c
	}
	return &enhanced
}

func (c *config) enhanceFromEnv() *config {
	enhanced := *c
	if err := envconfig.Process("", &enhanced); err != nil {
		log.Warn().Msgf(`Error reading environment; skipping. Cause: %s`, err)
		return c
	}
	return &enhanced
}

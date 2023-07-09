package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

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

func (c *config) enhanceFromFile(configFile string) *config {
	if err := configFileIsReadable(configFile); err != nil {
		log.Info().Msgf(
			`Config file "%s" not readable; using default configuration. Cause: %s`,
			configFile,
			err)
		return newDefaultConfig()
	}
	return c.openAndReadConfigFile(configFile)
}

func configFileIsReadable(configFile string) error {
	if _, err := os.Stat(configFile); err != nil {
		return err
	}
	return nil
}

func (c *config) openAndReadConfigFile(configFile string) *config {
	file, err := os.Open(configFile)
	defer func() {
		if err = file.Close(); err != nil {
			log.Warn().Msgf(`Unable to close file "%s". Cause: %s`, configFile, err)
		}
	}()
	if err != nil {
		log.Warn().Msgf(
			`Error opening config file "%s"; using default configuration. Cause: %s`,
			configFile,
			err)
		return newDefaultConfig()
	}
	return c.parseYaml(file)
}

func (c *config) parseYaml(file *os.File) *config {
	enhanced := *c
	if err := yaml.NewDecoder(file).Decode(&enhanced); err != nil {
		log.Warn().Msgf(
			`Error reading config file "%s"; using default configuration. Cause: %s`,
			defaultConfigFile,
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

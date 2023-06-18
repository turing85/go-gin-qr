package config

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Http    http    `yaml:"http"`
	Metrics metrics `yaml:"metrics"`
	Qr      qr      `yaml:"qr"`
}

type http struct {
	Host string `yaml:"host" envconfig:"HTTP_HOST"`
	Port int    `yaml:"port" envconfig:"HTTP_PORT"`
}

type metrics struct {
	Path string `yaml:"path" envconfig:"METRICS_PATH"`
}

type qr struct {
	Path string `yaml:"path" envconfig:"QR_PATH"`
}

const configFileName = "config.yml"

var emptyConfig = Config{}

var config = emptyConfig

var defaultConfig = Config{
	http{
		Host: "0.0.0.0",
		Port: 8080,
	},
	metrics{
		Path: "/metrics",
	},
	qr{
		Path: "/qr-code",
	},
}

func GetConfig() Config {
	if config == emptyConfig {
		config = enhanceConfigFromEnv(
			readConfigFromFile())
	}
	return config
}

func readConfigFromFile() Config {
	if _, err := os.Stat(configFileName); err != nil {
		log.Info().Msg(fmt.Sprintf(`Config file "%s" not found; using default configuration`, configFileName))
		return defaultConfig
	}
	file, err := os.Open(configFileName)
	defer func() {
		_ = file.Close()
	}()
	if err != nil {
		// do something
	}
	config := defaultConfig
	if err = yaml.NewDecoder(file).Decode(&config); err != nil {
		log.Warn().Msg(fmt.Sprintf(`Error reading config file "%s"; using default configuration"`, configFileName))
		return defaultConfig
	}
	return config
}

func enhanceConfigFromEnv(config Config) Config {
	copyConfig := config
	if err := envconfig.Process("", &copyConfig); err != nil {
		log.Warn().Msg(fmt.Sprintf(`Error reading environment; skipping: %s`, err))
	} else {
		config = copyConfig
	}
	return config
}

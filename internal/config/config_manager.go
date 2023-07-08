package config

import (
	"os"
	"sync"
)

const defaultConfigFile = "config.yml"

var instance *config = nil
var configMutex sync.Mutex

func GetConfig() Config {
	if instance == nil {
		configMutex.Lock()
		defer configMutex.Unlock()
		if instance == nil {
			instance = newDefaultConfig().
				enhanceFromFile(getConfigFile()).
				enhanceFromEnv()
		}
	}
	return instance
}

func getConfigFile() string {
	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		configFile = defaultConfigFile
	}
	return configFile
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

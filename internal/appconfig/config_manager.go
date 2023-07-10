package appconfig

import (
	"os"
	"sync"
)

const defaultConfigFile = "config.yml"

var instance = struct {
	config Config
	mutex  sync.Mutex
}{}

func GetConfig() Config {
	if instance.config == nil {
		instance.mutex.Lock()
		defer instance.mutex.Unlock()
		if instance.config == nil {
			instance.config = newDefaultConfig().
				enhanceFromFile(getConfigFile()).
				enhanceFromEnv()
		}
	}
	return instance.config
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

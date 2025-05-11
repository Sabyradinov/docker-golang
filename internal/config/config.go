package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configs struct {
	StopTimeoutMS int
	WebServer     WebServerConfig
}

type WebServerConfig struct {
	Port int
	GIN  GINConfig
}

type GINConfig struct {
	ReleaseMode bool
	UseLogger   bool
	UseRecovery bool
}

func Init(path string) (config *Configs, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("config file reading error, err : %v", err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("config file unmarshalling error, err : %v", err)
	}

	return
}

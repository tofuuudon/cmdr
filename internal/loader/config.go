package loader

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	Directives []string `json:"directives"`
}

func GetConfig() Config {
	configFile, _ := os.Open("./config.json")
	configBytes, _ := io.ReadAll(configFile)

	var config Config
	json.Unmarshal(configBytes, &config)

	return config
}

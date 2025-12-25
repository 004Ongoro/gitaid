package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	GeminiKey string `json:"gemini_key"`
	Model     string `json:"model"`
}

func LoadConfig() (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("could not find home directory: %v", err)
	}

	configPath := filepath.Join(home, ".config", "gitaid", "config.json")
	
	file, err := os.Open(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("config file not found at %s. Please create it with your gemini_key", configPath)
		}
		return nil, err
	}
	defer file.Close()

	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %v", err)
	}

	if cfg.Model == "" {
		cfg.Model = "gemini-1.5-flash"
	}

	return &cfg, nil
}
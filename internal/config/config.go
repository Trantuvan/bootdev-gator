package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName string = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	var config Config
	configPath, err := getConfigFilePath()

	if err != nil {
		return config, fmt.Errorf("get config path: %w", err)
	}

	file, err := os.Open(configPath)

	if err != nil {
		return config, fmt.Errorf("open file: %w", err)
	}

	defer file.Close()
	//* create decoder from place to read
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&config); err != nil {
		return config, fmt.Errorf("decode config: %w", err)
	}

	return config, nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	return write(*cfg)
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", fmt.Errorf("cannot find $HOME enviroment variable: %w", err)
	}

	configPath := filepath.Join(homeDir, configFileName)
	return configPath, nil
}

func write(cfg Config) error {
	configPath, err := getConfigFilePath()

	if err != nil {
		return fmt.Errorf("get path config: %w", err)
	}

	//* if file exist overwrite; create new file
	file, err := os.Create(configPath)

	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}

	//* create encoder from place to write
	encoder := json.NewEncoder(file)

	if err := encoder.Encode(cfg); err != nil {
		return fmt.Errorf("encode config: %w", err)
	}

	return nil
}

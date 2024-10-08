package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	permReadWrite         = 0644
	configFileName string = ".gatorconfig.json"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() Config {
	var config Config
	configPath, err := getConfigFilePath()

	if err != nil {
		log.Fatalf("%s", err)
		return config
	}

	file, err := os.ReadFile(configPath)

	if err != nil {
		log.Fatalf("cannot open file %s", err)
		return config
	}

	decoder := json.NewDecoder(bytes.NewBuffer(file))

	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("error parsing json into struct %s", err)
		return config
	}

	return config
}

func (cfg *Config) SetUser(username string) {
	cfg.CurrentUserName = username
	write(*cfg)
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", fmt.Errorf("error cannot find $HOME enviroment variable %w", err)
	}

	configPath := filepath.Join(homeDir, configFileName)
	return configPath, nil
}

func write(cfg Config) error {
	jsonData, err := json.Marshal(cfg)

	if err != nil {
		return fmt.Errorf("error cannot parse struct into json %w", err)
	}

	configPath, err := getConfigFilePath()

	if err != nil {
		return err
	}

	if err := os.WriteFile(configPath, jsonData, permReadWrite); err != nil {
		return fmt.Errorf("error cannot write to disk %w", err)
	}

	return nil
}

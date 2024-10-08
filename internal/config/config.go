package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const configFileName string = ".gatorconfig.json"

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

	file, err := os.Open(configPath)

	if err != nil {
		log.Fatalf("cannot open file %s", err)
		return config
	}

	defer file.Close()
	//* create decoder from place to read
	decoder := json.NewDecoder(file)

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
	configPath, err := getConfigFilePath()

	if err != nil {
		return err
	}

	//* if file exist overwrite; create new file
	file, err := os.Create(configPath)

	if err != nil {
		return fmt.Errorf("path error %w", err)
	}

	//* create encoder from place to write
	encoder := json.NewEncoder(file)

	if err := encoder.Encode(cfg); err != nil {
		return fmt.Errorf("error cannot parse struct into json %w", err)
	}

	return nil
}

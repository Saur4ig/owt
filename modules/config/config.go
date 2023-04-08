package config

import (
	"fmt"
	"os"
)

const (
	// names of required envs
	DB_NAME = "DBNAME"
	PORT    = "PORT"
)

type Config struct {
	ServerPort string
	Database   string
}

func GetConfig() (*Config, error) {
	err := checkAllEnvVariables()
	if err != nil {
		return nil, err
	}

	return &Config{
		ServerPort: os.Getenv(PORT),
		Database:   os.Getenv(DB_NAME),
	}, nil
}

// if at least one env params missing - error
func checkAllEnvVariables() error {
	if _, ok := os.LookupEnv(DB_NAME); !ok {
		return fmt.Errorf("%s is missing", DB_NAME)
	}

	if _, ok := os.LookupEnv(PORT); !ok {
		return fmt.Errorf("%s is missing", PORT)
	}

	return nil
}

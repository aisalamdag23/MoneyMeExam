package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v3"
	"log"
)

const configPathEnvName = "SPEC_FILE"

type (
	specWithMetaConfig struct {
		Spec Config `yaml:"spec"`
	}

	Config struct {
		// DB ...
		DB Database `yaml:"database" validate:"required"`
	}

	// Database config
	Database struct {
		// Host ...
		Host string `yaml:"host" validate:"required"`
		// Port ...
		Port string `yaml:"port" validate:"required"`
		// DBName ...
		DBName string `yaml:"name" validate:"required"`
		// User ...
		User string `yaml:"user" validate:"required"`
		// Pass ...
		Pass string `yaml:"pass" validate:"required"`
		// Conns ...
		Conns string `yaml:"conns"`
	}
)

// Load loads all configurations in to a new Config struct
func Load() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	configFilePath := os.Getenv(configPathEnvName)
	if configFilePath == "" {
		return nil, fmt.Errorf("env variable %s is not defined", configPathEnvName)
	}
	// reading app file config
	configFile, err := os.Open(configFilePath) // nolint:gosec
	if err != nil {
		return nil, errors.Wrap(err, "can not open config file")
	}

	var spec specWithMetaConfig
	err = yaml.NewDecoder(configFile).Decode(&spec)
	if err != nil {
		return nil, errors.Wrap(err, "can not unmarshal config data")
	}

	config := &spec.Spec

	// validating app file configs
	v := validator.New()
	err = v.Struct(config)
	if err != nil {
		return nil, errors.Wrap(err, "config file is not valid")
	}
	return config, nil
}

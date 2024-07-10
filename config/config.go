package config

import (
	"encoding/json"
	"os"

	"github.com/go-playground/validator"
)

const (
	path = "config/config.json"
)

type Config struct {
	HTTP struct {
		Address string
	}
	GRPC struct {
		Address string
	}
	Path string
}

func Load() (cfg *Config, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	if err = json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}

	if err = validator.New().Struct(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

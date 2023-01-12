package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTPServer `yaml:"http_server"`
		Redis      `yaml:"redis"`
	}

	HTTPServer struct {
		BaseURL string `yaml:"base_url" env:"HTTP_BASE_URL" env-default:""`
		Port    string `yaml:"port" env:"HTTP_PORT" env-default:""`
	}

	Redis struct {
		Enabled  bool   `yaml:"enabled" env:"REDIS_ENABLED" env-default:"false"`
		Address  string `yaml:"address" env:"REDIS_ADDRESS" env-default:""`
		Password string `yaml:"password" env:"REDIS_PASSWORD" env-default:""`
	}
)

func NewConfig(filename string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(filename, cfg)
	if err != nil {
		return nil, fmt.Errorf("config file error: %w", err)
	}

	return cfg, nil
}

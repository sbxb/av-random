package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTPServer `yaml:"http_server"`
		Redis      `yaml:"redis"`
		MongoDB    `yaml:"mongodb"`
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

	MongoDB struct {
		Enabled  bool   `yaml:"enabled" env:"MONGODB_ENABLED" env-default:"false"`
		Address  string `yaml:"address" env:"MONGODB_ADDRESS" env-default:""`
		User     string `yaml:"user" env:"MONGODB_USER" env-default:""`
		Password string `yaml:"password" env:"MONGODB_PASSWORD" env-default:""`
		DBName   string `yaml:"dbname" env:"MONGODB_DBNAME" env-default:""`
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

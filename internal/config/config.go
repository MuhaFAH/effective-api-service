package config

import (
	"github.com/MuhaFAH/effective-api-service/internal/storage/gorm"
)

type Config struct {
	Address  string `yaml:"address" env:"ADDRESS" env-default:"127.0.0.1:8080"`
	LogLevel string `yaml:"log_level" env:"LOG_LEVEL" env-default:"info"`

	gorm.PostgresConfig
}

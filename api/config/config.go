package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"gorm.io/gorm/logger"
)

var (
	App Config
)

type Config struct {
	Database        DatabaseConfig
	Port            int    `env:"PORT" env-default:"80"`
	LogLevel        string `env:"LOG_LEVEL" env-default:"debug"`
	SanctionFileURL string `env:"SANCTION_FILE_URL" env-default:"https://sigmaratings-public-static.s3.amazonaws.com/eu_sanctions.csv"`
}

type DatabaseConfig struct {
	Host     string          `env:"POSTGRES_HOST" env-default:"db"`
	User     string          `env:"POSTGRES_USER" env-default:"sanctions"`
	Password string          `env:"POSTGRES_PASSWORD" env-default:"are-fun"`
	Port     int             `env:"POSTGRES_PORT" env-default:"5432"`
	Name     string          `env:"POSTGRES_DATABASE" env-default:"sanctions"`
	LogLevel logger.LogLevel `env:"POSTGRES_LOG_LEVEL" env-default:"1"`
}

func init() {
	if err := cleanenv.ReadEnv(&App); err != nil {
		fmt.Printf("Could not load config: %v\n", err)
		os.Exit(-1)
	}
}

func Database() *DatabaseConfig {
	return &App.Database
}

func TestDatabase() *DatabaseConfig {
	*&App.Database = DatabaseConfig{
		Host:     "test-db",
		Port:     5433,
		User:     "sanctions",
		Password: "are-fun",
		Name:     "sanctions",
		LogLevel: 1,
	}

	return &App.Database
}

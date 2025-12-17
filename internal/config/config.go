// Package config provides utilities for config load
package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

// Config is a app config model
type Config struct {
	JWTSecret           string        `env:"JWT_SECRET"`
	AccessTokenExpires  time.Duration `env:"ACCESS_TOKEN_EXPIRES" envDefault:"15m"`
	RefreshTokenExpires time.Duration `env:"REFRESH_TOKEN_EXPIRES" envDefault:"72h"`
	ServerPort          int           `env:"PORT" envDefault:"8080"`
	AppEnv              string        `env:"APP_ENV" envDefault:"prod"`
	Timeout             time.Duration `env:"TIMEOUT" envDefault:"4s"`
	IdleTimeout         time.Duration `env:"IDLE_TIMEOUT" envDefault:"60s"`
	LogPath             string        `env:"LOG_PATH" envDefault:"./app.log"`
	ShutdownTimeout     time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"10s"`
	SwaggerMode         bool          `env:"SWAGGER_MODE" envDefault:"false"`
	DBHost              string        `env:"DB_HOST" envDefault:"postgres"`
	DBUser              string        `env:"DB_USER" envDefault:"tapok"`
	DBName              string        `env:"DB_NAME" envDefault:"drive"`
	DBPassword          string        `env:"DB_PASSWORD" envDefault:"password"`
}

// LoadConfig loads app configuration
func LoadConfig() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return cfg, fmt.Errorf("failed to load config: %w", err)
	}

	err := cfg.validate()
	if err != nil {
		return Config{}, fmt.Errorf("failed to validate config: %w", err)
	}

	return cfg, nil
}

func (c *Config) validate() error {
	if len(c.JWTSecret) < 32 {
		return fmt.Errorf("JWT secret must be at least 32 bytes long, got %d", len(c.JWTSecret))
	}

	if c.ServerPort < 0 || c.ServerPort > 65535 {
		return fmt.Errorf("invalid port: %d", c.ServerPort)
	}

	if c.AppEnv != "dev" && c.AppEnv != "prod" {
		return fmt.Errorf("app env must be 'dev' or 'prod', got %v", c.AppEnv)
	}

	if c.Timeout <= 0 {
		return fmt.Errorf("timeout must be positive, got %v", c.Timeout)
	}

	if c.IdleTimeout <= 0 {
		return fmt.Errorf("idle timeout must be positive, got %v", c.IdleTimeout)
	}

	return nil
}

// PostgresDSN create a db conn string
func (c *Config) PostgresDSN() string {
	return fmt.Sprintf(
		"host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost,
		c.DBUser,
		c.DBPassword,
		c.DBName,
	)
}

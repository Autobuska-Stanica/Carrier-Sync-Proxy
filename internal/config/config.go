package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

type Environment string

const (
	EnvLocal Environment = "local"
	EnvDev   Environment = "dev"
)

type Config struct {
	Server   Server
	AppName  string      `env:"APP_NAME" envDefault:"carrier-sync-proxy"`
	Env      Environment `env:"ENV" envDefault:"dev"`
	LogLevel int         `env:"LOG_LEVEL" envDefault:"0"`
}

type Server struct {
	Port            int           `env:"SERVER_PORT" envDefault:"8080"`
	Host            string        `env:"SERVER_HOST" envDefault:"0.0.0.0"`
	ReadTimeout     time.Duration `env:"SERVER_READ_TIMEOUT" envDefault:"5s"`
	ShutdownTimeout time.Duration `env:"SERVER_SHUTDOWN_TIMEOUT" envDefault:"10s"`
}

func New() (*Config, error) {
	cfg := &Config{}

	opts := env.Options{}
	if err := env.ParseWithOptions(cfg, opts); err != nil {
		return nil, fmt.Errorf("failed to parse env config: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return cfg, nil
}

func (c *Config) Validate() error {
	// Env
	switch c.Env {
	case EnvLocal, EnvDev:

	default:
		return fmt.Errorf("unknown env value %s", c.Env)
	}

	// Server
	if c.Server.Port < 1 || c.Server.Port > 65535 {
		return fmt.Errorf("server port must be between 1 and 65535, got %d", c.Server.Port)
	}
	return nil
}

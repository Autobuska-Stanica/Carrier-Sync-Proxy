package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Autobuska-Stanica/Carrier-Sync-Proxy/internal/config"
	"github.com/Autobuska-Stanica/Carrier-Sync-Proxy/internal/logger"
	"github.com/Autobuska-Stanica/Carrier-Sync-Proxy/internal/server"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
	}
}

func run() error {
	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("failed to create config: %w", err)
	}

	log := logger.New(cfg)
	slog.SetDefault(log)

	srv := server.New(cfg)
	return srv.Run()
}

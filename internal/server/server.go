package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Autobuska-Stanica/Carrier-Sync-Proxy/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func New(cfg *config.Config) *Server {
	s := &Server{}
	s.httpServer = &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      s.setupRoutes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	return s
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

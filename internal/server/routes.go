package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) setupRoutes() http.Handler {
	r := chi.NewRouter()

	// 1. Global Middleware (Applies to ALL routes)
	r.Use(middleware.RequestID)

	// 2. Unversioned Root / Health Endpoints
	r.Get("/alive", alive)
	r.Get("/ready", ready)

	// // 3. Initialize Handlers
	// v1Handler := v1.NewHandler()
	// v2Handler := v2.NewHandler()
	// adminHandler := admin.NewHandler()

	// // 4. Mount Sub-Routers
	// // Note: trailing slashes are handled cleanly by chi when mounting
	// r.Mount("/v1", v1.NewRouter(v1Handler))
	// r.Mount("/v2", v2.NewRouter(v2Handler))
	// r.Mount("/admin", admin.NewRouter(adminHandler, customMiddleware.AdminAuth))

	return r
}

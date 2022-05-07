package api

import (
	"context"
	"net/http"

	coreMiddleware "github.com/facily-tech/go-core/http/server/middleware"

	"github.com/LeandroAlcantara-1997/Beauty_salon/appointment/internal/container"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Handler(ctx context.Context, dep *container.Dependency) http.Handler {
	r := chi.NewMux()

	r.Use(dep.Components.Tracer.Middleware)             // must be first
	r.Use(middleware.RequestID)                         // must be second
	r.Use(coreMiddleware.Logger(dep.Components.Log))    // must be third
	r.Use(coreMiddleware.Recoverer(dep.Components.Log)) // must be forty

	r.Handle("/metrics", promhttp.Handler())
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {})

	return r
}

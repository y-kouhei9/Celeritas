package celeritas

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (c *Celeritas) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	if c.Debug {
		mux.Use(middleware.Logger)
	}
	mux.Use(middleware.Recoverer)

	return mux
}

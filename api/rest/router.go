package rest

import (
	"net/http"

	"github.com/sbxb/av-random/config"

	"github.com/go-chi/chi/v5"
)

func NewRouter(cfg config.HTTPServer /*, rs *random.Service*/) http.Handler {
	router := chi.NewRouter()
	rh := NewRouteHandler(cfg /*, rs*/)

	router.Route(cfg.BaseURL, func(r chi.Router) {
		r.Post("/generate", rh.PostGenerate)
		r.Get("/retrieve", rh.GetRetrieve)
	})

	// TODO health checking route

	return router
}

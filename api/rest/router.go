package rest

import (
	"net/http"

	"github.com/sbxb/av-random/config"
	"github.com/sbxb/av-random/service/random"

	"github.com/go-chi/chi/v5"
)

func NewRouter(cfg config.HTTPServer, rs *random.Service) http.Handler {
	router := chi.NewRouter()
	rh := NewRouteHandler(cfg, rs)

	router.Route(cfg.BaseURL, func(r chi.Router) {
		r.Post("/generate", rh.PostGenerate)
		r.Get("/retrieve/{id}", rh.GetRetrieve)
	})

	router.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})

	return router
}

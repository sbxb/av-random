package rest

import (
	"net/http"

	"github.com/sbxb/av-random/config"
	"github.com/sbxb/av-random/service/random"
)

type RouteHandler struct {
	config config.HTTPServer
	rs     *random.Service
}

func NewRouteHandler(cfg config.HTTPServer, rs *random.Service) RouteHandler {
	return RouteHandler{
		config: cfg,
		rs:     rs,
	}
}

func (rh RouteHandler) PostGenerate(w http.ResponseWriter, r *http.Request) {
	//
}

func (rh RouteHandler) GetRetrieve(w http.ResponseWriter, r *http.Request) {
	//
}

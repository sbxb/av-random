package rest

import (
	"log"
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

// e.g. curl -i -X POST http://localhost:8080/api/generate
func (rh RouteHandler) PostGenerate(w http.ResponseWriter, r *http.Request) {
	log.Println("PostGenerate handler hit")
	id, errID := rh.rs.GenerateID()
	value, errValue := rh.rs.GenerateRandomValue()

	if errID != nil || errValue != nil {
		http.Error(w, "error", http.StatusInternalServerError)
	}

	log.Printf("id = %s; value = %d\n", id, value)
	w.WriteHeader(http.StatusCreated)
}

func (rh RouteHandler) GetRetrieve(w http.ResponseWriter, r *http.Request) {
	log.Println("GetRetrieve handler hit")
	//
}

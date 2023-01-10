package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
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
	// FIXME set content-type in middleware
	log.Println("PostGenerate handler hit")

	id, errID := rh.rs.GenerateID()
	if errID != nil {
		http.Error(w, errID.Error(), http.StatusInternalServerError)
		return
	}

	value, errValue := rh.rs.GenerateRandomValue()
	if errValue != nil {
		http.Error(w, errValue.Error(), http.StatusInternalServerError)
		return
	}

	err := rh.rs.SaveRandomValue(id, value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("id = %s; value = %d\n", id, value)

	response := GenerateResponse{ID: id}
	rbytes, errJSON := json.Marshal(response)
	if errJSON != nil {
		http.Error(w, errJSON.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(rbytes)
}

// e.g. curl -i -X GET http://localhost:8080/api/retrieve/<id>
func (rh RouteHandler) GetRetrieve(w http.ResponseWriter, r *http.Request) {
	log.Println("GetRetrieve handler hit")

	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "no generation id provided", http.StatusBadRequest)
		return
	}

	re, err := rh.rs.RetrieveRandomValue(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	rbytes, errJSON := json.Marshal(convRandomEntityToRetrieveResponse(re))
	if errJSON != nil {
		http.Error(w, errJSON.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(rbytes)
}

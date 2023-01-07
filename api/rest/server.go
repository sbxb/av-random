package rest

import (
	"net/http"
	"time"

	"github.com/sbxb/av-random/config"
)

type HTTPServer struct {
	Server *http.Server
}

func NewHTTPServer(cfg config.HTTPServer, router http.Handler) *HTTPServer {
	server := &http.Server{
		Addr:         cfg.Port,
		Handler:      router,
		ReadTimeout:  8 * time.Second,
		WriteTimeout: 8 * time.Second,
		IdleTimeout:  36 * time.Second,
	}

	return &HTTPServer{
		Server: server,
	}
}

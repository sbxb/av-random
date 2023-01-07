package main

import (
	"flag"
	"log"

	"github.com/sbxb/av-random/api/rest"
	"github.com/sbxb/av-random/config"
)

func main() {
	log.SetPrefix("av-random: ")

	configFile := flag.String("c", "./config.yaml", "read config file")
	flag.Parse()

	cfg, err := config.NewConfig(*configFile)
	if err != nil {
		log.Fatalf("cannot read or parse config file: %v", err)
	}

	log.Printf("Config read: %+v", cfg)

	router := rest.NewRouter(cfg.HTTPServer /*, rs*/)

	server := rest.NewHTTPServer(cfg.HTTPServer, router)

	log.Println("HTTP Server starts")

	if err := server.Server.ListenAndServe(); err != nil {
		log.Printf("server failed to start with error: %v", err)
	}
}

package main

import (
	"flag"
	"log"

	"github.com/sbxb/av-random/api/rest"
	"github.com/sbxb/av-random/config"
	"github.com/sbxb/av-random/service/random"
	"github.com/sbxb/av-random/storage/inmemory"
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

	storage, err := inmemory.NewMemoryStorage()
	if err != nil {
		log.Fatalf("cannot create storage: %v", err)
	}

	randomService, err := random.New(storage)
	if err != nil {
		log.Fatalf("cannot create Random.Service: %v", err)
	}

	log.Println("Random.Service created")

	router := rest.NewRouter(cfg.HTTPServer, randomService)

	server := rest.NewHTTPServer(cfg.HTTPServer, router)

	log.Println("HTTP Server starts")

	if err := server.Server.ListenAndServe(); err != nil {
		log.Printf("server failed to start with error: %v", err)
	}
}

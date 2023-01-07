package main

import (
	"flag"
	"log"

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

}

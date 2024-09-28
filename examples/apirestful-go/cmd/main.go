package main

import (
	"apirestful-go/cmd/server"
	"apirestful-go/config"
	"log"
)

func main() {
	log.Println("Loading configuration...")
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	log.Println("Starting server...")
	srv := server.NewServerChi(cfg)
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}

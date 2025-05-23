package main

import (
	"auth-service/internal/config"
	"auth-service/internal/server"
	"log"
)

func main() {

	cfg := config.New()
	srv := server.New(*cfg)
	log.Printf("starting auth-server at: %s", cfg.ServerConfig.WebPort)
	err := srv.ListenAndServe()
	if err != nil {
		log.Printf("unable to start auth-server. Error: %s", err.Error())
		return

	}

}

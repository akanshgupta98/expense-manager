package main

import (
	"api-gateway/internal/config"
	"api-gateway/internal/server"

	"github.com/akanshgupta98/go-logger/v2"
)

func main() {
	cfg := config.New()
	err := logger.Init(cfg.LoggerCfg)
	if err != nil {
		panic(err)
	}
	srv := server.New(cfg)
	logger.Infof("starting gateway at: %s", cfg.ServerCfg.WebPort)
	err = srv.ListenAndServe()
	if err != nil {
		logger.Errorf("unable to start gateway server: %s", err.Error())
	}

}

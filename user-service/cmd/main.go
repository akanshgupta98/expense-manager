package main

import (
	"user-service/internal/config"
	"user-service/internal/server"

	"github.com/akanshgupta98/go-logger/v2"
)

func main() {
	cfg := config.New()
	err := logger.Init(cfg.LoggerCfg)
	if err != nil {
		panic(err)
	}
	srv := server.New(cfg)
	logger.Infof("starting user-service at port :%s", cfg.ServerCfg.WebPort)
	err = srv.ListenAndServe()
	if err != nil {
		logger.Errorf("unable to start user-service %s", err.Error())
		return
	}

}

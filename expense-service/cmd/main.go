package main

import (
	"expense-service/internal/config"
	"expense-service/internal/server"

	"github.com/akanshgupta98/go-logger/v2"
)

func main() {

	cfg := config.New()

	err := logger.Init(cfg.LoggerCfg)
	if err != nil {
		panic(err)
	}
	srv := server.New(cfg)
	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

package main

import (
	"auth-service/internal/amqp"
	"auth-service/internal/config"
	"auth-service/internal/database"
	"auth-service/internal/server"
	"auth-service/internal/service"

	"github.com/akanshgupta98/go-logger"
)

func main() {

	cfg := config.New()
	err := logger.Init(cfg.LoggerConfig)
	if err != nil {
		panic(err)
	}
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	service.Initialize(db, cfg)

	err = amqp.Connect(cfg)
	if err != nil {
		panic(err)
	}
	defer amqp.CloseConnection()

	srv := server.New(cfg)
	logger.Infof("starting auth-server on port: %s", cfg.ServerConfig.WebPort)
	err = srv.ListenAndServe()
	if err != nil {
		logger.Errorf("unable to start auth-server. Error: %s", err.Error())
		return

	}

}

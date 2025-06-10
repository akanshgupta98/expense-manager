package main

import (
	"user-service/internal/amqp"
	"user-service/internal/config"
	"user-service/internal/database"
	"user-service/internal/server"
	"user-service/internal/service"

	"github.com/akanshgupta98/go-logger/v2"
)

func main() {
	cfg := config.New()
	err := logger.Init(cfg.LoggerCfg)
	if err != nil {
		panic(err)
	}

	db, err := database.ConnectToDB(cfg.DBConfig.DSN)
	if err != nil {
		panic(err)
	}

	err = amqp.Connect(cfg)
	if err != nil {
		panic(err)
	}
	consumer := map[amqp.EXCHANGE]amqp.TOPICS{}
	for _, ex := range cfg.AMQPConfig.ConsumeExchanges {
		consumer[ex] = []string{"user.*"}
	}

	go amqp.ConsumeEvents(consumer)

	service.Initialize(db)
	srv := server.New(cfg)
	logger.Infof("starting user-service at port :%s", cfg.ServerCfg.WebPort)
	err = srv.ListenAndServe()
	if err != nil {
		logger.Errorf("unable to start user-service %s", err.Error())
		return
	}

}

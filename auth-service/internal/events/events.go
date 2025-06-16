package events

import (
	"auth-service/internal/config"
	"time"

	"github.com/akanshgupta98/go-logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

func New(conn *amqp.Connection, cfg config.Config) *Event {
	event := Event{
		URL:          cfg.AMQPConfig.Url,
		ExchangeName: cfg.AMQPConfig.PublishExhange,
		Conn:         conn,
	}
	return &event
}

func Connect(url string) (*amqp.Connection, error) {
	var conn *amqp.Connection
	var err error
	totalRetry := 5
	retryCount := 0
	for i := 0; i < totalRetry; i++ {
		conn, err = openConnection(url)
		if err != nil {
			logger.Infof("rabbitMQ not ready. Backing off...")
			retryCount++
			backOff := time.Duration(2 * retryCount)
			time.Sleep(backOff)

		} else {
			break
		}
	}

	return conn, nil

}

func (e *Event) CloseConnection() {
	e.Conn.Close()
	e.Channel.Close()
}

func openConnection(url string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	return conn, nil

}

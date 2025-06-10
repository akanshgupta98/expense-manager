package amqp

import (
	"auth-service/internal/config"
	"encoding/json"

	"github.com/rabbitmq/amqp091-go"
	mq "github.com/rabbitmq/amqp091-go"
)

func openConnection(url string) (*mq.Connection, error) {
	conn, err := mq.Dial(url)
	if err != nil {
		return nil, err
	}
	return conn, nil

}

func openChannel(conn *mq.Connection) (*mq.Channel, error) {

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil

}

func CloseConnection() {
	amqp.Conn.Close()
	amqp.Channel.Close()
}

func Connect(cfg config.Config) error {

	var err error

	amqp.Conn, err = openConnection(cfg.AMQPConfig.Url)
	if err != nil {
		return err
	}
	amqp.Channel, err = openChannel(amqp.Conn)
	if err != nil {
		return err
	}
	err = DeclareExchange(cfg.AMQPConfig.PublishExhange)
	if err != nil {
		return err
	}

	return nil
}

func DeclareExchange(name string) error {

	err := amqp.Channel.ExchangeDeclare(
		name,
		"topic",
		false,
		false,
		false,
		false,
		nil,
	)
	amqp.ExchangeName = name

	return err

}
func PublishMessage(exchangeName, routingKey string, data []byte) error {
	err := DeclareExchange(exchangeName)
	if err != nil {
		return err
	}

	err = amqp.Channel.Publish(
		exchangeName,
		routingKey,
		false,
		false,
		amqp091.Publishing{
			ContentType:     "application/json",
			ContentEncoding: "json",
			Body:            data,
		})

	if err != nil {
		return err
	}
	return nil

}

func PublishUserCreated(data UserCreatedEvent) error {
	msg, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = PublishMessage(amqp.ExchangeName, "user.created", msg)
	if err != nil {
		return err
	}
	return nil

}

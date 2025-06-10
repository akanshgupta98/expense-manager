package amqp

import "github.com/rabbitmq/amqp091-go"

var amqp AMQP

type AMQP struct {
	Conn         *amqp091.Connection
	Channel      *amqp091.Channel
	ExchangeName string
}

type UserCreatedEvent struct {
	UserID    int    `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Country   string `json:"country"`
	Email     string `json:"email"`
}

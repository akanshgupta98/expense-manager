package events

import (
	"github.com/akanshgupta98/expense-manager/contracts/eventspb"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Event struct {
	Conn         *amqp.Connection
	Channel      *amqp.Channel
	ExchangeName string
	URL          string
}

type EventData struct {
	Topic string
	Data  *eventspb.UserCreatedEvent
}

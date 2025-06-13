package events

import (
	"github.com/akanshgupta98/go-logger"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

func (e *Event) openChannel() (*amqp.Channel, error) {

	ch, err := e.Conn.Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil

}

func (e *Event) declarePublisherExchange() error {

	err := e.Channel.ExchangeDeclare(
		e.ExchangeName,
		"topic",
		false,
		false,
		false,
		false,
		nil,
	)

	return err

}
func (e *Event) SendEvent(payload EventData) error {
	if e.Channel == nil {
		logger.Debugf("opening a new channel")
		channel, err := e.openChannel()
		if err != nil {
			return err
		}
		e.Channel = channel

		err = e.declarePublisherExchange()
		if err != nil {
			return err
		}

	}

	serializedData, err := proto.Marshal(payload.Data)
	if err != nil {
		return err
	}

	err = e.Channel.Publish(
		e.ExchangeName,
		payload.Topic,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/x-protobuf",
			Body:        serializedData,
		})

	if err != nil {
		return err
	}

	return nil

}

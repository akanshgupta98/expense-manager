package amqp

import (
	"sync"
	"user-service/internal/config"

	"github.com/akanshgupta98/go-logger/v2"
	"github.com/rabbitmq/amqp091-go"
)

type AMQP struct {
	Conn    *amqp091.Connection
	Channel *amqp091.Channel
	Queues  map[string]amqp091.Queue
}

var amqp AMQP

type EXCHANGE = string
type TOPICS = []string

func Connect(cfg config.Config) error {

	conn, err := amqp091.Dial(cfg.AMQPConfig.URL)
	if err != nil {
		return err
	}
	amqp.Conn = conn

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	amqp.Channel = ch

	for _, eNames := range cfg.AMQPConfig.ConsumeExchanges {
		err := ch.ExchangeDeclare(eNames,
			"topic",
			false,
			false,
			false,
			false,
			nil)

		if err != nil {
			return err
		}
	}

	return nil

}

func DeclareRandomQueue() (amqp091.Queue, error) {
	// var err error
	q, err := amqp.Channel.QueueDeclare(
		"",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		return q, err
	}
	return q, nil
}
func ConsumeEvents(consume map[EXCHANGE]TOPICS) error {
	var wg sync.WaitGroup
	for ex, topics := range consume {
		q, err := DeclareRandomQueue()
		if err != nil {
			return err
		}

		for _, topic := range topics {
			amqp.Channel.QueueBind(
				q.Name,
				topic,
				ex,
				false,
				nil,
			)

		}

		msgs, err := amqp.Channel.Consume(q.Name, "", false, false, false, false, nil)
		if err != nil {
			return err
		}
		wg.Add(1)
		go func() {
			logger.Infof("started consumer")
			for d := range msgs {
				logger.Infof("recived event for topic: %s with data: %s", d.RoutingKey, string(d.Body))
			}
		}()

	}
	wg.Wait()
	return nil

}

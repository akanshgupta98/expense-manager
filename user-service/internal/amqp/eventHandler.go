package amqp

import (
	"user-service/internal/service"

	"github.com/akanshgupta98/expense-manager/contracts/eventspb"
	"github.com/akanshgupta98/go-logger/v2"
	"github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

func processEvent(event amqp091.Delivery) {

	eventType := event.RoutingKey
	logger.Infof("recieved event: %s", eventType)

	switch eventType {
	case "user.created":
		data := &eventspb.UserCreatedEvent{}
		err := proto.Unmarshal(event.Body, data)
		if err != nil {
			logger.Errorf("error during unmarshal: %s", err.Error())
		}
		if err := handleUserCreatedEvent(data); err != nil {
			event.Nack(false, true)
		} else {
			event.Ack(false)

		}

	}

}

func handleUserCreatedEvent(payload *eventspb.UserCreatedEvent) error {
	data := service.CreateProfileInput{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		UserID:    payload.UserId,
		Country:   payload.Country,
	}
	err := service.CreateProfile(data)
	if err != nil {
		logger.Errorf("unable to create user profile for id: %d. Err: %v", payload.UserId, err)
		return err
	}
	return nil

}

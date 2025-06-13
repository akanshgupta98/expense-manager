package service

import (
	"auth-service/internal/config"
	"auth-service/internal/events"
	"auth-service/internal/repository"
	"database/sql"
	"fmt"

	"github.com/akanshgupta98/expense-manager/contracts/eventspb"
	"github.com/akanshgupta98/go-logger"
	"github.com/rabbitmq/amqp091-go"
	"golang.org/x/crypto/bcrypt"
)

var service Service

func Initialize(db *sql.DB, mb *amqp091.Connection, cfg config.Config) {
	service = Service{
		model:  repository.New(db),
		secret: cfg.SecretKey,
		event:  events.New(mb, cfg),
	}

}

func RegisterUser(payload RegisterUserInput) (RegisterUserOutput, error) {
	response := RegisterUserOutput{}
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 12)
	if err != nil {
		return response, err
	}

	data := repository.User{
		Password: string(pwdHash),
		Email:    payload.Email,
	}
	response.UserID, err = service.model.User.CreateUser(data)
	if err != nil {
		return response, err
	}

	eventData := events.EventData{

		Topic: "user.created",
		Data: &eventspb.UserCreatedEvent{

			UserId:    response.UserID,
			Email:     payload.Email,
			FirstName: payload.FirstName,
			LastName:  payload.LastName,
			Country:   payload.Country,
		},
	}

	err = service.event.SendEvent(eventData)
	if err != nil {
		logger.Warnf("unable to publish event on message broker: %s", err.Error())

	} else {
		logger.Infof("published event of user creation")
	}

	return response, nil

}

func LoginUser(data Login) (Token, error) {
	var token Token
	user, err := service.model.User.FetchByEmail(data.Email)
	if err != nil {
		return token, err
	}
	claims := Claims{
		UserID: user.ID,
	}
	if data.Password == user.Password {
		token, err = issueToken(claims)
		if err != nil {
			return token, err
		}

	} else {
		logger.Errorf("password mismatch")
		err = fmt.Errorf("username or password are incorrect")
	}
	return token, err

}

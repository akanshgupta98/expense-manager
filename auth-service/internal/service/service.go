package service

import (
	"auth-service/internal/config"
	"auth-service/internal/repository"
	"database/sql"
	"fmt"

	"github.com/akanshgupta98/go-logger"
	"golang.org/x/crypto/bcrypt"
)

var service Service

func Initialize(db *sql.DB, cfg config.Config) {
	service = Service{
		model:  repository.New(db),
		secret: cfg.SecretKey,
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

	return response, nil

}

func FetchAllUsers() ([]RegisterUserInput, error) {
	var result []RegisterUserInput
	data, err := service.model.User.GetAllUsers()
	if err != nil {
		return nil, err
	}
	for _, u := range data {
		user := RegisterUserInput{
			// Name:     u.Name,
			Email:    u.Email,
			Password: u.Password,
		}
		result = append(result, user)
	}
	return result, nil

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

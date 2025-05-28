package service

import (
	"auth-service/internal/config"
	"auth-service/internal/repository"
	"database/sql"
)

var service Service

func Initialize(db *sql.DB, cfg config.Config) {
	service = Service{
		model:  repository.New(db),
		secret: cfg.SecretKey,
	}

}

func RegisterUser(payload User) error {

	data := repository.User{
		Name:     payload.Name,
		Password: payload.Password,
		Email:    payload.Email,
	}
	err := service.model.User.CreateUser(data)
	if err != nil {
		return err
	}

	return nil

}

func FetchAllUsers() ([]User, error) {
	var result []User
	data, err := service.model.User.GetAllUsers()
	if err != nil {
		return nil, err
	}
	for _, u := range data {
		user := User{
			Name:     u.Name,
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

	}
	return token, err

}

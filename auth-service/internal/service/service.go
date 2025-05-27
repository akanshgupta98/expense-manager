package service

import (
	"auth-service/internal/repository"
	"database/sql"
)

var model repository.Models

func Initialize(db *sql.DB) {
	model = repository.New(db)
}

func RegisterUser(payload User) error {

	data := repository.User{
		Name:     payload.Name,
		Password: payload.Password,
		Email:    payload.Email,
	}
	err := model.User.CreateUser(data)
	if err != nil {
		return err
	}

	return nil

}

func FetchAllUsers() ([]User, error) {
	var result []User
	data, err := model.User.GetAllUsers()
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

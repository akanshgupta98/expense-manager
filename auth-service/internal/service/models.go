package service

import (
	"auth-service/internal/events"
	"auth-service/internal/repository"
)

type RegisterUserInput struct {
	Password  string
	Email     string
	FirstName string
	LastName  string
	Country   string
}

type RegisterUserOutput struct {
	UserID int64
}

type Login struct {
	Email    string
	Password string
}

type Claims struct {
	UserID int
}

type Token struct {
	JWTToken     string
	RefreshToken string
}

type Service struct {
	model  repository.Models
	secret string
	event  *events.Event
}

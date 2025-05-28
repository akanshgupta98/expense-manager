package service

import "auth-service/internal/repository"

type User struct {
	Name     string
	Password string
	Email    string
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
}

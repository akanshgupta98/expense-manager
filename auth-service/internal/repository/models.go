package repository

import "time"

type Models struct {
	User  User
	Token Token
}

type User struct {
	ID       int
	Name     string
	Password string
	Email    string
}

type Token struct {
	ID           int
	RefreshToken string
	UserID       int
	Expiry       time.Duration
	CreatedAt    time.Time
}

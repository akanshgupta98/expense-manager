package service

import "user-service/internal/repo"

var models *repo.Models

type CreateProfileInput struct {
	FirstName string
	LastName  string
	UserID    int
	Country   string
	Email     string
}

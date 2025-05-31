package service

import (
	"database/sql"
	"user-service/internal/repo"

	"github.com/akanshgupta98/go-logger/v2"
)

func Initialize(db *sql.DB) {
	models = repo.New(db)
}
func CreateProfile(profile CreateProfileInput) error {

	data := repo.UserProfile{
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		Email:     profile.Email,
		Country:   profile.Country,
		UserID:    profile.UserID,
	}
	err := models.UserProfile.AddProfile(data)
	if err != nil {
		return err
	}
	logger.Debugf("profile created successfully")

	return nil
}

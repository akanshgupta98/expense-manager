package repo

import (
	"database/sql"
	"time"
)

var db *sql.DB

type Models struct {
	UserProfile UserProfile
}

type UserProfile struct {
	ID        int
	FirstName string
	LastName  string
	UserID    int
	Country   string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

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

type TABLE_NAME = string
type COLUMN_NAMES = []string

const (
	AUTH_TABLE    = "AUTH"
	AUTH_TABLE_PK = "ID"
)

var dbData map[TABLE_NAME]COLUMN_NAMES = map[TABLE_NAME]COLUMN_NAMES{
	AUTH_TABLE: []string{"EMAIL", "PASSWORD"},
}

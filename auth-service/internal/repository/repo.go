package repository

import (
	"database/sql"
	"time"

	"github.com/akanshgupta98/go-logger"
)

var db *sql.DB

func New(dbpool *sql.DB) Models {
	db = dbpool
	return Models{
		User:  User{},
		Token: Token{},
	}
}

func (u *User) CreateUser(data User) (int, error) {

	var userID int
	query := `INSERT INTO USERS (NAME,EMAIL,PASSWORD) VALUES ($1,$2,$3) RETURNING ID`

	rows, err := db.Query(query, data.Name, data.Email, data.Password)
	if err != nil {
		return userID, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&userID)
	}
	return userID, err

}

func (u *User) GetAllUsers() ([]User, error) {
	result := make([]User, 0)
	query := `SELECT * FROM USERS`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, user)
	}
	return result, nil

}

func (u *User) FetchByEmail(email string) (User, error) {
	query := `SELECT * FROM USERS WHERE EMAIL=$1`
	var user User

	rows, err := db.Query(query, email)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
		)
		if err != nil {
			return user, err
		}
	}
	return user, nil
}

func (t *Token) CreateToken(data Token) error {
	query := `INSERT INTO TOKEN (R_TOKEN, USER_ID,EXPIRY,CREATED_AT) VALUES($1,$2,$3,$4)`
	expiry := time.Now().Add(data.Expiry)
	expiry = expiry.UTC()
	logger.Debugf("Expiry time is: %v", expiry)
	logger.Debugf("Current time is: %v", time.Now().UTC())

	rows, err := db.Query(query, data.RefreshToken, data.UserID, expiry, time.Now().UTC())
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

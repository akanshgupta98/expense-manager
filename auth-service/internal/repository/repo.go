package repository

import "database/sql"

var db *sql.DB

func New(dbpool *sql.DB) Models {
	db = dbpool
	return Models{User: User{}}
}

func (u *User) CreateUser(data User) error {

	query := `INSERT INTO USERS (NAME,EMAIL,PASSWORD) VALUES ($1,$2,$3)`

	_, err := db.Query(query, data.Name, data.Email, data.Password)
	if err != nil {
		return err
	}
	return nil

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

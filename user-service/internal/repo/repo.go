package repo

import (
	"database/sql"
	"time"

	"github.com/akanshgupta98/go-logger/v2"
)

func New(dbpool *sql.DB) *Models {
	db = dbpool
	return &Models{
		UserProfile: UserProfile{},
	}
}

func (up *UserProfile) AddProfile(profile UserProfile) error {
	query := `INSERT INTO USERS (FIRST_NAME,LAST_NAME,USER_ID,EMAIL,COUNTRY,CREATED_AT,UPDATED_AT) VALUES ($1,$2,$3,$4,$5,$6,$7)`
	rows, err := db.Query(query,
		profile.FirstName, // FIRST_NAME
		profile.LastName,  // LAST_NAME
		profile.UserID,    // USER_ID
		profile.Email,     // EMAIL
		profile.Country,   // COUNTRY
		time.Now().UTC(),  // CREATED_AT
		time.Now().UTC(),  // UPDATED_AT
	)
	if err != nil {
		return err
	}

	defer rows.Close()
	logger.Debugf("SQL Query result is: %v", rows)
	return nil

}

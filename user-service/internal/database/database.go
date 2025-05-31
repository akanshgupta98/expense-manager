package database

import (
	"database/sql"
	"time"

	"github.com/akanshgupta98/go-logger/v2"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func connect(DSN string) (*sql.DB, error) {

	db, err := sql.Open("pgx", DSN)
	if err != nil {
		return db, err
	}
	err = db.Ping()
	if err != nil {
		return db, err
	}
	return db, nil
}

func ConnectToDB(DSN string) (*sql.DB, error) {
	var db *sql.DB
	var err error
	for count := 0; count < 10; count++ {
		db, err = connect(DSN)
		if err == nil {
			return db, nil
		} else {
			logger.Debugf("unable to connect to DB")
			logger.Debugf("Backing off for 2 seconds...")
			time.Sleep(2 * time.Second)
		}
	}
	return db, err

}

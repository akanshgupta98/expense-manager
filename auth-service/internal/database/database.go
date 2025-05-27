package database

import (
	"database/sql"
	"os"
	"time"

	"github.com/akanshgupta98/go-logger"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func Connect() (*sql.DB, error) {
	var db *sql.DB
	var err error
	dsn := os.Getenv("DSN")

	for count := 0; count < 10; count++ {
		db, err = sql.Open("pgx", dsn)
		if err != nil {
			return db, err
		}

		err = db.Ping()
		if err != nil {
			logger.Debugf("Unable to connect to DB. Retrying...")
			time.Sleep(2 * time.Second)

		} else {
			logger.Debugf("successfully connected to DB.")
			break
		}
	}

	return db, nil
}

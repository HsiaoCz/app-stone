package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var db *sql.DB

func InitDB() error {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}

	logrus.Info("Successfully connected to the database")

	return nil
}

func GetDB() *sql.DB {
	return db
}

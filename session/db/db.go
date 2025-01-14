package db

import (
	"log"
	"os"

	"github.com/anthdm/superkit/db"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"

	_ "github.com/mattn/go-sqlite3"
)

// By default this is a pre-configured Gorm DB instance.
// Change this type based on the database package of your likings.
var dbInstance *bun.DB

// Get returns the instantiated DB instance.
func Get() *bun.DB {
	return dbInstance
}

func Init() error {
	// Create a default *sql.DB exposed by the superkit/db package
	// based on the given configuration.
	config := db.Config{
		Driver:   os.Getenv("DB_DRIVER"),
		Name:     os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		Host:     os.Getenv("DB_HOST"),
	}
	dbinst, err := db.NewSQL(config)
	if err != nil {
		return err
	}
	// Based on the driver create the corresponding DB instance.
	// By default, the SuperKit boilerplate comes with a pre-configured
	// ORM called Gorm. https://gorm.io.
	//
	// You can change this to any other DB interaction tool
	// of your liking. EG:
	// - uptrace bun -> https://bun.uptrace.dev
	// - SQLC -> https://github.com/sqlc-dev/sqlc
	// - gojet -> https://github.com/go-jet/jet
	switch config.Driver {
	case db.DriverSqlite3:
		dbInstance = bun.NewDB(dbinst, sqlitedialect.New())
	case db.DriverMysql:
		// ...
	default:
		log.Fatal("invalid driver:", config.Driver)
	}
	return nil
}

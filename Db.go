package main

import (
	"fmt"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"

	"github.com/jmoiron/sqlx"
)

// connectDB creates and returns a database connection.
func ConnectDB(cfg DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open(cfg.Type, cfg.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("error connecting to db : %v", err)
	}

	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxActiveConns)
	db.SetConnMaxLifetime(time.Second * cfg.ConnectTimeout)

	// Ping database to check for connection issues.
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to db : %v", err)
	}

	return db, nil
}

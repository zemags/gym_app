package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	DBName string
}

func NewSQLiteDB(cfg Config) (*sqlx.DB, error) {
	conString := fmt.Sprintf("../../%s", cfg.DBName)
	db, err := sqlx.Open("sqlite3", conString)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

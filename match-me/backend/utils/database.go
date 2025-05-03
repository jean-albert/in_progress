package utils

import (
	"database/sql"
	"fmt"
	"match-me/backend/models"
)

func InitDB(cfg config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func CreateUser(db *sql.DB, email, password string) (int, error) {
	// Implementation for user creation
}

func GetProfile(db *sql.DB, userID int) (models.Profile, error) {
	// Implementation for getting profile
}

// Add other database operations similarly

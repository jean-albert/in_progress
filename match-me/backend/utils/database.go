package utils

import (
	"database/sql"
	"fmt"
	"match-me-backend/config"
	"match-me-backend/models"

	"github.com/lib/pq"
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

func GetBio(db *sql.DB, userID int) (models.Bio, error) {
	var bio models.Bio
	err := db.QueryRow("SELECT user_id, interests, looking_for, age, gender, max_distance FROM bios WHERE user_id = $1", userID).Scan(
		&bio.UserID,
		pq.Array(&bio.Interests),
		&bio.LookingFor,
		&bio.Age,
		&bio.Gender,
		&bio.MaxDistance,
	)
	if err != nil {
		return models.Bio{}, fmt.Errorf("failed to get bio: %w", err) // Explicit return
	}
	return bio, nil // Explicit return
}

package db

import (
	"database/sql"
	"news-aggregator/internal/models"
)

// CreateUser inserts a new user into the database.
func CreateUser(u *models.User) error {
	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", u.Username, u.Password)
	return err
}

// GetUserByUsername retrieves a user by their username.
func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetAllUsers retrieves all users from the database.
func GetAllUsers() ([]models.User, error) {
	rows, err := db.Query("SELECT id, username, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

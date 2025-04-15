package repository

import (
	"database/sql"
	"fmt"
	"spotify-cli/internal/models"
	"strings"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {
	query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	// Trim whitespace from the input
	username = strings.TrimSpace(username)

	// Optional debug logging
	fmt.Println("Looking up user with username:", username)

	query := "SELECT id, username, password, created_at FROM users WHERE username = ?"
	row := r.db.QueryRow(query, username)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) UsernameExists(username string) (bool, error) {
	query := "SELECT 1 FROM users WHERE username = ? LIMIT 1"
	row := r.db.QueryRow(query, username)

	var exists int
	err := row.Scan(&exists)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *UserRepository) Create(user *models.User) (int, error) {
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	result, err := r.db.Exec(query, user.Username, user.Password)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

func (r *UserRepository) Update(user *models.User) error {
	query := "UPDATE users SET username = ?, password = ? WHERE id = ?"
	_, err := r.db.Exec(query, user.Username, user.Password, user.ID)
	return err
}

func (r *UserRepository) Delete(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}

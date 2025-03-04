package repositories

import (
	"api/internal/models"
	"database/sql"
)

type DefaultUserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) models.UserRepository {
	return &DefaultUserRepository{db}
}

func (r *DefaultUserRepository) CreateUser(user *models.User) error {
	return nil
}

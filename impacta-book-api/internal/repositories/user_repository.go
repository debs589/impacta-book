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

func (r *DefaultUserRepository) CreateUser(user models.User) (int, error) {
	statement, err := r.db.Prepare("INSERT INTO users(name, nickName, email, password) values(?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

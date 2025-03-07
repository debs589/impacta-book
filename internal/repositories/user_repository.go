package repositories

import (
	"api/internal/models"
	"database/sql"
	"fmt"
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

func (r *DefaultUserRepository) GetUsers(nameOrNick string) ([]models.User, error) {

	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	rows, err := r.db.Query("SELECT id, name, nickName, email, createdAt FROM users WHERE name LIKE ? or nickName LIKE ?", nameOrNick, nameOrNick)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}

	for rows.Next() {
		user := models.User{}

		err = rows.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *DefaultUserRepository) GetUser(id int) (models.User, error) {
	user := models.User{}

	row := r.db.QueryRow("SELECT id, name, nickName, email, createdAt FROM users WHERE id = ?", id)
	err := row.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

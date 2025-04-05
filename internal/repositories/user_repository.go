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

func (r *DefaultUserRepository) GetUserByEmail(email string) (models.User, error) {
	user := models.User{}

	row := r.db.QueryRow("SELECT id, password FROM users WHERE email = ?", email)
	err := row.Scan(&user.ID, &user.Password)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *DefaultUserRepository) UpdateUser(id int, user models.User) error {
	statement, err := r.db.Prepare("UPDATE users SET name = ?, nickName = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Nickname, user.Email, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *DefaultUserRepository) DeleteUser(id int) error {
	statement, err := r.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (r *DefaultUserRepository) FollowUser(followerId, userId int) error {
	statement, err := r.db.Prepare("INSERT INTO followers (follower_id, user_id) values(?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(followerId, userId)
	if err != nil {
		return err
	}

	return nil
}

func (r *DefaultUserRepository) UnfollowUser(userId, followerId int) error {
	statement, err := r.db.Prepare("DELETE FROM followers WHERE user_id = ? and follower_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(userId, followerId)
	if err != nil {
		return err
	}

	return nil
}

func (r *DefaultUserRepository) GetFollowers(userId int) ([]models.User, error) {
	rows, err := r.db.Query("SELECT id, name, nickName, email, createdAt FROM users u "+
		"INNER JOIN followers f ON u.id = f.follower_id WHERE f.user_id = ?", userId)

	if err != nil {
		return []models.User{}, err
	}
	defer rows.Close()

	followers := []models.User{}

	for rows.Next() {
		user := models.User{}

		err = rows.Scan(&user.ID, &user.Name, &user.Nickname, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}

		followers = append(followers, user)
	}

	err = rows.Err()
	if err != nil {
		return []models.User{}, err
	}

	return followers, nil

}

package models

type User struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Nickname  string `json:"nickname,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type UserService interface {
	CreateUser(User, string) (int, error)
	GetUsers(string) ([]User, error)
	GetUser(int) (User, error)
}

type UserRepository interface {
	CreateUser(user User) (int, error)
	GetUsers(string) ([]User, error)
	GetUser(int) (User, error)
}

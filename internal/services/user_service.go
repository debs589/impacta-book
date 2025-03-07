package services

import (
	"api/internal/models"
	"api/internal/security"
	"api/internal/utils"
	"errors"
	"github.com/badoux/checkmail"
	"strings"
)

type DefaultUserService struct {
	rp models.UserRepository
}

func NewUserService(rp models.UserRepository) models.UserService {
	return &DefaultUserService{rp}
}

func (s *DefaultUserService) CreateUser(user models.User, step string) (int, error) {
	verify := s.validate(user, step)

	if verify != nil {
		return 0, utils.ErrInvalidArguments
	}

	formatUser, err := s.format(user, step)
	if err != nil {
		return 0, utils.ErrInvalidArguments
	}

	id, err := s.rp.CreateUser(formatUser)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *DefaultUserService) validate(user models.User, step string) error {
	if len(user.Name) == 0 {
		return errors.New("Name is required and cannot be empty")
	}

	if len(user.Nickname) == 0 {
		return errors.New("Nickname is required and cannot be empty")
	}

	if len(user.Email) == 0 {
		return errors.New("Email is required and cannot be empty")
	}

	if emailCheck := checkmail.ValidateFormat(user.Email); emailCheck != nil {
		return emailCheck
	}

	if len(user.Password) == 0 && step == "register" {
		return errors.New("Password is required and cannot be empty")
	}

	return nil
}

func (s *DefaultUserService) format(user models.User, step string) (models.User, error) {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Nickname = strings.TrimSpace(user.Nickname)

	if step == "register" {
		passwordHash, err := security.Hash(user.Password)
		if err != nil {
			return models.User{}, err
		}

		user.Password = string(passwordHash)
	}

	return user, nil
}

func (s *DefaultUserService) GetUsers(nameOrNick string) ([]models.User, error) {
	users, err := s.rp.GetUsers(nameOrNick)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *DefaultUserService) GetUser(id int) (models.User, error) {
	user, err := s.rp.GetUser(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

package services

import (
	"api/internal/models"
	"errors"
	"github.com/badoux/checkmail"
)

type DefaultUserService struct {
	rp models.UserRepository
}

func NewUserService(rp models.UserRepository) models.UserService {
	return &DefaultUserService{rp}
}

func (s *DefaultUserService) CreateUser(user models.User) (int, error) {
	verify := s.validate(user)

	if verify != nil {
		return 0, verify
	}

	id, err := s.rp.CreateUser(user)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *DefaultUserService) validate(user models.User) error {
	if len(user.Name) == 0 {
		return errors.New("Name is required and cannot be empty")
	}

	if len(user.Nickname) == 0 {
		return errors.New("Nickname is required and cannot be empty")
	}

	if len(user.Email) == 0 {
		return errors.New("Email is required and cannot be empty")
	}

	if error := checkmail.ValidateFormat(user.Email); error != nil {
		return errors.New("Email is invalid")
	}

	return nil
}

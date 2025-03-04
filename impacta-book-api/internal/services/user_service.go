package services

import "api/internal/models"

type DefaultUserService struct {
	rp models.UserRepository
}

func NewUserService(rp models.UserRepository) models.UserService {
	return &DefaultUserService{rp}
}

func (s *DefaultUserService) CreateUser(user *models.User) error {
	return nil
}

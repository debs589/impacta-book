package services

import (
	"api/internal/models"
	"api/internal/utils"
	"errors"
	"strings"
)

type DefaultPublicationService struct {
	rp models.PublicationRepository
}

func NewPublicationService(rp models.PublicationRepository) models.PublicationService {
	return &DefaultPublicationService{rp}
}

func (s *DefaultPublicationService) CreatePublication(publication models.Publication) (int, error) {
	verify := s.validate(publication)

	if verify != nil {
		return 0, utils.ErrInvalidArguments
	}

	s.format(&publication)

	id, err := s.rp.CreatePublication(publication)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *DefaultPublicationService) GetPublication(id int) (models.Publication, error) {
	publication, err := s.rp.GetPublication(id)
	if err != nil {
		return models.Publication{}, err
	}
	return publication, nil
}

func (s *DefaultPublicationService) GetPublications(id int) ([]models.Publication, error) {
	publication, err := s.rp.GetPublications(id)
	if err != nil {
		return []models.Publication{}, err
	}
	return publication, nil
}

func (s *DefaultPublicationService) validate(publication models.Publication) error {
	if len(publication.Title) == 0 {
		return errors.New("Title is required and cannot be empty")
	}

	if len(publication.Content) == 0 {
		return errors.New("Content is required and cannot be empty")
	}

	return nil
}

func (s *DefaultPublicationService) format(publication *models.Publication) {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)
}

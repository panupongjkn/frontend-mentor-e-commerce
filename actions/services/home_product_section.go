package services

import (
	"e-commerce-api/actions/repositories"
	"e-commerce-api/models"
)

type homeProductSectionService struct {
	homeProductSectionRepository repositories.HomeProductSectionRepository
}

func NewHomeProductSectionService(homeProductSectionRepository repositories.HomeProductSectionRepository) homeProductSectionService {
	return homeProductSectionService{homeProductSectionRepository: homeProductSectionRepository}
}

func (s homeProductSectionService) GetAll() ([]models.HomeProductSection, error) {
	homeProductSections, err := s.homeProductSectionRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return homeProductSections, nil
}

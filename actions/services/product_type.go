package services

import (
	"e-commerce-api/actions/repositories"
	"e-commerce-api/models"
)

type productTypeService struct {
	productTypeRepository repositories.ProductTypeRepository
}

func NewProductTypeService(productTypeRepository repositories.ProductTypeRepository) productTypeService {
	return productTypeService{productTypeRepository: productTypeRepository}
}

func (s productTypeService) GetByHomeSection() ([]models.ProductType, error) {
	productTypes, err := s.productTypeRepository.GetByHomeSection()
	if err != nil {
		return nil, err
	}
	return productTypes, nil
}

func (s productTypeService) GetBySlug(slug string) (*models.ProductType, error) {
	productType, err := s.productTypeRepository.GetBySlug(slug)
	if err != nil {
		return nil, err
	}
	return productType, nil
}

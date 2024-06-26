package services

import (
	"e-commerce-api/actions/repositories"
	"e-commerce-api/models"
)

type productService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) productService {
	return productService{productRepository: productRepository}
}

func (s productService) GetBySlug(slug string) (*models.Product, error) {
	product, err := s.productRepository.GetBySlug(slug)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s productService) GetByProductTypeSlug(productTypeSlug string) ([]models.ProductByProductType, error) {
	products, err := s.productRepository.GetByProductType(productTypeSlug)
	if err != nil {
		return nil, err
	}
	return products, nil
}

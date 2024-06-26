package services

import "e-commerce-api/models"

type ProductTypeService interface {
	GetByHomeSection() ([]models.ProductType, error)
	GetBySlug(slug string) (*models.ProductType, error)
}

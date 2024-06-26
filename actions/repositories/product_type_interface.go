package repositories

import "e-commerce-api/models"

type ProductTypeRepository interface {
	GetByHomeSection() ([]models.ProductType, error)
	GetBySlug(slug string) (*models.ProductType, error)
}

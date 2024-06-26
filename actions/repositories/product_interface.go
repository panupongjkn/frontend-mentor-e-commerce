package repositories

import "e-commerce-api/models"

type ProductRepository interface {
	GetBySlug(slug string) (*models.Product, error)
	GetByProductType(productTypeSlug string) ([]models.ProductByProductType, error)
}

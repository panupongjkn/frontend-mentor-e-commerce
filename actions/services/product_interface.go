package services

import "e-commerce-api/models"

type ProductService interface {
	GetBySlug(slug string) (*models.Product, error)
	GetByProductTypeSlug(productTypeSlug string) ([]models.ProductByProductType, error)
}

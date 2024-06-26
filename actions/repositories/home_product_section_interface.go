package repositories

import "e-commerce-api/models"

type HomeProductSectionRepository interface {
	GetAll() ([]models.HomeProductSection, error)
}

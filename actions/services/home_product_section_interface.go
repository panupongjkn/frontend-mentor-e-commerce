package services

import "e-commerce-api/models"

type HomeProductSectionService interface {
	GetAll() ([]models.HomeProductSection, error)
}

package services

import "e-commerce-api/models"

type BannerService interface {
	GetAll() ([]models.Banner, error)
}

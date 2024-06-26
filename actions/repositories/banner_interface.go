package repositories

import "e-commerce-api/models"

type BannerRepository interface {
	GetAll() ([]models.Banner, error)
}

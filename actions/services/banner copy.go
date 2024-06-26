package services

import (
	"e-commerce-api/actions/repositories"
	"e-commerce-api/models"
)

type bannerService struct {
	bannerRepository repositories.BannerRepository
}

func NewBannerService(bannerRepository repositories.BannerRepository) bannerService {
	return bannerService{bannerRepository: bannerRepository}
}

func (s bannerService) GetAll() ([]models.Banner, error) {
	banners, err := s.bannerRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return banners, nil
}

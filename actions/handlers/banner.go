package handlers

import (
	"e-commerce-api/actions/services"

	"github.com/gin-gonic/gin"
)

type bannerHandler struct {
	bannerService services.BannerService
}

func NewBannerHandler(bannerService services.BannerService) bannerHandler {
	return bannerHandler{bannerService: bannerService}
}

func (h bannerHandler) GetAll(c *gin.Context) {
	banners, err := h.bannerService.GetAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"data":    banners,
			"message": "Get banners success",
		})
	}
}

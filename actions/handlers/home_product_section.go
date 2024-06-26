package handlers

import (
	"e-commerce-api/actions/services"

	"github.com/gin-gonic/gin"
)

type homeProductSectionHandler struct {
	homeProductSectionService services.HomeProductSectionService
}

func NewHomeProductSectionHandler(homeProductSectionService services.HomeProductSectionService) homeProductSectionHandler {
	return homeProductSectionHandler{homeProductSectionService: homeProductSectionService}
}

func (h homeProductSectionHandler) GetAll(c *gin.Context) {
	homeProductSections, err := h.homeProductSectionService.GetAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"data":    homeProductSections,
			"message": "Get home product section success.",
		})
	}
}

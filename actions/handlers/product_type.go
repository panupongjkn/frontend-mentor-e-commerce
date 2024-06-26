package handlers

import (
	"e-commerce-api/actions/services"

	"github.com/gin-gonic/gin"
)

type productTypeHandler struct {
	productTypeService services.ProductTypeService
}

func NewProductTypeHandler(productTypeService services.ProductTypeService) productTypeHandler {
	return productTypeHandler{productTypeService: productTypeService}
}

func (h productTypeHandler) GetByHomeSection(c *gin.Context) {
	productTypes, err := h.productTypeService.GetByHomeSection()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"data":    productTypes,
			"message": "Get home product type section success.",
		})
	}
}

func (h productTypeHandler) GetBySlug(c *gin.Context) {
	productType, err := h.productTypeService.GetBySlug(c.Param("slug"))
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"data":    productType,
			"message": "Get product type success.",
		})
	}
}

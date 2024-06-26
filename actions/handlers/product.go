package handlers

import (
	"e-commerce-api/actions/services"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) productHandler {
	return productHandler{productService: productService}
}

func (h productHandler) GetBySlug(c *gin.Context) {
	product, err := h.productService.GetBySlug(c.Param("slug"))
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"data":    product,
			"message": "Get product success.",
		})
	}
}

func (h productHandler) GetByProductTypeSlug(c *gin.Context) {
	products, err := h.productService.GetByProductTypeSlug(c.Param("slug"))
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"data":    products,
			"message": "Get products by product type success.",
		})
	}
}

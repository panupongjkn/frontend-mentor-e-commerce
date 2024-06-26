package main

import (
	"e-commerce-api/actions/handlers"
	"e-commerce-api/actions/repositories"
	"e-commerce-api/actions/services"
	"e-commerce-api/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	os.Setenv("TZ", "Asia/Bangkok")
	err := godotenv.Load()
	if err != nil {
		log.Print("ERROR: loading .env fail")
	}
}

// func initValidator() {

// }

func main() {
	db := database.Init()
	r := gin.Default()
	// validate := validator.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	bannerRepository := repositories.NewBannerRepositoryDB(db)
	productTypeRepository := repositories.NewProductTypeRepositoryDB(db)
	productRepository := repositories.NewProductRepositoryDB(db)
	homeProductSectionRepository := repositories.NewHomeProductSectionRepositoryDB(db)

	bannerService := services.NewBannerService(bannerRepository)
	productTypeService := services.NewProductTypeService(productTypeRepository)
	productService := services.NewProductService(productRepository)
	homeProductSectionService := services.NewHomeProductSectionService(homeProductSectionRepository)

	bannerHandlers := handlers.NewBannerHandler(bannerService)
	productTypeHandlers := handlers.NewProductTypeHandler(productTypeService)
	productHandlers := handlers.NewProductHandler(productService)
	homeProductSectionHandlers := handlers.NewHomeProductSectionHandler(homeProductSectionService)

	r.GET("/home/banner", bannerHandlers.GetAll)
	r.GET("/home/product-type-section", productTypeHandlers.GetByHomeSection)
	r.GET("/home/product-section", homeProductSectionHandlers.GetAll)

	r.GET("/product/:slug", productHandlers.GetBySlug)

	r.GET("/product-type/:slug", productTypeHandlers.GetBySlug)
	r.GET("/product-type/:slug/product", productHandlers.GetByProductTypeSlug)

	r.Run(":" + os.Getenv("PORT")) // listen and serve on 0.0.0.0:8080
}

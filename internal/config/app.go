package config

import (
	"github.com/aldisypu/go-simple-pos/internal/controller"
	"github.com/aldisypu/go-simple-pos/internal/repository"
	"github.com/aldisypu/go-simple-pos/internal/router"
	"github.com/aldisypu/go-simple-pos/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	categoryRepository := repository.NewCategoryRepository(config.Log)
	productRepository := repository.NewProductRepository(config.Log)
	saleRepository := repository.NewSaleRepository(config.Log)
	salesDetailRepository := repository.NewSalesDetailRepository(config.Log)

	// setup service
	categoryService := service.NewCategoryService(config.DB, config.Log, config.Validate, categoryRepository)
	productService := service.NewProductService(config.DB, config.Log, config.Validate, productRepository)
	saleService := service.NewSaleService(config.DB, config.Log, config.Validate, saleRepository, productRepository, salesDetailRepository)
	salesDetailService := service.NewSalesDetailService(config.DB, config.Log, config.Validate, productRepository, saleRepository, salesDetailRepository)

	// setup controller
	categoryController := controller.NewCategoryController(categoryService, config.Log)
	productController := controller.NewProductController(productService, config.Log)
	saleController := controller.NewSaleController(saleService, config.Log)
	salesDetailController := controller.NewSalesDetailController(salesDetailService, config.Log)

	routeConfig := router.RouteConfig{
		App:                   config.App,
		CategoryController:    categoryController,
		ProductController:     productController,
		SaleController:        saleController,
		SalesDetailController: salesDetailController,
	}
	routeConfig.Setup()
}

package router

import (
	"github.com/aldisypu/go-simple-pos/internal/controller"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                   *fiber.App
	CategoryController    controller.CategoryController
	ProductController     controller.ProductController
	SaleController        controller.SaleController
	SalesDetailController controller.SalesDetailController
}

func (c *RouteConfig) Setup() {
	c.SetupRoute()
}

func (c *RouteConfig) SetupRoute() {
	c.App.Post("/api/categories", c.CategoryController.Create)
	c.App.Put("/api/categories/:categoryId", c.CategoryController.Update)
	c.App.Delete("/api/categories/:categoryId", c.CategoryController.Delete)
	c.App.Get("/api/categories/:categoryId", c.CategoryController.Get)
	c.App.Get("/api/categories", c.CategoryController.List)

	c.App.Post("/api/products", c.ProductController.Create)
	c.App.Put("/api/products/:productId", c.ProductController.Update)
	c.App.Delete("/api/products/:productId", c.ProductController.Delete)
	c.App.Get("/api/products/:productId", c.ProductController.Get)
	c.App.Get("/api/products", c.ProductController.List)

	c.App.Post("/api/sales", c.SaleController.Create)
	c.App.Delete("/api/sales/:saleId", c.SaleController.Delete)
	c.App.Get("/api/sales/:saleId", c.SaleController.Get)
	c.App.Get("/api/sales", c.SaleController.List)

	c.App.Put("/api/sales/:saleId/details/:salesDetailId", c.SalesDetailController.Update)
	c.App.Delete("/api/sales/:saleId/details/:salesDetailId", c.SalesDetailController.Delete)
	c.App.Get("/api/sales/:saleId/details/:salesDetailId", c.SalesDetailController.Get)
	c.App.Get("/api/sales/:saleId/details", c.SalesDetailController.List)
}

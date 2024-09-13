package controller

import (
	"github.com/aldisypu/go-simple-pos/internal/model/web"
	"github.com/aldisypu/go-simple-pos/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
	Log            *logrus.Logger
}

func NewProductController(productService service.ProductService, log *logrus.Logger) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
		Log:            log,
	}
}

func (c *ProductControllerImpl) Create(ctx *fiber.Ctx) error {
	request := new(web.CreateProductRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parsing request body")
		return fiber.ErrBadRequest
	}

	response, err := c.ProductService.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to creating product")
		return err
	}

	return ctx.JSON(web.WebResponse[*web.ProductResponse]{Data: response})
}

func (c *ProductControllerImpl) Update(ctx *fiber.Ctx) error {
	request := new(web.UpdateProductRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parsing request body")
		return fiber.ErrBadRequest
	}

	request.ID = ctx.Params("productId")

	response, err := c.ProductService.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to updating product")
		return err
	}

	return ctx.JSON(web.WebResponse[*web.ProductResponse]{Data: response})
}

func (c *ProductControllerImpl) Delete(ctx *fiber.Ctx) error {
	productId := ctx.Params("productId")

	request := &web.DeleteProductRequest{
		ID: productId,
	}

	if err := c.ProductService.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to deleting product")
		return err
	}

	return ctx.JSON(web.WebResponse[bool]{Data: true})
}

func (c *ProductControllerImpl) Get(ctx *fiber.Ctx) error {
	request := &web.GetProductRequest{
		ID: ctx.Params("productId"),
	}

	response, err := c.ProductService.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to getting product")
		return err
	}

	return ctx.JSON(web.WebResponse[*web.ProductResponse]{Data: response})
}

func (c *ProductControllerImpl) List(ctx *fiber.Ctx) error {
	responses, err := c.ProductService.List(ctx.UserContext())
	if err != nil {
		c.Log.WithError(err).Error("failed to to list products")
		return err
	}

	return ctx.JSON(web.WebResponse[[]web.ProductResponse]{Data: responses})
}

package controller

import (
	"github.com/aldisypu/go-simple-pos/internal/model/web"
	"github.com/aldisypu/go-simple-pos/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
	Log             *logrus.Logger
}

func NewCategoryController(categoryService service.CategoryService, log *logrus.Logger) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
		Log:             log,
	}
}

func (c *CategoryControllerImpl) Create(ctx *fiber.Ctx) error {
	request := new(web.CreateCategoryRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parsing request body")
		return fiber.ErrBadRequest
	}

	response, err := c.CategoryService.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to creating category")
		return err
	}

	return ctx.JSON(web.WebResponse[*web.CategoryResponse]{Data: response})
}

func (c *CategoryControllerImpl) Update(ctx *fiber.Ctx) error {
	request := new(web.UpdateCategoryRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parsing request body")
		return fiber.ErrBadRequest
	}

	request.ID = ctx.Params("categoryId")

	response, err := c.CategoryService.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to updating category")
		return err
	}

	return ctx.JSON(web.WebResponse[*web.CategoryResponse]{Data: response})
}

func (c *CategoryControllerImpl) Delete(ctx *fiber.Ctx) error {
	categoryId := ctx.Params("categoryId")

	request := &web.DeleteCategoryRequest{
		ID: categoryId,
	}

	if err := c.CategoryService.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to deleting category")
		return err
	}

	return ctx.JSON(web.WebResponse[bool]{Data: true})
}

func (c *CategoryControllerImpl) Get(ctx *fiber.Ctx) error {
	request := &web.GetCategoryRequest{
		ID: ctx.Params("categoryId"),
	}

	response, err := c.CategoryService.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to getting category")
		return err
	}

	return ctx.JSON(web.WebResponse[*web.CategoryResponse]{Data: response})
}

func (c *CategoryControllerImpl) List(ctx *fiber.Ctx) error {
	responses, err := c.CategoryService.List(ctx.UserContext())
	if err != nil {
		c.Log.WithError(err).Error("failed to list categories")
		return err
	}

	return ctx.JSON(web.WebResponse[[]web.CategoryResponse]{Data: responses})
}

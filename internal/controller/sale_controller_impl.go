package controller

import (
	"github.com/aldisypu/go-simple-pos/internal/model/web"
	"github.com/aldisypu/go-simple-pos/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type SaleControllerImpl struct {
	SaleService service.SaleService
	Log         *logrus.Logger
}

func NewSaleController(saleService service.SaleService, log *logrus.Logger) SaleController {
	return &SaleControllerImpl{
		SaleService: saleService,
		Log:         log,
	}
}

func (c *SaleControllerImpl) Create(ctx *fiber.Ctx) error {
	request := new(web.CreateSaleRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parsing request body")
		return fiber.ErrBadRequest
	}

	response, err := c.SaleService.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to creating sale")
		return err
	}

	return ctx.JSON(web.WebResponse[*web.SaleResponse]{Data: response})
}

func (c *SaleControllerImpl) Delete(ctx *fiber.Ctx) error {
	saleId := ctx.Params("saleId")

	request := &web.DeleteSaleRequest{
		ID: saleId,
	}

	if err := c.SaleService.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to deleting sale")
		return err
	}

	return ctx.JSON(web.WebResponse[bool]{Data: true})
}

func (c *SaleControllerImpl) Get(ctx *fiber.Ctx) error {
	request := &web.GetSaleRequest{
		ID: ctx.Params("saleId"),
	}

	response, err := c.SaleService.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to getting sale")
		return err
	}

	return ctx.JSON(web.WebResponse[*web.SaleResponse]{Data: response})
}

func (c *SaleControllerImpl) List(ctx *fiber.Ctx) error {
	responses, err := c.SaleService.List(ctx.UserContext())
	if err != nil {
		c.Log.WithError(err).Error("failed to list sales")
		return err
	}

	return ctx.JSON(web.WebResponse[[]web.SaleResponse]{Data: responses})
}

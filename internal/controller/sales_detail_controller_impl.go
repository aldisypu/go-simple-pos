package controller

import (
	"github.com/aldisypu/go-simple-pos/internal/model/web"
	"github.com/aldisypu/go-simple-pos/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type SalesDetailControllerImpl struct {
	SalesDetailService service.SalesDetailService
	Log                *logrus.Logger
}

func NewSalesDetailController(salesDetailService service.SalesDetailService, log *logrus.Logger) SalesDetailController {
	return &SalesDetailControllerImpl{
		SalesDetailService: salesDetailService,
		Log:                log,
	}
}

func (c *SalesDetailControllerImpl) Update(ctx *fiber.Ctx) error {
	request := new(web.UpdateSalesDetailRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parsing request body")
		return fiber.ErrBadRequest
	}

	request.SaleId = ctx.Params("saleId")
	request.ID = ctx.Params("salesDetailId")

	response, err := c.SalesDetailService.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to updating sales detail")
		return err
	}

	return ctx.JSON(web.WebResponse[*web.SalesDetailResponse]{Data: response})
}

func (c *SalesDetailControllerImpl) Delete(ctx *fiber.Ctx) error {
	saleId := ctx.Params("saleId")
	salesDetailId := ctx.Params("salesDetailId")

	request := &web.DeleteSalesDetailRequest{
		SaleId: saleId,
		ID:     salesDetailId,
	}

	if err := c.SalesDetailService.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to deleting sales detail")
		return err
	}

	return ctx.JSON(web.WebResponse[bool]{Data: true})
}

func (c *SalesDetailControllerImpl) Get(ctx *fiber.Ctx) error {
	saleId := ctx.Params("saleId")
	salesDetailId := ctx.Params("salesDetailId")

	request := &web.GetSalesDetailRequest{
		SaleId: saleId,
		ID:     salesDetailId,
	}

	response, err := c.SalesDetailService.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to getting sales detail")
		return err
	}

	return ctx.JSON(web.WebResponse[*web.SalesDetailResponse]{Data: response})
}

func (c *SalesDetailControllerImpl) List(ctx *fiber.Ctx) error {
	saleId := ctx.Params("saleId")

	request := &web.ListSalesDetailRequest{
		SaleId: saleId,
	}

	responses, err := c.SalesDetailService.List(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to list sales details")
		return err
	}

	return ctx.JSON(web.WebResponse[[]web.SalesDetailResponse]{Data: responses})
}

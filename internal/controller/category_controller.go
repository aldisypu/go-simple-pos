package controller

import "github.com/gofiber/fiber/v2"

type CategoryController interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	Get(ctx *fiber.Ctx) error
	List(ctx *fiber.Ctx) error
}

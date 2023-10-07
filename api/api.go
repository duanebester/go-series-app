package api

import (
	"go-series-app/services"

	"github.com/gofiber/fiber/v2"
)

func NewApi(service services.Services) *fiber.App {
	api := fiber.New()

	api.Get("/report", func(c *fiber.Ctx) error {
		report := service.GetReport()
		return c.JSON(report)
	})

	return api
}

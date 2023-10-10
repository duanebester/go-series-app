package handlers

import (
	"go-series-app/services"

	"github.com/gofiber/fiber/v2"
)

func GetProductByID(productService services.ProductService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on get product request", "errors": "invalid product id"})
		}

		product, err := productService.GetProductByID(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
		}

		return c.JSON(fiber.Map{"status": "success", "data": product})
	}
}

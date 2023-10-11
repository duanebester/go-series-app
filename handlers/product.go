package handlers

import (
	"go-series-app/models"
	"go-series-app/services"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	GetProduct(ctx *fiber.Ctx) error
	CreateProduct(ctx *fiber.Ctx) error
	UpdateProduct(ctx *fiber.Ctx) error
	DeleteProduct(ctx *fiber.Ctx) error
}

type productHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) ProductHandler {
	return &productHandler{
		productService: productService,
	}
}

func (h *productHandler) GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on get product request", "errors": "invalid product id"})
	}

	product, err := h.productService.GetProductByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "data": product})
}

func (h *productHandler) CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create product", "data": err})
	}

	createdProduct, err := h.productService.CreateProduct(*product)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create product", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created product", "data": createdProduct})
}

func (h *productHandler) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on get product request", "errors": "invalid product id"})
	}

	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update product", "data": err})
	}

	updatedProduct, err := h.productService.UpdateProduct(id, *product)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update product", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Updated product", "data": updatedProduct})
}

func (h *productHandler) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on get product request", "errors": "invalid product id"})
	}

	_, err := h.productService.DeleteProduct(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete product", "data": err})
	}

	// if deleted.DeletedAt. == nil {
	// 	return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete product", "data": err})
	// }

	return c.JSON(fiber.Map{"status": "success", "message": "Deleted product"})
}

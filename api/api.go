package api

import (
	"go-series-app/handlers"
	"go-series-app/middleware"
	"go-series-app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewApi(service services.Services) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "go-series-app",
	})
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}  ${status} - ${latency} ${method} ${path}\n",
	}))

	api := app.Group("/api")

	// Auth
	auth := api.Group("/auth")
	authHandler := handlers.NewAuthHandler(service.UserService())
	auth.Post("/login", authHandler.Login)

	// User
	// user := api.Group("/user")
	// user.Get("/:id", handler.GetUser)
	// user.Post("/", handler.CreateUser)
	// user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
	// user.Delete("/:id", middleware.Protected(), handler.DeleteUser)

	// Product
	productHandler := handlers.NewProductHandler(service.ProductService())
	product := api.Group("/product", middleware.Protected())
	product.Post("/", productHandler.CreateProduct)
	product.Get("/:id", productHandler.GetProduct)
	product.Put("/:id", productHandler.UpdateProduct)
	product.Delete("/:id", productHandler.DeleteProduct)

	// Report
	// report := api.Group("/report", middleware.Protected())
	// report.Get("/", func(c *fiber.Ctx) error {
	// 	report := service.GetReport()
	// 	return c.JSON(report)
	// })

	return app
}

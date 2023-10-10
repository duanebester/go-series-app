package api

import (
	"go-series-app/handlers"
	"go-series-app/middleware"
	"go-series-app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	// "github.com/golang-jwt/jwt/v5"
)

func NewApi(service services.Services) *fiber.App {
	app := fiber.New()

	api := app.Group("/api", logger.New())

	// Auth
	auth := api.Group("/auth")
	// userService := services.NewUserService(configs.DB)
	// authHandler := handlers.NewAuthHandler(service.UserService())

	auth.Post("/login", handlers.Login(service.UserService()))
	// User
	// user := api.Group("/user")
	// user.Get("/:id", handler.GetUser)
	// user.Post("/", handler.CreateUser)
	// user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
	// user.Delete("/:id", middleware.Protected(), handler.DeleteUser)

	// Product
	product := api.Group("/product", middleware.Protected())
	product.Get("/:id", handlers.GetProductByID(service.ProductService()))

	// Report
	// report := api.Group("/report", middleware.Protected())
	// report.Get("/", func(c *fiber.Ctx) error {
	// 	report := service.GetReport()
	// 	return c.JSON(report)
	// })

	return app
}

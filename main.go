package main

import (
	"log"
	"soulcode-pre-test/domain"
	"soulcode-pre-test/handler"
	"soulcode-pre-test/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	productRepo := domain.NewMemoryProductRepository()
	productCache := domain.NewMemoryProductCache()
	productUsecase := usecase.NewProductUsecase(productRepo, productCache)
	productHandler := handler.NewProductHandler(productUsecase)

	productHandler.RegisterRoutes(app)
	handler.RegisterSwaggerRoute(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Product API is running",
		})
	})

	log.Fatal(app.Listen(":3000"))
}

package handler

import "github.com/gofiber/fiber/v2"

func RegisterSwaggerRoute(app *fiber.App) {
	app.Static("/swagger", "./swagger", fiber.Static{
		Index: "swagger.html",
	})
}

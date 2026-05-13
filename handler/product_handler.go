package handler

import (
	"errors"
	"soulcode-pre-test/domain"
	"soulcode-pre-test/usecase"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	usecase  *usecase.ProductUsecase
	validate *validator.Validate
}

func NewProductHandler(usecase *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		usecase:  usecase,
		validate: validator.New(),
	}
}

func (h *ProductHandler) RegisterRoutes(app *fiber.App) {
	app.Post("/products", h.Create)
	app.Get("/products", h.GetAll)
	app.Get("/products/:id", h.GetByID)
	app.Put("/products/:id", h.Update)
	app.Delete("/products/:id", h.Delete)
}

func (h *ProductHandler) Create(c *fiber.Ctx) error {
	var req domain.CreateProductRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request body",
		})
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "validation failed",
			"error":   err.Error(),
		})
	}

	product, err := h.usecase.Create(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create product",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "product created successfully",
		"data":    product,
	})
}

func (h *ProductHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	product, err := h.usecase.GetByID(id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "product not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to get product",
		})
	}

	return c.JSON(fiber.Map{
		"message": "product retrieved successfully",
		"data":    product,
	})
}

func (h *ProductHandler) GetAll(c *fiber.Ctx) error {
	products, err := h.usecase.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to get products",
		})
	}

	return c.JSON(fiber.Map{
		"message": "products retrieved successfully",
		"data":    products,
	})
}

func (h *ProductHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var req domain.UpdateProductRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request body",
		})
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "validation failed",
			"error":   err.Error(),
		})
	}

	product, err := h.usecase.Update(id, req)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "product not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to update product",
		})
	}

	return c.JSON(fiber.Map{
		"message": "product updated successfully",
		"data":    product,
	})
}

func (h *ProductHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.usecase.Delete(id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "product not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to delete product",
		})
	}

	return c.JSON(fiber.Map{
		"message": "product deleted successfully",
	})
}

package domain

import "time"

type Product struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Price     float64   `json:"price" validate:"required,gt=0"`
	Stock     int       `json:"stock" validate:"gte=0"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateProductRequest struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required,gt=0"`
	Stock int     `json:"stock" validate:"gte=0"`
}

type UpdateProductRequest struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required,gt=0"`
	Stock int     `json:"stock" validate:"gte=0"`
}

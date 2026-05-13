package domain

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("product not found")

type MemoryProductRepository struct {
	mu       sync.RWMutex
	products map[string]Product
}

func NewMemoryProductRepository() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[string]Product),
	}
}

func (r *MemoryProductRepository) Create(product Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.products[product.ID] = product
	return nil
}

func (r *MemoryProductRepository) GetByID(id string) (Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	product, exists := r.products[id]
	if !exists {
		return Product{}, ErrNotFound
	}
	return product, nil
}

func (r *MemoryProductRepository) GetAll() ([]Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	products := make([]Product, 0, len(r.products))
	for _, product := range r.products {
		products = append(products, product)
	}
	return products, nil
}

func (r *MemoryProductRepository) Update(product Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.products[product.ID]; !exists {
		return ErrNotFound
	}
	r.products[product.ID] = product
	return nil
}

func (r *MemoryProductRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.products[id]; !exists {
		return ErrNotFound
	}
	delete(r.products, id)
	return nil
}

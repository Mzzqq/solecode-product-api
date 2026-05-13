package domain

import "sync"

type MemoryProductCache struct {
	mu       sync.RWMutex
	products map[string]Product
}

func NewMemoryProductCache() *MemoryProductCache {
	return &MemoryProductCache{
		products: make(map[string]Product),
	}
}

func (c *MemoryProductCache) Get(id string) (Product, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	product, exists := c.products[id]

	return product, exists
}

func (c *MemoryProductCache) Set(product Product) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.products[product.ID] = product
}

func (c *MemoryProductCache) Delete(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.products, id)
}

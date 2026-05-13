package usecase

import (
	"soulcode-pre-test/domain"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	repo := domain.NewMemoryProductRepository()
	cache := domain.NewMemoryProductCache()
	uc := NewProductUsecase(repo, cache)

	tests := []struct {
		name    string
		req     domain.CreateProductRequest
		wantErr bool
	}{
		{
			name: "success create product",
			req: domain.CreateProductRequest{
				Name:  "Laptop",
				Price: 15000000,
				Stock: 10,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product, err := uc.Create(tt.req)

			if tt.wantErr && err == nil {
				t.Fatal("expected error but got nil")
			}

			if !tt.wantErr && err != nil {
				t.Fatalf("expected nil error but got %v", err)
			}

			if product.ID == "" {
				t.Fatal("expected product ID to be generated")
			}

			if product.Name != tt.req.Name {
				t.Fatalf("expected name %s but got %s", tt.req.Name, product.Name)
			}
		})
	}
}

func TestGetProductByID(t *testing.T) {
	repo := domain.NewMemoryProductRepository()
	cache := domain.NewMemoryProductCache()
	uc := NewProductUsecase(repo, cache)

	createdProduct, _ := uc.Create(domain.CreateProductRequest{
		Name:  "Mouse",
		Price: 250000,
		Stock: 20,
	})

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "success get product by id",
			id:      createdProduct.ID,
			wantErr: false,
		},
		{
			name:    "product not found",
			id:      "invalid-id",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product, err := uc.GetByID(tt.id)

			if tt.wantErr && err == nil {
				t.Fatal("expected error but got nil")
			}

			if !tt.wantErr && err != nil {
				t.Fatalf("expected nil error but got %v", err)
			}

			if !tt.wantErr && product.ID != tt.id {
				t.Fatalf("expected product ID %s but got %s", tt.id, product.ID)
			}
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	repo := domain.NewMemoryProductRepository()
	cache := domain.NewMemoryProductCache()
	uc := NewProductUsecase(repo, cache)

	createdProduct, _ := uc.Create(domain.CreateProductRequest{
		Name:  "Keyboard",
		Price: 500000,
		Stock: 15,
	})

	tests := []struct {
		name    string
		id      string
		req     domain.UpdateProductRequest
		wantErr bool
	}{
		{
			name: "success update product",
			id:   createdProduct.ID,
			req: domain.UpdateProductRequest{
				Name:  "Mechanical Keyboard",
				Price: 800000,
				Stock: 8,
			},
			wantErr: false,
		},
		{
			name: "product not found",
			id:   "invalid-id",
			req: domain.UpdateProductRequest{
				Name:  "Gaming Keyboard",
				Price: 900000,
				Stock: 5,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product, err := uc.Update(tt.id, tt.req)

			if tt.wantErr && err == nil {
				t.Fatal("expected error but got nil")
			}

			if !tt.wantErr && err != nil {
				t.Fatalf("expected nil error but got %v", err)
			}

			if !tt.wantErr && product.Name != tt.req.Name {
				t.Fatalf("expected product name %s but got %s", tt.req.Name, product.Name)
			}
		})
	}
}

func TestDeleteProduct(t *testing.T) {
	repo := domain.NewMemoryProductRepository()
	cache := domain.NewMemoryProductCache()
	uc := NewProductUsecase(repo, cache)

	createdProduct, _ := uc.Create(domain.CreateProductRequest{
		Name:  "Monitor",
		Price: 2000000,
		Stock: 3,
	})

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "success delete product",
			id:      createdProduct.ID,
			wantErr: false,
		},
		{
			name:    "product not found",
			id:      "invalid-id",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.Delete(tt.id)

			if tt.wantErr && err == nil {
				t.Fatal("expected error but got nil")
			}

			if !tt.wantErr && err != nil {
				t.Fatalf("expected nil error but got %v", err)
			}
		})
	}
}

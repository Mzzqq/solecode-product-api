package usecase

import (
	"soulcode-pre-test/domain"
	"time"

	"github.com/google/uuid"
)

type ProductUsecase struct {
	repo  domain.ProductRepository
	cache domain.ProductCache
}

func NewProductUsecase(repo domain.ProductRepository, cache domain.ProductCache) *ProductUsecase {
	return &ProductUsecase{
		repo:  repo,
		cache: cache,
	}
}

func (u *ProductUsecase) Create(req domain.CreateProductRequest) (domain.Product, error) {
	product := domain.Product{
		ID:        uuid.NewString(),
		Name:      req.Name,
		Price:     req.Price,
		Stock:     req.Stock,
		CreatedAt: time.Now(),
	}

	if err := u.repo.Create(product); err != nil {
		return domain.Product{}, err
	}

	u.cache.Set(product)

	return product, nil
}

func (u *ProductUsecase) GetByID(id string) (domain.Product, error) {
	if product, exists := u.cache.Get(id); exists {
		return product, nil
	}

	product, err := u.repo.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}

	u.cache.Set(product)

	return product, nil
}

func (u *ProductUsecase) GetAll() ([]domain.Product, error) {
	return u.repo.GetAll()
}

func (u *ProductUsecase) Update(id string, req domain.UpdateProductRequest) (domain.Product, error) {
	existingProduct, err := u.repo.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}

	updatedProduct := domain.Product{
		ID:        existingProduct.ID,
		Name:      req.Name,
		Price:     req.Price,
		Stock:     req.Stock,
		CreatedAt: existingProduct.CreatedAt,
	}

	if err := u.repo.Update(updatedProduct); err != nil {
		return domain.Product{}, err
	}

	u.cache.Set(updatedProduct)

	return updatedProduct, nil
}

func (u *ProductUsecase) Delete(id string) error {
	if err := u.repo.Delete(id); err != nil {
		return err
	}
	u.cache.Delete(id)

	return nil
}

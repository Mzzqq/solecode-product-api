package domain

type ProductRepository interface {
	Create(product Product) error
	GetByID(id string) (Product, error)
	GetAll() ([]Product, error)
	Update(product Product) error
	Delete(id string) error
}

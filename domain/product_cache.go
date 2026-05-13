package domain

type ProductCache interface {
	Get(id string) (Product, bool)
	Set(product Product)
	Delete(id string)
}

package application

import (
	"233338-R-C2/src/productos/domain"
	"233338-R-C2/src/productos/domain/entities"
)

type AddProduct struct {
	db domain.IProduct
}

func NewAddProduct(db domain.IProduct) *AddProduct {
	return &AddProduct{db: db}
}

func (ap *AddProduct) Execute(product *entities.Product) error {
	return ap.db.Guardar(product)
}

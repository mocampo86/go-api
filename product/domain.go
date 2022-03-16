package product

import (
	"fmt"
	"time"
)

//--- Product
type DAOProduct struct {
	ProductId   int
	Name        string
	Description string
	Price       float32
	SKU         string
	CreatedOn   time.Time
	UpdatedOn   time.Time
	DeletedOn   time.Time
}

type Product struct {
	ProductId   int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
}

func (p DAOProduct) ToModel() Product {
	return Product{
		ProductId:   p.ProductId,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		SKU:         p.SKU,
	}
}

func (p Product) ToDAO() DAOProduct {
	return DAOProduct{
		ProductId:   p.ProductId,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		SKU:         p.SKU,
	}
}

//---Error
var ErrFuncNotImplemented = fmt.Errorf("Not implemented function")

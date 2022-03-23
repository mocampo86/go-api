package product

import (
	"context"
)

// Repository access to data
type Repository interface {
	// Remove a product by ID from database.
	Remove(ctx context.Context, id int64) error
	// Find retrieves a product from database based on a product's id
	Find(ctx context.Context, filter ProductFilterRequest) (products []DAOProduct, err error)
	// GetAll retrieves all product from database as an array of product
	GetAll(ctx context.Context) (products []DAOProduct, err error)
	// Update changes product information on the Product.Id passed in.
	Update(ctx context.Context, product DAOProduct) error
	// Insert a product to a database. Return an error
	Insert(ctx context.Context, product DAOProduct) error
}

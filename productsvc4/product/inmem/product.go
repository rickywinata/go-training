package inmem

import (
	"context"

	"github.com/rickywinata/go-training/productsvc4/product"
)

// ProductRepository implements ProductRepository.
type ProductRepository struct {
	Data []*product.Product
}

// Insert inserts a new product.
func (r *ProductRepository) Insert(ctx context.Context, p *product.Product) error {
	r.Data = append(r.Data, p)
	return nil
}

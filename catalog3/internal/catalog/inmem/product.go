package inmem

import (
	"context"

	"github.com/rickywinata/go-training/catalog3/internal/catalog/model"
)

// ProductRepository implements ProductRepository.
type ProductRepository struct {
	Data []*model.Product
}

// Insert inserts a new product.
func (r *ProductRepository) Insert(ctx context.Context, p *model.Product) error {
	r.Data = append(r.Data, p)
	return nil
}

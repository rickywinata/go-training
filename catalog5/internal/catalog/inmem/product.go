package inmem

import (
	"context"

	"github.com/rickywinata/go-training/catalog5/internal/catalog"
)

// ProductRepository implements ProductRepository.
type ProductRepository struct {
	Data []*catalog.Product
}

// Insert inserts a new catalog.
func (r *ProductRepository) Insert(ctx context.Context, p *catalog.Product) error {
	r.Data = append(r.Data, p)
	return nil
}

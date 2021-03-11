package main

import "context"

// InmemProductRepository implements ProductRepository.
type InmemProductRepository struct {
	Data []*Product
}

// Insert inserts a new product.
func (r *InmemProductRepository) Insert(ctx context.Context, product *Product) error {
	r.Data = append(r.Data, product)
	return nil
}

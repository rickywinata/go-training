package main

import (
	"context"
)

type (
	// GetProductQuery represents the parameters for getting a product.
	GetProductQuery struct {
		Name string `json:"name"`
	}
)

// ProductView is an interface for product query operations.
type ProductView interface {
	GetProduct(ctx context.Context, q *GetProductQuery) (*Product, error)
}

type productView struct {
	productRepo *InmemProductRepository
}

// NewProductView creates a new product view.
func NewProductView(productRepo1 *InmemProductRepository) ProductView {
	return &productView{
		productRepo: productRepo1,
	}
}

func (s *productView) GetProduct(ctx context.Context, q *GetProductQuery) (*Product, error) {
	for _, d := range s.productRepo.Data {
		if d.Name == q.Name {
			return d, nil
		}
	}
	return nil, nil
}

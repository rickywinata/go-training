package view

import (
	"context"

	"github.com/rickywinata/go-training/catalog3/internal/catalog"
	"github.com/rickywinata/go-training/catalog3/internal/catalog/inmem"
)

type (
	// GetProductQuery represents the parameters for getting a product.
	GetProductQuery struct {
		Name string `json:"name"`
	}
)

// ProductView is an interface for product query operations.
type ProductView interface {
	GetProduct(ctx context.Context, q *GetProductQuery) (*catalog.Product, error)
}

type productView struct {
	productRepo *inmem.ProductRepository
}

// NewProductView creates a new product view.
func NewProductView(productRepo1 *inmem.ProductRepository) ProductView {
	return &productView{
		productRepo: productRepo1,
	}
}

func (s *productView) GetProduct(ctx context.Context, q *GetProductQuery) (*catalog.Product, error) {
	for _, d := range s.productRepo.Data {
		if d.Name == q.Name {
			return d, nil
		}
	}
	return nil, nil
}

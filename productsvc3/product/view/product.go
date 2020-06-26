package view

import (
	"context"

	"github.com/rickywinata/go-training/productsvc3/product"
	"github.com/rickywinata/go-training/productsvc3/product/inmem"
)

type (
	// GetProductQuery represents the parameters for getting a product.
	GetProductQuery struct {
		Name string `json:"name"`
	}
)

// ProductView is an interface for product query operations.
type ProductView interface {
	GetProduct(ctx context.Context, q *GetProductQuery) (*product.Product, error)
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

func (s *productView) GetProduct(ctx context.Context, q *GetProductQuery) (*product.Product, error) {
	for _, d := range s.productRepo.Data {
		if d.Name == q.Name {
			return d, nil
		}
	}
	return nil, nil
}

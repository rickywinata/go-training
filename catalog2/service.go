package main

import (
	"context"
)

type (
	// CreateProductInput represents the input for creating a product.
	CreateProductInput struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
	}
)

// ProductService is an interface for product operations.
type ProductService interface {
	CreateProduct(ctx context.Context, input *CreateProductInput) (*Product, error)
}

type productService struct {
	productRepo ProductRepository
}

// NewProductService creates a new product service.
func NewProductService(productRepo1 ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo1,
	}
}

func (s *productService) CreateProduct(ctx context.Context, input *CreateProductInput) (*Product, error) {
	product := &Product{
		Name:  input.Name,
		Price: input.Price,
	}

	if err := s.productRepo.Insert(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

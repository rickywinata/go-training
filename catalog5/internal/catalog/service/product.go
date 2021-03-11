package service

import (
	"context"

	"github.com/rickywinata/go-training/catalog5/internal/catalog"
	"gopkg.in/go-playground/validator.v9"
)

type (
	// CreateProductCommand represents the parameters for creating a catalog.
	CreateProductCommand struct {
		Name  string `json:"name" validate:"min=3"`
		Price int    `json:"price"`
	}
)

// ProductService is an interface for product operations.
type ProductService interface {
	CreateProduct(ctx context.Context, cmd *CreateProductCommand) (*catalog.Product, error)
}

type productService struct {
	productRepo catalog.Repository
}

// NewProductService creates a new product service.
func NewProductService(productRepo1 catalog.Repository) ProductService {
	return &productService{
		productRepo: productRepo1,
	}
}

func (s *productService) CreateProduct(ctx context.Context, cmd *CreateProductCommand) (*catalog.Product, error) {
	validate := validator.New()
	if err := validate.Struct(cmd); err != nil {
		return nil, err
	}

	p := &catalog.Product{
		Name:  cmd.Name,
		Price: cmd.Price,
	}

	if err := s.productRepo.Insert(ctx, p); err != nil {
		return nil, err
	}

	return p, nil
}

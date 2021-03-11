package catalog

import (
	"context"

	"github.com/rickywinata/go-training/catalog3/internal/catalog/model"
)

type (
	// CreateProductCommand represents the parameters for creating a catalog\.
	CreateProductCommand struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
	}
)

// Service is an interface for catalog use cases.
type Service interface {
	CreateProduct(ctx context.Context, cmd *CreateProductCommand) (*model.Product, error)
}

type service struct {
	productRepo model.Repository
}

// NewService creates a new product service.
func NewService(productRepo1 model.Repository) Service {
	return &service{
		productRepo: productRepo1,
	}
}

func (s *service) CreateProduct(ctx context.Context, cmd *CreateProductCommand) (*model.Product, error) {
	p := &model.Product{
		Name:  cmd.Name,
		Price: cmd.Price,
	}

	if err := s.productRepo.Insert(ctx, p); err != nil {
		return nil, err
	}

	return p, nil
}

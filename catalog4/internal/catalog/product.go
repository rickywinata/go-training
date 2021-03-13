package catalog

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rickywinata/go-training/catalog4/internal/catalog/model"
	"github.com/rickywinata/go-training/catalog4/internal/catalog/view"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type (
	// CreateProductInput represents the input for creating a product.
	CreateProductInput struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
	}

	// CreateProductOutput represents the output for creating a product.
	CreateProductOutput struct {
		*view.Product
	}
)

// Service is an interface for catalog use cases.
type Service interface {
	CreateProduct(ctx context.Context, input *CreateProductInput) (*CreateProductOutput, error)
}

type service struct {
	db *sqlx.DB
}

// NewService creates a new product service.
func NewService(db *sqlx.DB) Service {
	return &service{
		db: db,
	}
}

func (s *service) CreateProduct(ctx context.Context, input *CreateProductInput) (*CreateProductOutput, error) {
	p := &model.Product{
		Name:  input.Name,
		Price: null.IntFrom(input.Price),
	}

	err := p.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &CreateProductOutput{
		Product: &view.Product{
			Name:  p.Name,
			Price: p.Price.Int,
		},
	}, nil
}

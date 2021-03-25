package catalog

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rickywinata/go-training/catalog5/internal/catalog/database/model"
	"github.com/rickywinata/go-training/catalog5/internal/catalog/view"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type (
	CreateProductInput struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
	}

	CreateProductOutput struct {
		*view.Product
	}

	GetProductInput struct {
		Name string `json:"name"`
	}

	GetProductOutput struct {
		*view.Product
	}
)

type Service interface {
	CreateProduct(ctx context.Context, input *CreateProductInput) (*CreateProductOutput, error)
	GetProduct(ctx context.Context, input *GetProductInput) (*GetProductOutput, error)
}

type service struct {
	db *sqlx.DB
}

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
		&view.Product{
			Name:  p.Name,
			Price: p.Price.Int,
		},
	}, nil
}

func (s *service) GetProduct(ctx context.Context, input *GetProductInput) (*GetProductOutput, error) {
	p, err := view.GetProductByName(ctx, s.db, input.Name)
	return &GetProductOutput{p}, err
}

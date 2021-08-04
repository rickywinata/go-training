package main

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rickywinata/go-training/catalog2/querier"
)

type (
	CreateProductInput struct {
		Name  string `json:"name"`
		Price int64  `json:"price"`
	}

	GetProductInput struct {
		Name string `json:"name"`
	}
)

type ProductView struct {
	Name  string `json:"name" validate:"max=5"`
	Price int64  `json:"price"`
}

type ProductService interface {
	CreateProduct(ctx context.Context, input *CreateProductInput) (*ProductView, error)
	GetProduct(ctx context.Context, input *GetProductInput) (*ProductView, error)
}

type productService struct {
	db *sqlx.DB
	q  querier.Querier
}

func NewProductService(db *sqlx.DB) ProductService {
	return &productService{
		db: db,
		q:  querier.New(db),
	}
}

func (s *productService) CreateProduct(ctx context.Context, input *CreateProductInput) (*ProductView, error) {
	p, err := s.q.InsertProduct(ctx, querier.InsertProductParams{
		Name:  input.Name,
		Price: input.Price,
	})
	if err != nil {
		return nil, err
	}

	return &ProductView{
		Name:  p.Name,
		Price: p.Price,
	}, nil
}

func (s *productService) GetProduct(ctx context.Context, input *GetProductInput) (*ProductView, error) {
	p, err := s.q.FindProduct(ctx, input.Name)
	if err != nil {
		return nil, err
	}

	return &ProductView{
		Name:  p.Name,
		Price: p.Price,
	}, err
}

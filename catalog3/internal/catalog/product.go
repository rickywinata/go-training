package catalog

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/rickywinata/go-training/catalog3/internal/database/querier"
)

var (
	ErrNotFound = errors.New("product is not found")
)

type (
	CreateProductInput struct {
		Name  string `json:"name"`
		Price int64  `json:"price"`
	}

	CreateProductOutput struct {
		*ProductView
	}

	GetProductInput struct {
		Name string `json:"name"`
	}

	GetProductOutput struct {
		*ProductView
	}
)

type ProductView struct {
	Name  string `json:"name" validate:"max=5"`
	Price int64  `json:"price"`
}

type Service interface {
	CreateProduct(ctx context.Context, input *CreateProductInput) (*CreateProductOutput, error)
	GetProduct(ctx context.Context, input *GetProductInput) (*GetProductOutput, error)
}

type service struct {
	db *sqlx.DB
	q  querier.Querier
	sb squirrel.StatementBuilderType
}

func NewService(db *sqlx.DB) Service {
	return &service{
		db: db,
		q:  querier.New(db),
		sb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (s *service) CreateProduct(ctx context.Context, input *CreateProductInput) (*CreateProductOutput, error) {
	p, err := s.q.InsertProduct(ctx, querier.InsertProductParams{
		Name:  input.Name,
		Price: input.Price,
	})
	if err != nil {
		return nil, err
	}

	return &CreateProductOutput{
		&ProductView{
			Name:  p.Name,
			Price: p.Price,
		},
	}, nil
}

func (s *service) GetProduct(ctx context.Context, input *GetProductInput) (*GetProductOutput, error) {
	q := s.sb.Select("name", "price").
		From("product").
		Where(squirrel.Eq{"name": input.Name})
		
	sql, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	var res ProductView
	err = s.db.Get(&res, sql, args...)
	if err != nil {
		return nil, err
	}

	return &GetProductOutput{
		ProductView: &res,
	}, err
}

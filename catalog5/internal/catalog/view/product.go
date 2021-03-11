package view

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/rickywinata/go-training/catalog5/internal/database/model"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// List of errors.
var (
	ErrNotFound = errors.New("product is not found")
)

type (
	// GetProductQuery represents the parameters for getting a product.
	GetProductQuery struct {
		Name string `json:"name"`
	}
)

// Product is a view-only representation of model.Product.
type Product struct {
	Name  string `json:"name" validate:"max=5"`
	Price int    `json:"price"`
}

// ProductView is an interface for product query operations.
type ProductView interface {
	GetProduct(ctx context.Context, q *GetProductQuery) (*Product, error)
}

type productView struct {
	db *sqlx.DB
}

// NewProductView creates a new product view.
func NewProductView(db *sqlx.DB) ProductView {
	return &productView{
		db: db,
	}
}

func (v *productView) GetProduct(ctx context.Context, q *GetProductQuery) (*Product, error) {
	mproduct, err := model.Products(qm.Where("name = ?", q.Name)).One(ctx, v.db)
	if err != nil {
		return nil, err
	}

	return &Product{
		Name:  mproduct.Name,
		Price: mproduct.Price.Int,
	}, nil
}

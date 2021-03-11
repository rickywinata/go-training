package view

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rickywinata/go-training/catalog5/internal/catalog"
	"github.com/rickywinata/go-training/catalog5/internal/postgres/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// // List of errors.
var (
	ErrNotFound = errors.New("product is not found")
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
	db *sqlx.DB
}

// NewProductView creates a new product view.
func NewProductView(db *sqlx.DB) ProductView {
	return &productView{
		db: db,
	}
}

func (v *productView) GetProduct(ctx context.Context, q *GetProductQuery) (*catalog.Product, error) {
	mproduct, err := models.Products(qm.Where("name = ?", q.Name)).One(ctx, v.db)
	if err != nil {
		return nil, err
	}

	return &catalog.Product{
		Name:  mproduct.Name,
		Price: mproduct.Price.Int,
	}, nil
}

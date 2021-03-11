package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rickywinata/go-training/catalog5/internal/catalog"
	"github.com/rickywinata/go-training/catalog5/internal/postgres/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// ProductRepository implements ProductRepository.
type ProductRepository struct {
	db *sqlx.DB
}

// NewProductRepository creates a new repository.
func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

// Insert inserts a new product.
func (r *ProductRepository) Insert(ctx context.Context, p *catalog.Product) error {
	dp := &models.Product{
		Name:  p.Name,
		Price: null.IntFrom(p.Price),
	}

	err := dp.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

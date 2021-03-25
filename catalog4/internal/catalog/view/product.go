package view

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/rickywinata/go-training/catalog4/internal/catalog/database/model"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var (
	ErrNotFound = errors.New("product is not found")
)

type Product struct {
	Name  string `json:"name" validate:"max=5"`
	Price int    `json:"price"`
}

func GetProductByName(ctx context.Context, db *sqlx.DB, name string) (*Product, error) {
	mproduct, err := model.Products(qm.Where("name = ?", name)).One(ctx, db)
	if err != nil {
		return nil, err
	}

	return &Product{
		Name:  mproduct.Name,
		Price: mproduct.Price.Int,
	}, nil
}

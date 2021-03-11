package model

import "context"

// Product represents an item to be sold.
type Product struct {
	Name  string `json:"name" validate:"max=5"`
	Price int    `json:"price"`
}

// Repository is an interface for product storage functions.
type Repository interface {
	Insert(ctx context.Context, product *Product) error
}

package main

import "context"

// Product represents an item to be sold.
type Product struct {
	Name  string `json:"name" validate:"max=5"`
	Price int    `json:"price"`
}

// ProductRepository is an interface for product storage functions.
type ProductRepository interface {
	Insert(ctx context.Context, product *Product) error
}

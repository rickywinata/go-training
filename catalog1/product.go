package main

// Product represents an item to be sold.
type Product struct {
	Name  string `json:"name" validate:"max=5"`
	Price int    `json:"price"`
}

package main

// Product .
type Product struct {
	Name  string `json:"name" validate:"max=5"`
	Price int    `json:"price"`
}

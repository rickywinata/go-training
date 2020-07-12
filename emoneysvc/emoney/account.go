package emoney

import (
	"context"
)

// Account represents an item to be sold.
type Account struct {
	ID      string `json:"id"`
	Balance int    `json:"balance"`
}

// Repository is an interface for product storage functions.
type Repository interface {
	Insert(ctx context.Context, account *Account) error
}

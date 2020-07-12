package emoney

import (
	"context"
)

// Account represents an item to be sold.
type Account struct {
	ID      string `json:"id"`
	Balance int    `json:"balance"`
}

// AccountEntry represents an entry on an account.
type AccountEntry struct {
	AccountID string
	Amount    int
}

// AccountRepository is an interface for product storage functions.
type AccountRepository interface {
	Insert(ctx context.Context, account *Account) error
}

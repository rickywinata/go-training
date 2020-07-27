package emoney

import (
	"context"
	"errors"
	"time"
)

// List of errors.
var (
	ErrAccountNotFound = errors.New("account is not found")
)

// Account represents an item to be sold.
type Account struct {
	ID         string `json:"id"`
	Balance    int    `json:"balance"`
	Registered bool   `json:"registered"`
}

// AccountEntry represents an entry on an account.
type AccountEntry struct {
	ID        string    `json:"id"`
	AccountID string    `json:"account_id"`
	Amount    int       `json:"amount"`
	BookedAt  time.Time `json:"booked_at"`
}

// AccountRepository is an interface for account storage functions.
type AccountRepository interface {
	FindByID(ctx context.Context, accID string) (*Account, error)
	Insert(ctx context.Context, acc *Account) error
	Update(ctx context.Context, acc *Account) error
}

// AccountEntryRepository is an interface for account entry storage operations.
type AccountEntryRepository interface {
	Insert(ctx context.Context, accEntry *AccountEntry) error
}

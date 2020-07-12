package inmem

import (
	"context"

	"github.com/rickywinata/go-training/emoneysvc/emoney"
)

// AccountRepository implements emoney.AccountRepository.
type AccountRepository struct {
	Data []*emoney.Account
}

// Insert inserts a new account.
func (r *AccountRepository) Insert(ctx context.Context, acc *emoney.Account) error {
	r.Data = append(r.Data, acc)
	return nil
}

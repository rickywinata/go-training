package inmem

import (
	"context"

	"github.com/rickywinata/go-training/emoneysvc/emoney"
)

// AccountRepository implements emoney.AccountRepository.
type AccountRepository struct {
	Data []*emoney.Account
}

// FindByID finds an account by its id.
func (r *AccountRepository) FindByID(ctx context.Context, accID string) (*emoney.Account, error) {
	for _, d := range r.Data {
		if d.ID == accID {
			return d, nil
		}
	}
	return nil, emoney.ErrAccountNotFound
}

// Insert inserts a new account.
func (r *AccountRepository) Insert(ctx context.Context, acc *emoney.Account) error {
	r.Data = append(r.Data, acc)
	return nil
}

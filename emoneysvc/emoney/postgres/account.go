package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rickywinata/go-training/emoneysvc/emoney"
	"github.com/rickywinata/go-training/emoneysvc/postgres/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// AccountRepository implements emoney.AccountRepository.
type AccountRepository struct {
	db *sqlx.DB
}

// NewAccountRepository creates a new repository.
func NewAccountRepository(db *sqlx.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

// Insert inserts a new account.
func (r *AccountRepository) Insert(ctx context.Context, acc *emoney.Account) error {
	dacc := &models.Account{
		ID:      acc.ID,
		Balance: null.IntFrom(acc.Balance),
	}

	err := dacc.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

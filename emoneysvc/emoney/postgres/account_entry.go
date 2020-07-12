package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rickywinata/go-training/emoneysvc/emoney"
	"github.com/rickywinata/go-training/emoneysvc/postgres/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// AccountEntryRepository implements emoney.AccountEntryRepository.
type AccountEntryRepository struct {
	db *sqlx.DB
}

// NewAccountEntryRepository creates a new repository.
func NewAccountEntryRepository(db *sqlx.DB) *AccountEntryRepository {
	return &AccountEntryRepository{
		db: db,
	}
}

// Insert inserts a new account entry.
func (r *AccountEntryRepository) Insert(ctx context.Context, accEntry *emoney.AccountEntry) error {
	maccEntry := &models.AccountEntry{
		ID:        accEntry.ID,
		AccountID: null.StringFrom(accEntry.AccountID),
		Amount:    null.IntFrom(accEntry.Amount),
		BookedAt:  null.TimeFrom(accEntry.BookedAt),
	}

	if err := maccEntry.Insert(ctx, r.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}

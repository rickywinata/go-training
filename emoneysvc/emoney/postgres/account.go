package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rickywinata/go-training/emoneysvc/emoney"
	"github.com/rickywinata/go-training/emoneysvc/postgres/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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

// FindByID finds an account by its id.
func (r *AccountRepository) FindByID(ctx context.Context, accID string) (*emoney.Account, error) {
	macc, err := models.Accounts(qm.Where("id = ?", accID)).One(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return &emoney.Account{
		ID:      macc.ID,
		Balance: macc.Balance.Int,
	}, nil
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

// Update updates an account.
func (r *AccountRepository) Update(ctx context.Context, acc *emoney.Account) error {
	macc, err := models.Accounts(qm.Where("id = ?", acc.ID)).One(ctx, r.db)
	if err != nil {
		return err
	}

	macc.Balance = null.IntFrom(acc.Balance)

	if _, err := macc.Update(ctx, r.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}

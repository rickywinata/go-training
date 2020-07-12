package view

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rickywinata/go-training/emoneysvc/emoney"
	"github.com/rickywinata/go-training/emoneysvc/postgres/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// // List of errors.
var (
	ErrNotFound = errors.New("account is not found")
)

type (
	// GetAccountQuery represents the parameters for getting a account.
	GetAccountQuery struct {
		AccountID string
	}
)

// AccountView is an interface for account query operations.
type AccountView interface {
	GetAccount(ctx context.Context, q *GetAccountQuery) (*emoney.Account, error)
}

type accountView struct {
	db *sqlx.DB
}

// NewAccountView creates a new account view.
func NewAccountView(db *sqlx.DB) AccountView {
	return &accountView{
		db: db,
	}
}

func (v *accountView) GetAccount(ctx context.Context, q *GetAccountQuery) (*emoney.Account, error) {
	macc, err := models.Accounts(qm.Where("id = ?", q.AccountID)).One(ctx, v.db)
	if err != nil {
		return nil, err
	}

	return &emoney.Account{
		ID:      macc.ID,
		Balance: macc.Balance.Int,
	}, nil
}

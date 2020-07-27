package emoney

import (
	"errors"

	"github.com/google/uuid"
)

const maxBalance = 2000000

// List of errors.
var (
	ErrMaxBalance       = errors.New("balance is at max limit")
	ErrBalanceNotEnough = errors.New("balance is not enough")
)

// Trx represents the created accounting entries when a transaction happens.
type Trx struct {
	Entries []*AccountEntry
}

// Topup deposits money to an account.
//
// It creates a Trx with one entry:
// - Add some amount to an account.
//
func Topup(acc *Account, amount int) (*Trx, error) {
	if acc.Balance+amount > maxBalance {
		return nil, ErrMaxBalance
	}

	acc.Balance += amount

	trx := &Trx{
		Entries: []*AccountEntry{
			{
				ID:        uuid.New().String(),
				AccountID: acc.ID,
				Amount:    amount,
			},
		},
	}

	return trx, nil
}

// Transfer sends money from one account to another account.
//
// It creates a Trx with 2 entries:
// - Remove some amount from 1 account.
// - Add some amount to another account.
//
func Transfer(accFrom *Account, accTo *Account, amount int) (*Trx, error) {
	if accFrom.Balance-amount < 0 {
		return nil, ErrBalanceNotEnough
	}

	if accTo.Balance+amount > maxBalance {
		return nil, ErrMaxBalance
	}

	accFrom.Balance -= amount
	accTo.Balance += amount

	trx := &Trx{
		Entries: []*AccountEntry{
			{
				ID:        uuid.New().String(),
				AccountID: accFrom.ID,
				Amount:    -amount,
			},
			{
				ID:        uuid.New().String(),
				AccountID: accTo.ID,
				Amount:    amount,
			},
		},
	}

	return trx, nil
}

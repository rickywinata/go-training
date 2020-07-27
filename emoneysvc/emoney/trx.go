package emoney

import (
	"errors"

	"github.com/google/uuid"
)

const maxUnregisteredBalance = 2000000
const maxRegisteredBalance = 10000000

// List of errors.
var (
	ErrMaxBalance       = errors.New("balance is at max limit")
	ErrBalanceNotEnough = errors.New("balance is not enough")
)

// Trx represents a money movement from one account to another account.
type Trx struct {
	Entries []*AccountEntry
}

// Topup deposits money to an account.
//
// It creates a Trx with one entry:
// - Add some amount to an account.
//
func Topup(acc *Account, amount int) (*Trx, error) {
	maxBalance := maxUnregisteredBalance
	if acc.Registered {
		maxBalance = maxRegisteredBalance
	}

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
	maxBalance := maxUnregisteredBalance
	if accTo.Registered {
		maxBalance = maxRegisteredBalance
	}

	if accTo.Balance+amount > maxBalance {
		return nil, ErrMaxBalance
	}

	if accFrom.Balance-amount < 0 {
		return nil, ErrBalanceNotEnough
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

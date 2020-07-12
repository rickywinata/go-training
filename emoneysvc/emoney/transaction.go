package emoney

import (
	"github.com/google/uuid"
)

// Transaction represents entries.
type Transaction struct {
	Entries []*AccountEntry
}

// Topup deposits money to an account.
//
// It creates a transaction with one entry:
// - Add some amount to an account.
//
func Topup(acc *Account, amount int) *Transaction {
	acc.Balance += amount

	trx := &Transaction{
		Entries: []*AccountEntry{
			{
				ID:        uuid.New().String(),
				AccountID: acc.ID,
				Amount:    amount,
			},
		},
	}

	return trx
}

// Transfer sends money from one account to another account.
//
// It creates a transaction with 2 entries:
// - Remove some amount from 1 account.
// - Add some amount to another account.
//
func Transfer(accFrom *Account, accTo *Account, amount int) *Transaction {
	accFrom.Balance -= amount
	accTo.Balance += amount

	trx := &Transaction{
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

	return trx
}

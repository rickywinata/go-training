package emoney

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestTopupSuccess(t *testing.T) {
	tests := map[string]struct {
		acc         *Account
		amount      int
		wantBalance int
		wantTrx     *Trx
	}{
		"success": {
			acc:         &Account{ID: "1"},
			amount:      10000,
			wantBalance: 10000,
			wantTrx: &Trx{
				Entries: []*AccountEntry{
					{
						AccountID: "1",
						Amount:    10000,
					},
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			c := qt.New(t)
			trx, err := Topup(tc.acc, tc.amount)
			if err != nil {
				t.Fatal(err)
			}

			c.Assert(tc.acc.Balance, qt.Equals, tc.wantBalance)
			c.Assert(trx, qt.CmpEquals(cmpopts.IgnoreFields(AccountEntry{}, "ID")), tc.wantTrx)
		})
	}
}

func TestTopupFailed(t *testing.T) {
	tests := map[string]struct {
		acc    *Account
		amount int
		want   error
	}{
		"balance already at max limit": {
			acc:    &Account{ID: "1", Balance: 2000000},
			amount: 10000,
			want:   ErrMaxBalance,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			c := qt.New(t)
			_, err := Topup(tc.acc, tc.amount)
			c.Assert(err, qt.Equals, tc.want)
		})
	}
}

func TestTransferSuccess(t *testing.T) {
	tests := map[string]struct {
		acc1            *Account
		acc2            *Account
		amount          int
		wantAcc1Balance int
		wantAcc2Balance int
		wantTrx         *Trx
	}{
		"success": {
			acc1:            &Account{ID: "1", Balance: 15000},
			acc2:            &Account{ID: "2", Balance: 0},
			amount:          10000,
			wantAcc1Balance: 5000,
			wantAcc2Balance: 10000,
			wantTrx: &Trx{
				Entries: []*AccountEntry{
					{
						AccountID: "1",
						Amount:    -10000,
					},
					{
						AccountID: "2",
						Amount:    10000,
					},
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			c := qt.New(t)
			trx, err := Transfer(tc.acc1, tc.acc2, tc.amount)
			if err != nil {
				t.Fatal(err)
			}

			c.Assert(tc.acc1.Balance, qt.Equals, tc.wantAcc1Balance)
			c.Assert(tc.acc2.Balance, qt.Equals, tc.wantAcc2Balance)
			c.Assert(trx, qt.CmpEquals(cmpopts.IgnoreFields(AccountEntry{}, "ID")), tc.wantTrx)
		})
	}

}

func TestTransferFailed(t *testing.T) {
	tests := map[string]struct {
		acc1   *Account
		acc2   *Account
		amount int
		want   error
	}{
		"balance acc1 is not enough": {
			acc1:   &Account{ID: "1", Balance: 5000},
			acc2:   &Account{ID: "1", Balance: 2000000},
			amount: 10000,
			want:   ErrBalanceNotEnough,
		},
		"balance acc2 already at max limit": {
			acc1:   &Account{ID: "1", Balance: 15000},
			acc2:   &Account{ID: "1", Balance: 2000000},
			amount: 10000,
			want:   ErrMaxBalance,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			c := qt.New(t)
			_, err := Transfer(tc.acc1, tc.acc2, tc.amount)
			c.Assert(err, qt.Equals, tc.want)
		})
	}

}

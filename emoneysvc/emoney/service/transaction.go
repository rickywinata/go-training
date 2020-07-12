package service

import (
	"context"

	"github.com/rickywinata/go-training/emoneysvc/emoney"
)

type (
	// TopupCommand represents the parameters for topping emoney.
	TopupCommand struct {
		AccountID string `json:"account_id"`
		Amount    int    `json:"amount"`
	}

	// TransferCommand transfers
	TransferCommand struct {
		FromAccountID string `json:"from_account_id"`
		ToAccountID   string `json:"to_account_id"`
		Amount        int    `json:"amount"`
	}
)

// TransactionService is an interface for transaction operations.
type TransactionService interface {
	Topup(ctx context.Context, cmd *TopupCommand) error
	Transfer(ctx context.Context, cmd *TransferCommand) error
}

type transactionService struct {
	accRepo      emoney.AccountRepository
	accEntryRepo emoney.AccountEntryRepository
}

// NewTransactionService creates a new transaction service.
func NewTransactionService(
	accRepo emoney.AccountRepository,
	accEntryRepo emoney.AccountEntryRepository,
) TransactionService {
	return &transactionService{accRepo, accEntryRepo}
}

// Topup tops up an account with some amount.
func (s *transactionService) Topup(ctx context.Context, cmd *TopupCommand) error {
	acc, err := s.accRepo.FindByID(ctx, cmd.AccountID)
	if err != nil {
		return err
	}

	trx := emoney.Topup(acc, cmd.Amount)

	if err := s.accRepo.Update(ctx, acc); err != nil {
		return err
	}

	for _, entry := range trx.Entries {
		if err := s.accEntryRepo.Insert(ctx, entry); err != nil {
			return err
		}
	}

	return nil
}

// Transfer transfers some amount from one account to another account.
func (s *transactionService) Transfer(ctx context.Context, cmd *TransferCommand) error {
	fromAcc, err := s.accRepo.FindByID(ctx, cmd.FromAccountID)
	if err != nil {
		return err
	}

	toAcc, err := s.accRepo.FindByID(ctx, cmd.ToAccountID)
	if err != nil {
		return err
	}

	trx := emoney.Transfer(fromAcc, toAcc, cmd.Amount)

	if err := s.accRepo.Update(ctx, fromAcc); err != nil {
		return err
	}

	if err := s.accRepo.Update(ctx, toAcc); err != nil {
		return err
	}

	for _, entry := range trx.Entries {
		if err := s.accEntryRepo.Insert(ctx, entry); err != nil {
			return err
		}
	}

	return nil
}

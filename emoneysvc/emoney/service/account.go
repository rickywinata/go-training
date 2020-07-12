package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/rickywinata/go-training/emoneysvc/emoney"
	"gopkg.in/go-playground/validator.v9"
)

type (
	// CreateAccountCommand represents the parameters for creating a emoney.
	CreateAccountCommand struct{}
)

// AccountService is an interface for account operations.
type AccountService interface {
	CreateAccount(ctx context.Context, cmd *CreateAccountCommand) (*emoney.Account, error)
}

type accountService struct {
	accountRepo emoney.Repository
}

// NewAccountService creates a new account service.
func NewAccountService(accountRepo1 emoney.Repository) AccountService {
	return &accountService{
		accountRepo: accountRepo1,
	}
}

func (s *accountService) CreateAccount(ctx context.Context, cmd *CreateAccountCommand) (*emoney.Account, error) {
	validate := validator.New()
	if err := validate.Struct(cmd); err != nil {
		return nil, err
	}

	acc := &emoney.Account{
		ID:      uuid.New().String(),
		Balance: 0,
	}

	if err := s.accountRepo.Insert(ctx, acc); err != nil {
		return nil, err
	}

	return acc, nil
}

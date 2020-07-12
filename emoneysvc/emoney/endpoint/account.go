package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/rickywinata/go-training/emoneysvc/emoney/service"
	"github.com/rickywinata/go-training/emoneysvc/emoney/view"
)

// GetAccount creates endpoint for getting an account.
func GetAccount(accountView view.AccountView) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		qry := request.(*view.GetAccountQuery)
		acc, err := accountView.GetAccount(ctx, qry)
		return acc, err
	}
}

// CreateAccount creates endpoint for creating an account.
func CreateAccount(svc service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		cmd := request.(*service.CreateAccountCommand)
		acc, err := svc.CreateAccount(ctx, cmd)
		return acc, err
	}
}

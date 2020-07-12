package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/rickywinata/go-training/emoneysvc/emoney/service"
)

// Topup creates endpoint for topup operation.
func Topup(svc service.TransactionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		cmd := request.(*service.TopupCommand)
		err := svc.Topup(ctx, cmd)
		return nil, err
	}
}

// Transfer creates endpoint for transfer operation.
func Transfer(svc service.TransactionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		cmd := request.(*service.TransferCommand)
		err := svc.Transfer(ctx, cmd)
		return nil, err
	}
}

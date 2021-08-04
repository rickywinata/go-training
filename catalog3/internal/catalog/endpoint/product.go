package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/rickywinata/go-training/catalog3/internal/catalog"
)

func GetProduct(svc catalog.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		qry := request.(*catalog.GetProductInput)
		product, err := svc.GetProduct(ctx, qry)
		return product, err
	}
}

func CreateProduct(svc catalog.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(*catalog.CreateProductInput)
		product, err := svc.CreateProduct(ctx, input)
		return product, err
	}
}

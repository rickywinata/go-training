package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/rickywinata/go-training/catalog3/internal/catalog"
	"github.com/rickywinata/go-training/catalog3/internal/catalog/view"
)

// GetProduct creates endpoint.
func GetProduct(productView view.ProductView) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		qry := request.(*view.GetProductQuery)
		product, err := productView.GetProduct(ctx, qry)
		return product, err
	}
}

// CreateProduct creates endpoint.
func CreateProduct(svc catalog.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		cmd := request.(*catalog.CreateProductCommand)
		product, err := svc.CreateProduct(ctx, cmd)
		return product, err
	}
}

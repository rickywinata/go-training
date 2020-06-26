package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/rickywinata/go-training/productsvc3/product/endpoint"
	"github.com/rickywinata/go-training/productsvc3/product/service"
	"github.com/rickywinata/go-training/productsvc3/product/view"
)

// GetProduct creates http handler.
func GetProduct(productView view.ProductView) http.Handler {
	return httptransport.NewServer(
		// Endpoint.
		endpoint.GetProduct(productView),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var qry view.GetProductQuery
			qry.Name = chi.URLParam(r, "product_name")
			return &qry, nil
		},

		// Encoder.
		encodeResponse,
	)
}

// CreateProduct creates http handler.
func CreateProduct(svc service.ProductService) http.Handler {
	return httptransport.NewServer(
		// Endpoint.
		endpoint.CreateProduct(svc),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var cmd service.CreateProductCommand
			if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
				return nil, err
			}
			return &cmd, nil
		},

		// Encoder.
		encodeResponse,
	)
}

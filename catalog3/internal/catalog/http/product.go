package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/rickywinata/go-training/catalog3/internal/catalog"
	"github.com/rickywinata/go-training/catalog3/internal/catalog/endpoint"
	"github.com/rickywinata/go-training/catalog3/internal/catalog/view"
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
func CreateProduct(svc catalog.Service) http.Handler {
	return httptransport.NewServer(
		// Endpoint.
		endpoint.CreateProduct(svc),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var input catalog.CreateProductInput
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				return nil, err
			}
			return &input, nil
		},

		// Encoder.
		encodeResponse,
	)
}

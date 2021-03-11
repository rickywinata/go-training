package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	repo := &InmemProductRepository{
		Data: []*Product{},
	}
	svc := NewProductService(repo)
	view := NewProductView(repo)

	createProductHandler := httptransport.NewServer(
		// Endpoint.
		func(ctx context.Context, request interface{}) (interface{}, error) {
			input := request.(*CreateProductInput)
			product, err := svc.CreateProduct(ctx, input)
			return product, err
		},

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var input CreateProductInput
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				return nil, err
			}
			return &input, nil
		},

		// Encoder.
		encodeResponse,
	)

	getProductHandler := httptransport.NewServer(
		// Endpoint.
		func(ctx context.Context, request interface{}) (interface{}, error) {
			qry := request.(*GetProductQuery)
			product, err := view.GetProduct(ctx, qry)
			return product, err
		},

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var qry GetProductQuery
			qry.Name = chi.URLParam(r, "product_name")
			return &qry, nil
		},

		// Encoder.
		encodeResponse,
	)

	r := chi.NewRouter()
	r.Post("/products", createProductHandler.ServeHTTP)
	r.Get("/products/{product_name}", getProductHandler.ServeHTTP)

	log.Println("Listening on :8080 ...")
	http.ListenAndServe(":8080", r)
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

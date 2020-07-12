package http

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/rickywinata/go-training/emoneysvc/emoney/endpoint"
	"github.com/rickywinata/go-training/emoneysvc/emoney/service"
)

// Topup creates http handler for topup operation.
func Topup(svc service.TransactionService) http.Handler {
	return httptransport.NewServer(
		// Endpoint.
		endpoint.Topup(svc),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var cmd service.TopupCommand
			if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
				return nil, err
			}
			return &cmd, nil
		},

		// Encoder.
		encodeResponse,

		// Error encoder.
		httptransport.ServerErrorEncoder(errorEncoder),
	)
}

// Transfer creates http handler for transfer operation.
func Transfer(svc service.TransactionService) http.Handler {
	return httptransport.NewServer(
		// Endpoint.
		endpoint.Transfer(svc),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var cmd service.TransferCommand
			if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
				return nil, err
			}
			return &cmd, nil
		},

		// Encoder.
		encodeResponse,

		// Error encoder.
		httptransport.ServerErrorEncoder(errorEncoder),
	)
}

package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/rickywinata/go-training/emoneysvc/emoney/endpoint"
	"github.com/rickywinata/go-training/emoneysvc/emoney/service"
	"github.com/rickywinata/go-training/emoneysvc/emoney/view"
)

// GetAccount creates http handler for getting an account.
func GetAccount(accountView view.AccountView) http.Handler {
	return httptransport.NewServer(
		// Endpoint.
		endpoint.GetAccount(accountView),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var qry view.GetAccountQuery
			qry.AccountID = chi.URLParam(r, "account_id")
			return &qry, nil
		},

		// Encoder.
		encodeResponse,

		// Error encoder.
		httptransport.ServerErrorEncoder(errorEncoder),
	)
}

// CreateAccount creates http handler for creating an account.
func CreateAccount(svc service.AccountService) http.Handler {
	return httptransport.NewServer(
		// Endpoint.
		endpoint.CreateAccount(svc),

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var cmd service.CreateAccountCommand
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

package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/rickywinata/go-training/emoneysvc/emoney"
	"github.com/rickywinata/go-training/productsvc5/product/view"
)

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	statusCode := err2code(err)
	w.WriteHeader(statusCode)

	if statusCode == http.StatusInternalServerError {
		log.Printf("system error: %+v\n", err)
		json.NewEncoder(w).Encode(errorWrapper{Error: "internal server error"})
		return
	}

	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

func err2code(err error) int {
	if err == view.ErrNotFound {
		return http.StatusNotFound
	}
	if err == emoney.ErrBalanceNotEnough {
		return http.StatusBadRequest
	}
	if err == emoney.ErrMaxBalance {
		return http.StatusBadRequest
	}
	if err == emoney.ErrAccountNotFound {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}

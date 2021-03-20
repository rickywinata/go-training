package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/rickywinata/go-training/catalog5/internal/catalog/view"
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
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}

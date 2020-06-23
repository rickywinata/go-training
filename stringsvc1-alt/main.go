package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
)

// StringService provides operations on strings.
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty string")

type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

func main() {
	svc := stringService{}

	uppercaseHandler := httptransport.NewServer(
		func(_ context.Context, request interface{}) (interface{}, error) {
			req := request.(uppercaseRequest)
			v, err := svc.Uppercase(req.S)
			if err != nil {
				return uppercaseResponse{v, err.Error()}, nil
			}
			return uppercaseResponse{v, ""}, nil
		},
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var request uppercaseRequest
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				return nil, err
			}
			return request, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	countHandler := httptransport.NewServer(
		func(_ context.Context, request interface{}) (interface{}, error) {
			req := request.(countRequest)
			v := svc.Count(req.S)
			return countResponse{v}, nil
		},
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var request countRequest
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				return nil, err
			}
			return request, nil
		},
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
	)

	r := chi.NewRouter()
	r.Post("/uppercase", uppercaseHandler.ServeHTTP)
	r.Post("/count", countHandler.ServeHTTP)

	log.Println("Listening on :8080 ...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

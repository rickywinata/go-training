package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	r := chi.NewRouter()
	r.Get("/greet", handle)
	log.Println("Listening on :8080 ...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

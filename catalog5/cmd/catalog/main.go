package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rickywinata/go-training/catalog5/internal/catalog"
	cataloghttp "github.com/rickywinata/go-training/catalog5/internal/catalog/http"
)

func main() {
	db, err := sqlx.Open("postgres", "dbname=catalog user=postgres password=password sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	catalogSvc := catalog.NewService(db)

	r := chi.NewRouter()
	r.Post("/products", cataloghttp.CreateProduct(catalogSvc).ServeHTTP)
	r.Get("/products/{product_name}", cataloghttp.GetProduct(catalogSvc).ServeHTTP)

	log.Println("Listening on :8080 ...")
	http.ListenAndServe(":8080", r)
}

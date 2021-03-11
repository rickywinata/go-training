package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	cataloghttp "github.com/rickywinata/go-training/catalog5/internal/catalog/http"
	"github.com/rickywinata/go-training/catalog5/internal/catalog/postgres"
	"github.com/rickywinata/go-training/catalog5/internal/catalog/service"
	"github.com/rickywinata/go-training/catalog5/internal/catalog/view"
)

func main() {
	db, err := sqlx.Open("postgres", "dbname=catalogsvc user=postgres password=password sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pgrepo := postgres.NewProductRepository(db)
	catalogsvc := service.NewProductService(pgrepo)
	productView := view.NewProductView(db)

	r := chi.NewRouter()
	r.Post("/products", cataloghttp.CreateProduct(catalogsvc).ServeHTTP)
	r.Get("/products/{product_name}", cataloghttp.GetProduct(productView).ServeHTTP)

	log.Println("Listening on :8080 ...")
	http.ListenAndServe(":8080", r)
}

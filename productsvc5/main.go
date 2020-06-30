package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	producthttp "github.com/rickywinata/go-training/productsvc5/product/http"
	"github.com/rickywinata/go-training/productsvc5/product/postgres"
	"github.com/rickywinata/go-training/productsvc5/product/service"
	"github.com/rickywinata/go-training/productsvc5/product/view"
)

func main() {
	db, err := sqlx.Open("postgres", "dbname=productsvc user=postgres password=password sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pgrepo := postgres.NewProductRepository(db)
	productSvc := service.NewProductService(pgrepo)
	productView := view.NewProductView(db)

	r := chi.NewRouter()
	r.Post("/products", producthttp.CreateProduct(productSvc).ServeHTTP)
	r.Get("/products/{product_name}", producthttp.GetProduct(productView).ServeHTTP)

	log.Println("Listening on :8080 ...")
	http.ListenAndServe(":8080", r)
}

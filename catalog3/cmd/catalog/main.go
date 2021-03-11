package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rickywinata/go-training/catalog3/internal/catalog"
	cataloghttp "github.com/rickywinata/go-training/catalog3/internal/catalog/http"
	"github.com/rickywinata/go-training/catalog3/internal/catalog/inmem"
	"github.com/rickywinata/go-training/catalog3/internal/catalog/service"
	"github.com/rickywinata/go-training/catalog3/internal/catalog/view"
)

func main() {
	repo := &inmem.ProductRepository{
		Data: []*catalog.Product{},
	}
	catalogsvc := service.NewProductService(repo)
	productView := view.NewProductView(repo)

	r := chi.NewRouter()
	r.Post("/products", cataloghttp.CreateProduct(catalogsvc).ServeHTTP)
	r.Get("/products/{product_name}", cataloghttp.GetProduct(productView).ServeHTTP)

	log.Println("Listening on :8080 ...")
	http.ListenAndServe(":8080", r)
}

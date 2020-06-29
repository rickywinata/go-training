package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rickywinata/go-training/productsvc3/product"
	producthttp "github.com/rickywinata/go-training/productsvc3/product/http"
	"github.com/rickywinata/go-training/productsvc3/product/inmem"
	"github.com/rickywinata/go-training/productsvc3/product/service"
	"github.com/rickywinata/go-training/productsvc3/product/view"
)

func main() {
	repo := &inmem.ProductRepository{
		Data: []*product.Product{},
	}
	productSvc := service.NewProductService(repo)
	productView := view.NewProductView(repo)

	r := chi.NewRouter()
	r.Post("/products", producthttp.CreateProduct(productSvc).ServeHTTP)
	r.Get("/products/{product_name}", producthttp.GetProduct(productView).ServeHTTP)

	log.Println("Listening on :8080 ...")
	http.ListenAndServe(":8080", r)
}

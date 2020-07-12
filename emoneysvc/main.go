package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	accounthttp "github.com/rickywinata/go-training/emoneysvc/emoney/http"
	"github.com/rickywinata/go-training/emoneysvc/emoney/postgres"
	"github.com/rickywinata/go-training/emoneysvc/emoney/service"
	"github.com/rickywinata/go-training/emoneysvc/emoney/view"
)

func main() {
	db, err := sqlx.Open("postgres", "dbname=emoneysvc user=postgres password=password sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pgrepo := postgres.NewAccountRepository(db)
	accountSvc := service.NewAccountService(pgrepo)
	accountView := view.NewAccountView(db)

	r := chi.NewRouter()
	r.Post("/accounts", accounthttp.CreateAccount(accountSvc).ServeHTTP)
	r.Get("/accounts/{account_id}", accounthttp.GetAccount(accountView).ServeHTTP)

	log.Println("Listening on :8080 ...")
	http.ListenAndServe(":8080", r)
}

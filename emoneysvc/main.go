package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	emoneyhttp "github.com/rickywinata/go-training/emoneysvc/emoney/http"
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

	accRepo := postgres.NewAccountRepository(db)
	accEntryRepo := postgres.NewAccountEntryRepository(db)
	accountSvc := service.NewAccountService(accRepo)
	accountView := view.NewAccountView(db)
	transactionSvc := service.NewTransactionService(accRepo, accEntryRepo)

	r := chi.NewRouter()
	r.Post("/accounts", emoneyhttp.CreateAccount(accountSvc).ServeHTTP)
	r.Get("/accounts/{account_id}", emoneyhttp.GetAccount(accountView).ServeHTTP)
	r.Post("/topups", emoneyhttp.Topup(transactionSvc).ServeHTTP)
	r.Post("/transfers", emoneyhttp.Transfer(transactionSvc).ServeHTTP)

	log.Println("Listening on :8080 ...")
	http.ListenAndServe(":8080", r)
}

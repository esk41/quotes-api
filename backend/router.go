package main

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/quotes", createQuoteHandler).Methods("POST")
	r.HandleFunc("/quotes", getAllQuotesHandler).Methods("GET")
	r.HandleFunc("/quotes/random", getRandomQuoteHandler).Methods("GET")
	r.HandleFunc("/quotes/{id}", deleteQuoteHandler).Methods("DELETE")
	r.HandleFunc("/quotes/{id:[0-9]+}", GetQuoteByIDHandler).Methods("GET")
	return r
}

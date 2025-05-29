package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var store = NewStore()

func createQuoteHandler(w http.ResponseWriter, r *http.Request) {
	var q Quote
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	q = store.AddQuote(q)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(q)
}

func getAllQuotesHandler(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	quotes := store.GetAll(author)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}

func getRandomQuoteHandler(w http.ResponseWriter, r *http.Request) {
	q, err := store.GetRandom()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(q)
}

func deleteQuoteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	if !store.DeleteByID(id) {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

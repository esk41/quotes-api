package main

import (
	"log"
	"net/http"
)

func main() {
	r := NewRouter()
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

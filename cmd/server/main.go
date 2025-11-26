package main

import (
	"log"
	"net/http"

	"github.com/afiffaizun/golang-web/internal/http"
)

func main() {
	router := http.NewRouter()

	log.Println("server running on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("failed to start server:  %v", err)
	}

}
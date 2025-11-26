package main

import (
	"log"
	stdhttp "net/http"

	apphttp "github.com/afiffaizun/golang-web/internal/http"
)

func main() {
	router := apphttp.NewRouter()

	log.Println("server running on :8081")
	if err := stdhttp.ListenAndServe(":8081", router); err != nil {
		log.Fatalf("failed to start server:  %v", err)
	}

}

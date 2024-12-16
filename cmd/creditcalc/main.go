package main

import (
	"log"
	"net/http"
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/dsemenov12/creditcalc/internal/handlers"
)

func main() {
	if error := run(); error != nil {
        fmt.Println(error)
    }
	
}

func run() error {
	router := chi.NewRouter()

	router.Get("/", handlers.GetRequest)
	router.Post("/", handlers.PostRequest)

	err := http.ListenAndServe(":8080", router)
	log.Println("Server started on :8080")
    if err != nil {
        return err
    }

	return nil
}
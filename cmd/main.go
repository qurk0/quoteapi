package main

import (
	"log"
	"net/http"
	"quotesapi/internal/quote"
)

func main() {
	router := http.NewServeMux()
	quote.NewQuoteHandler(router)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Printf("Server is running and listening address: %s\n=============================================\n", server.Addr)
	server.ListenAndServe()
}

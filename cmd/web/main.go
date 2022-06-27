package main

import (
	"log"
	"net/http"

	"github.com/LuisEduardo-M/Chat/internal/handlers"
)

func main() {
	mux := routes()

	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()

	log.Println("http://localhost:8000/")

	_ = http.ListenAndServe(":8000", mux)
}

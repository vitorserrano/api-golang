package main

import (
	"api/golang/internal/http"
	"log"
)

func main() {
	log.Println("Starting server 🚀")

	err := http.StartServer()
	log.Fatal(err)
}

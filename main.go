package main

import (
	"log"
	"net/http"

	"github.com/GNUalex/go-api/handler"
)

func main() {
	// Define routes
	http.HandleFunc("/posts", handler.GetPosts)
	http.HandleFunc("/post/{slug}", handler.GetPost)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

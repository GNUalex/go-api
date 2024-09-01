package main

import (
	"github.com/GNUalex/go-api/httpserver"
)

func main() {
	httpServer := &httpserver.HttpServer{}
	httpServer.ListenAndServe(":8080")
}

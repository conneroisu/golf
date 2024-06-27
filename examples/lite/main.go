package main

import (
	"fmt"

	"github.com/conneroisu/golf"
	"github.com/conneroisu/golf/domain"
)

// newR returns a new rename handler
func newR() golf.HandlerFunc {
	return func(w golf.ResponseWriter, r *domain.Request) {
		// TODO: Implement Rename
		fmt.Println("Rename Request Received" + r.Method)
		fmt.Println("writer:", w)
	}
}

// main is the entry point for the server
func main() {
	server := golf.DefaultMux

	AddRoutes(server)
}

// AddRoutes adds the routes to the server
func AddRoutes(server *golf.ServeMux) {
	server.Handle(domain.MethodTextDocumentRename, newR())
}

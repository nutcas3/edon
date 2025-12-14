package main

import (
	"log"

	"github.com/katungi/edon/internal/server"
)

func main() {
	srv, err := server.New("cmd/web/static")
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}
	defer srv.Close()

	if err := srv.Start("cmd/web/static"); err != nil {
		log.Fatal(err)
	}
}

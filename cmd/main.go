package main

import (
	"log"
	"makves/pkg/items"
)

func main() {
	server := items.NewRestServer()
	log.Fatal(server.Run())
}

package main

import (
	"log"

	"github.com/thaessaep/billingSystem/internal/httpserver"
)

// @title App Api
// @version 1.0
// @description Api Server for Avito

// @host localhost:8080

func main() {
	config := httpserver.NewConfig()

	s := httpserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}

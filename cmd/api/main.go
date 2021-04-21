package main

import (
	"log"

	"github.com/gabrielopesantos/filesharing_api/pkg/utl/server"
)

func main() {
	Run()
}

func Run() {
	srv, err := server.New()
	if err != nil {
		log.Fatalf("Error setting up server: %v\n", err)
	}

	srv.Run(":8080")
}

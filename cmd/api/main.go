package main

import (
	"log"

	postgres "github.com/gabrielopesantos/myDrive-api/pkg/utl/database"
	"github.com/gabrielopesantos/myDrive-api/pkg/utl/server"
)

func main() {

	psqlDB, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatalf("Postgres init: %s", err)
	} else {
		// log.Infof("Postgres connected, status %#v", psqlDB.Status())
		log.Println("Postgres connected, psqlDB.Stats()")
	}
	defer psqlDB.Close()

	srv := server.NewServer(psqlDB)
	if err = srv.Run(); err != nil {
		log.Fatal(err)
	}
}

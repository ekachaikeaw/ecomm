package main

import (
	"log"

	"github.com/ekachaikeaw/ecomm/db"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()

	log.Println("successfully connected to database")

	// do something with database
	// st := storer.NewMySQLStorer(db.GetDB())

}
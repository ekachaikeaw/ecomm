package main

import (
	"log"

	"github.com/ekachaikeaw/ecomm/db"
	"github.com/ekachaikeaw/ecomm/ecomm-api/handler"
	"github.com/ekachaikeaw/ecomm/ecomm-api/server"
	"github.com/ekachaikeaw/ecomm/ecomm-api/storer"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()

	log.Println("successfully connected to database")

	// do something with database
	st := storer.NewMySQLStorer(db.GetDB())
	srv := server.NewServer(st)
	hdl := handler.NewHandler(srv)
	handler.RegisterRoutes(hdl)
	handler.Start(":8080")
}

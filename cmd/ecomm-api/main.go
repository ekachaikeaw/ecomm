package main

import (
	"log"

	"github.com/ekachaikeaw/ecomm/db"
	"github.com/ekachaikeaw/ecomm/ecomm-api/handler"
	"github.com/ekachaikeaw/ecomm/ecomm-api/server"
	"github.com/ekachaikeaw/ecomm/ecomm-api/storer"
	"github.com/ianschenck/envflag"
)

const minSecretKeySize = 32

func main() {
	var secretKey = envflag.String("SECRET_KEY", "01234567890123456789012345678901")
	if len(*secretKey) < minSecretKeySize {
		log.Fatal("SECRET_KEY must be at least %d characters", minSecretKeySize)
	}

	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()

	log.Println("successfully connected to database")

	// do something with database
	st := storer.NewMySQLStorer(db.GetDB())
	srv := server.NewServer(st)
	hdl := handler.NewHandler(srv, *secretKey)
	handler.RegisterRoutes(hdl)
	handler.Start(":8080")
}

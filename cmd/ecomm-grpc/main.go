package main

import (
	"log"
	"net"

	"github.com/ekachaikeaw/ecomm/db"
	"github.com/ekachaikeaw/ecomm/ecomm-grpc/pb"
	"github.com/ekachaikeaw/ecomm/ecomm-grpc/server"
	"github.com/ekachaikeaw/ecomm/ecomm-grpc/storer"
	"github.com/ianschenck/envflag"
	"google.golang.org/grpc"
)

func main() {
	var (
		svcAddr = envflag.String("SVC_ADDR", "0.0.0.0:9091", "address where ecomm-grpc service is listening on")
	)

	// instantiate db
	// instantiate server
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()

	log.Println("successfully connected to database")

	// do something with database
	st := storer.NewMySQLStorer(db.GetDB())
	srv := server.NewServer(st)
	// register our server with the gRPC server
	grpcSrv := grpc.NewServer()
	pb.RegisterEcommServer(grpcSrv, srv)

	listener, err := net.Listen("tcp", *svcAddr)
	if err != nil {
		log.Fatalf("listener failed: %v", err)
	}

	log.Printf("server listening on: %s", *svcAddr)
	err = grpcSrv.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

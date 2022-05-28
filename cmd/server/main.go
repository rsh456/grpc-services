package main

import (
	"log"

	"github.com/rsh456/grpc-services/internal/db"
	"github.com/rsh456/grpc-services/internal/rocket"
)

func Run() error {
	// responsible for initializing and starting
	// our gRPC server
	rocketStore, err := db.New()
	if err != nil {
		return err
	}
	_ = rocket.New(rocketStore)

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"

	"github.com/404th/grpcserver/client"
	"github.com/404th/grpcserver/server"
)

func main() {
	address := ":6767"

	// server
	server := server.NewServer()
	if err := server.Run(address); err != nil {
		log.Fatalf("Error while running new server: %e", err)
	}

	// client
	cl := client.NewClient()
	client, err := cl.InitClient(address)
	if err != nil {
		log.Fatalf("Error wgile initializing client: %e", err)
	}

	// runnning interfaces
}

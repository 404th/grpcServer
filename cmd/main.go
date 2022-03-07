package main

import (
	"context"
	"log"
	"time"

	"github.com/404th/grpcserver/client"
	pb "github.com/404th/grpcserver/generated/user_service"
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

	//// context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	users := make(map[string]int32)

	users["Jojo"] = 75
	users["Samuel"] = 54
	users["Inna"] = 19

	for name, age := range users {
		us, err := client.CreateUser(ctx, &pb.NewUser{
			Name: name,
			Age:  age,
		})
		if err != nil {
			log.Fatalf("Error while creating user: %e", err)
		}

		log.Printf(`Details:
						Name: %s,
						Age: %d,
						ID: %d,		
		`, us.GetName(), us.GetAge(), us.GetId())
	}
}

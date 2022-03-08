package main

import (
	"context"
	"log"
	"time"

	pb "github.com/404th/grpcserver/generated/user_service"
	"google.golang.org/grpc"
)

const (
	address = ":8080"
)

func main() {
	conn, err := grpc.Dial("tcp", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Error while dialing: %e", err)
	}
	defer conn.Close()

	cc := pb.NewUserManagementClient(conn)

	// ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// creating users
	us, err := cc.CreateUser(ctx, &pb.NewUser{Name: "Sulaqmon", Age: 35})
	if err != nil {
		log.Fatalf("error first user : %e", err)
	}

	log.Printf("New User: %v", us)

	cc.CreateUser(ctx, &pb.NewUser{Name: "Luqmon", Age: 71})
	r, err := cc.CreateUser(ctx, &pb.NewUser{Name: "Haydaralo", Age: 102})
	if err != nil {
		log.Fatalf("Error while creating third user: %e", err)
	}
	// getting all users
	n, err := cc.GetUsers(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Error while getting all users: %e", err)
	}
	log.Printf("ALL USERS: %v", n)

	// deleting
	d, err := cc.DeleteUser(ctx, &pb.IDTracker{Id: r.GetId()})
	if err != nil {
		log.Fatalf("Error while all users after deleting one of them : %e", err)
	}

	log.Printf("AFTER DELETED: %v", d)

	// after deleting the third one getting all users
	au, err := cc.GetUsers(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Error while all users after deleting one of them : %e", err)
	}
	log.Printf("ALL USERS AFTER DELETED: %v", au)
}

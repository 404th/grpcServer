package main

import (
	"context"
	"log"
	"time"

	pb "github.com/404th/grpcserver/generated/user_service"
	"google.golang.org/grpc"
)

var (
	address = "localhost:6060"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Error while dialing: %v", err)
	}
	defer conn.Close()

	cc := pb.NewUserManagementClient(conn)

	//
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//
	var urs = make(map[string]int32)
	urs["Humoyun"] = 25
	urs["Javlon"] = 76
	urs["Koke"] = 33

	var d_id int32

	for name, age := range urs {
		u, err := cc.CreateUser(ctx, &pb.NewUser{
			Name: name,
			Age:  age,
		})

		if err != nil {
			log.Fatalf("Error while creating %v-user %v\n", name, err)
		}

		d_id = u.GetId()
		log.Printf("Details: NAME: %v, AGE: %v, ID: %v\n", u.GetName(), u.GetAge(), u.GetId())
	}

	// getting all
	ul, err := cc.GetUsers(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Error while getting users: %v\n", err)
	}

	log.Printf("ALL USERS: %v\n", ul.GetUsersList())

	// delete
	del, err := cc.DeleteUser(ctx, &pb.IDTracker{Id: d_id})
	if err != nil {
		log.Printf("Error while deleting user: %v\n", err)
	}

	log.Printf("DELETED RESULT: %s\n", del.GetDetailsOfDeleted())

	// after deleted all users
	au, err := cc.GetUsers(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Error while getting users after deleting one of them: %v", err)
	}

	log.Printf("AFTER DELETED ALL USERS: %v\n", au.GetUsersList())
}

package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"

	"github.com/404th/grpcserver/db"
	pb "github.com/404th/grpcserver/generated/user_service"
	"google.golang.org/grpc"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
	users_list *pb.UsersList
}

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{
		users_list: &pb.UsersList{},
	}
}

const (
	address = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error while listening %e", err)
	}

	gServer := grpc.NewServer()

	// registering services
	pb.RegisterUserManagementServer(gServer, &UserManagementServer{})
	log.Printf("Server is running on PORT: %d", address)

	// serving
	if err = gServer.Serve(lis); err != nil {
		log.Fatalf("Error while serving %e", err)
	}
}

func (s *UserManagementServer) CreateUser(ctx context.Context, us *pb.NewUser) (*pb.User, error) {
	log.Printf("Receiving details: NAME: %s AGE: %d", us.GetName(), us.GetAge())

	id := rand.Intn(10000)

	created := &pb.User{
		Name: us.GetName() + " has been added",
		Age:  us.GetAge(),
		Id:   int32(id),
	}

	// add new element to fake db
	db.Database.UsersList = append(db.Database.UsersList, created)

	return created, nil
}

func (s *UserManagementServer) GetUsers(ctx context.Context, mt *pb.Empty) (*pb.UsersList, error) {
	return db.Database, nil
}

func (s *UserManagementServer) DeleteUser(ctx context.Context, id *pb.IDTracker) (*pb.Deleted, error) {
	var deleteUser *pb.User
	for ind, user := range db.Database.UsersList {
		if user.GetId() == id.GetId() {
			db.Database.UsersList = append(db.Database.UsersList[:ind], db.Database.UsersList[ind+1:]...)
			deleteUser = user
		}
	}

	// returner
	str := fmt.Sprintf("DELETE ID: %d NAME: %s AGE: %d", deleteUser.GetId(), deleteUser.GetName(), deleteUser.GetAge())

	return &pb.Deleted{
		DetailsOfDeleted: str,
	}, nil
}

package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"

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

var (
	address = ":6060"
)

func (s *UserManagementServer) Run() error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error while listening on port: %v", address)
		return err
	}
	gs := grpc.NewServer()
	pb.RegisterUserManagementServer(gs, s)
	return gs.Serve(lis)
}

func main() {
	var ss *UserManagementServer = NewUserManagementServer()
	if err := ss.Run(); err != nil {
		log.Fatalf("Error while running server: %v", err)
	}
}

//////////////////////////
func (s *UserManagementServer) CreateUser(ctx context.Context, nu *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: Name: %s Age: %d\n", nu.GetName(), nu.GetAge())

	var user_id int32 = int32(rand.Intn(10000))
	created_user := &pb.User{Name: nu.GetName(), Age: nu.GetAge(), Id: user_id}
	s.users_list.UsersList = append(s.users_list.UsersList, created_user)

	return created_user, nil
}

func (s *UserManagementServer) GetUsers(ctx context.Context, mt *pb.Empty) (*pb.UsersList, error) {
	return s.users_list, nil
}

func (s *UserManagementServer) DeleteUser(ctx context.Context, id *pb.IDTracker) (*pb.Deleted, error) {
	var del_usr *pb.User

	for ind, usr := range s.users_list.UsersList {
		if usr.GetId() == id.GetId() {
			del_usr = usr
			s.users_list.UsersList = append(s.users_list.UsersList[:ind], s.users_list.UsersList[ind+1:]...)
		}
	}

	str := fmt.Sprintf("Deleted item details: NAME: %s AGE: %d ID: %d", del_usr.GetName(), del_usr.GetAge(), del_usr.GetId())
	return &pb.Deleted{
		DetailsOfDeleted: str,
	}, nil
}

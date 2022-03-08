package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"

	pb "github.com/404th/grpcserver/generated/user_service"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{}
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

	readByte, err := ioutil.ReadFile("users.json")
	var users_list *pb.UsersList = &pb.UsersList{}

	var user_id int32 = int32(rand.Intn(10000))
	created_user := &pb.User{Name: nu.GetName(), Age: nu.GetAge(), Id: user_id}

	if err != nil {
		if os.IsNotExist(err) {
			log.Print("File not found.Creating new one.")
			users_list.UsersList = append(users_list.UsersList, created_user)

			jsonBytes, err := protojson.Marshal(users_list)
			if err != nil {
				log.Fatalf("Error while marshaling users_list : %v", err)
			}
			if err := ioutil.WriteFile("users.json", jsonBytes, 0664); err != nil {
				log.Fatalf("Error while writing to file: %v", err)
			}

			return created_user, nil
		} else {
			log.Fatalf("Error reading file: %v", err)
		}
	}

	if err = protojson.Unmarshal(readByte, users_list); err != nil {
		log.Fatalf("Error while un marshalling %v", err)
		return nil, err
	}

	users_list.UsersList = append(users_list.UsersList, created_user)
	writeByte, err := protojson.Marshal(users_list)
	if err != nil {
		log.Fatalf("Error while marshalling %v", err)
		return nil, err
	}

	if err := ioutil.WriteFile("users.json", writeByte, 0664); err != nil {
		log.Fatalf("Error while writing to file: %v", err)
		return nil, err
	}

	return created_user, nil
}

func (s *UserManagementServer) GetUsers(ctx context.Context, mt *pb.Empty) (*pb.UsersList, error) {
	readByte, err := ioutil.ReadFile("users.json")
	var users_list *pb.UsersList = &pb.UsersList{}
	if err != nil {
		log.Fatalf("Error while reading file: %v", err)
		return nil, err
	}
	if err := protojson.Unmarshal(readByte, users_list); err != nil {
		log.Fatalf("Error while unmarshalling: %v", err)
		return nil, err
	}

	return users_list, nil
}

func (s *UserManagementServer) DeleteUser(ctx context.Context, id *pb.IDTracker) (*pb.Deleted, error) {
	var del_usr *pb.User
	var users_list *pb.UsersList = &pb.UsersList{}

	readByte, err := ioutil.ReadFile("users.json")
	if err != nil {
		log.Fatalf("Error while reading in deleting one of them: %v", err)
		return nil, err
	}

	if err := protojson.Unmarshal(readByte, users_list); err != nil {
		log.Fatalf("Error while unmarshalling deleting one of them: %v", err)
		return nil, err
	}

	for ind, usr := range users_list.UsersList {
		if usr.GetId() == id.GetId() {
			del_usr = usr
			users_list.UsersList = append(users_list.UsersList[:ind], users_list.UsersList[ind+1:]...)
		}
	}

	wrJson, err := protojson.Marshal(users_list)
	if err != nil {
		log.Fatalf("Error while marshalling: %v", err)
		return nil, err
	}

	if err := ioutil.WriteFile("users.json", wrJson, 0664); err != nil {
		log.Fatalf("Error while writin to file in deleting one: %v", err)
		return nil, err
	}

	str := fmt.Sprintf("Deleted item details: NAME: %s AGE: %d ID: %d", del_usr.GetName(), del_usr.GetAge(), del_usr.GetId())
	return &pb.Deleted{
		DetailsOfDeleted: str,
	}, nil
}

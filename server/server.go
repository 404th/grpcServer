package server

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/404th/grpcserver/generated/user_service"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedUserManagementServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(port string) error {
	// listening
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error while listening on PORT%s", port)
		return err
	}

	// creating grpc server
	gs := grpc.NewServer()

	if err = gs.Serve(lis); err != nil {
		log.Fatalf("Error while serving gRPC: %e", err)
		return err
	}

	return nil
}

//////// IMPLEMENTATION ////////
func (s *Server) CreateUser(ctx context.Context, nu *pb.NewUser) (*pb.User, error) {
	// generating ID
	id := int32(rand.Intn(10000))

	log.Printf("Got details: NAME: %v, AGE: %d while creating new user", nu.GetName(), nu.GetAge())

	return &pb.User{
		Name: nu.GetName(),
		Age:  nu.GetAge(),
		Id:   id,
	}, nil
}

package client

import (
	"log"

	pb "github.com/404th/grpcserver/generated/user_service"
	"google.golang.org/grpc"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) InitClient(port string) (pb.UserManagementClient, error) {
	conn, err := grpc.Dial(port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Error while dialing: %e", err)
	}
	defer conn.Close()
	us := pb.NewUserManagementClient(conn)

	return us, nil
}

//// IMPLEMENTATION

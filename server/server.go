package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/404th/grpcserver/generated/user_service"
	pgx "github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
	conn *pgx.Conn
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
	psql_url := "postgres://postgres:postgres@localhost:5432/uss"
	conn, err := pgx.Connect(context.Background(), psql_url)
	if err != nil {
		log.Fatalf("Error while connecting to PSQL db: %v", err)
		return
	}
	defer conn.Close(context.Background())
	var ss *UserManagementServer = NewUserManagementServer()
	ss.conn = conn
	if err := ss.Run(); err != nil {
		log.Fatalf("Error while running server: %v", err)
	}
}

//////////////////////////
func (s *UserManagementServer) CreateUser(ctx context.Context, nu *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: Name: %s Age: %d\n", nu.GetName(), nu.GetAge())
	created_user := &pb.User{Name: nu.GetName(), Age: nu.GetAge()}

	qy := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL NOT NULL PRIMARY KEY,
			name TEXT,
			age INT
		);
	`)

	_, err := s.conn.Exec(context.Background(), qy)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while creating an user and write to psql: %v", err)
		os.Exit(1)
		return nil, err
	}

	tx, err := s.conn.Begin(context.Background())
	if err != nil {
		log.Fatalf("failed writing to psql: %v", err)
	}

	_, err = tx.Exec(context.Background(), "INSERT INTO users(name, age) VALUES($1, $2);", nu.GetName(), nu.GetAge())
	if err != nil {
		log.Fatalf("Error while inserting to psql: %v", err)
		tx.Rollback(context.Background())
	}

	tx.Commit(context.Background())

	return created_user, nil
}

func (s *UserManagementServer) GetUsers(ctx context.Context, mt *pb.Empty) (*pb.UsersList, error) {
	var users_list *pb.UsersList = &pb.UsersList{}

	sql := `SELECT * FROM users;`

	rows, err := s.conn.Query(context.Background(), sql)
	if err != nil {
		log.Fatalf("Error while getting all users: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		user := pb.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			log.Fatalf("Error while scanning rows (getting all users): %v", err)
		}
		users_list.UsersList = append(users_list.UsersList, &user)
	}

	return users_list, nil
}

func (s *UserManagementServer) DeleteUser(ctx context.Context, id *pb.IDTracker) (*pb.Deleted, error) {
	sql := `DELETE FROM users WHERE id = $1;`

	_, err := s.conn.Exec(context.Background(), sql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while deleting item: %v", err)
		os.Exit(1)
	}

	return &pb.Deleted{
		DetailsOfDeleted: "DELETED",
	}, nil
}

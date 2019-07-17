package main

import (
	"fmt"
	pb "github.com/fernandochristyanto/todogrpc/proto/todo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"path/filepath"
)

const (
	address    = "localhost:8080"
	certPath   = "creds/server.crt"
	serverName = "localhost" // Common Name
)

func main() {
	// Create TLS credentials
	cert, _ := filepath.Abs(certPath)
	creds, err := credentials.NewClientTLSFromFile(cert, serverName)
	if err != nil {
		log.Fatalf("Error creating TLS creds: %s", err)
	}

	// Setup connection to gRPC server
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	// Setup client for service
	todoClient := pb.NewTodoTransactionClient(conn)

	getTodosResponse, err := todoClient.GetTodos(context.Background(), &pb.GetTodosRequest{})
	if err != nil {
		log.Fatalf("Could not transact: %v", err)
	}

	for _, todo := range getTodosResponse.Todos {
		fmt.Println(todo.GetTaskName())
	}
}

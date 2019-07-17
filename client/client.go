package main

import (
	"fmt"
	cred "github.com/fernandochristyanto/todogrpc/creds"
	pb "github.com/fernandochristyanto/todogrpc/proto/todo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
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
	grpcDialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}
	// Setup connection to gRPC server
	conn, err := grpc.Dial(address, grpcDialOpts...)
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	// Setup client for service
	todoClient := pb.NewTodoTransactionClient(conn)

	/**
	 * Calling the service
	 */
	// Create a new metadata with serverkey
	md := metadata.Pairs("serverkey", cred.ServerKey)
	// Attach metadata to context
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	getTodosResponse, err := todoClient.GetTodos(ctx, &pb.GetTodosRequest{})

	if err != nil {
		log.Fatalf("Could not transact: %v", err)
	}

	// Print values
	for _, todo := range getTodosResponse.Todos {
		fmt.Println(todo.GetTaskName())
	}
}

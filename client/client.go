package main

import (
	"fmt"
	pb "github.com/fernandochristyanto/todogrpc/proto/todo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:8080"
)

func main() {
	// Setup connection to gRPC server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
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

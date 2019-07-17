package main

import (
	pb "github.com/fernandochristyanto/todogrpc/proto/todo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":8080"
)

// Implements the TodoTransaction interface
type todoTransactionImpl struct{}

func (todoTransaction *todoTransactionImpl) GetTodos(ctx context.Context, in *pb.GetTodosRequest) (*pb.GetTodosResponse, error) {
	todos := []*pb.Todo{
		&pb.Todo{
			TaskName:  "Name 1",
			Completed: true,
		},
		&pb.Todo{
			TaskName:  "Name 2",
			Completed: true,
		},
		&pb.Todo{
			TaskName:  "Name 3",
			Completed: true,
		},
		&pb.Todo{
			TaskName:  "Name 4",
			Completed: true,
		},
	}
	todosResponse := pb.GetTodosResponse{
		Todos: todos,
	}
	return &todosResponse, nil
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Fail to open connection: %v", err)
	}

	// Declare new gRPC service
	grpcServer := grpc.NewServer()

	// Register transaction servers
	pb.RegisterTodoTransactionServer(grpcServer, &todoTransactionImpl{})

	// Register gRPC server
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

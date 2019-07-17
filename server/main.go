package main

import (
	pb "github.com/fernandochristyanto/todogrpc/proto/todo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	certPath = "../creds/server.crt"
	keyPath  = "../creds/server.key"
	port     = ":8080"
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

/**
* Steps in creating the credentials (SSL/TLS Certificates with openssl)
* Files needed to be generated:
*  - server.key: a private RSA key to sign and authenticate the public key
*  - server.pem/server.crt: self-signed X.509 public keys for distribution
*  - rootca.crt: a certificate authority public key for signing .csr files
*  - host.csr: a certificate signing request to access the CA
*
* ------------------------------------------------------------------------
* Commands:
*  - openssl genrsa -out server.key 2048
* 	 2048 bit RSA key (stronger keys are available as well).
*  - openssl req -new -x509 -sha256 -key server.key \
			 -out server.crt -days 3650
	 Generate the certificate, and will also prompt you for some questions about the location, organization, and contact of the certificate holder.
*
*  - openssl req -new -sha256 -key server.key -out server.csr
*  - openssl x509 -req -sha256 -in server.csr -signkey server.key \
             -out server.crt -days 3650
*    Generate a certificate signing request (.csr)
*/

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Fail to open connection: %v", err)
	}

	// Create new TLS credentials
	creds, err := credentials.NewServerTLSFromFile(certPath, keyPath)
	if err != nil {
		log.Fatalf("Could not load TLS keys: %s", err)
	}
	// Declare new gRPC server with TLS
	grpcServer := grpc.NewServer(grpc.Creds(creds))

	// Register transaction servers
	pb.RegisterTodoTransactionServer(grpcServer, &todoTransactionImpl{})

	// Register gRPC server
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

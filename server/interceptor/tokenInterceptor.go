package interceptor

import (
	"github.com/fernandochristyanto/todogrpc/creds"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

var (
	errMissingMetadata  = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidServerKey = status.Errorf(codes.Unauthenticated, "invalid server key")
)

func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}

	serverKey := strings.TrimSpace(authorization[0])
	return serverKey == creds.ServerKey
}

// EnsureValidStaticApplicationKey validates every Unary call made to gRPC server provides a ["serverkey"] in the metadata with a correct value
func EnsureValidStaticApplicationKey(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	metadata, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}

	if !valid(metadata["serverkey"]) {
		return nil, errInvalidServerKey
	}
	// Continue execution of handler after ensuring a valid token.
	return handler(ctx, req)
}

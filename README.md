# go-todogrpc
A simple sandbox project(s) (actually collection of projects, divided per branch) covering gRPC core features. Each branch demonstrates different feature of gRPC.

## Feature Covered
- Basics
- Authentication (TLS, static token)

## Branches
- **basic/single-model**  
  A basic client server model with one message type (todo)

- **auth/tls**  
  Simple client server model with TLS Security (first one of two security features offered by gRPC)

- **auth/global-unary-interceptor-statictoken**  
  Simple client server model with TLS Security + Global unary interceptor for serverkey (token) validation


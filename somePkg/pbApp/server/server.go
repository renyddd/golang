package main

// https://grpc.io/docs/languages/go/basics/

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/renyddd/golang/somePkg/protocolbuffers"
	"google.golang.org/grpc"
)

type Server struct {
	protocolbuffers.UnimplementedGreeterServer
	greeting string
}

func (s *Server) SayHello(ctx context.Context, in *protocolbuffers.HelloRequest) (*protocolbuffers.HelloReply, error) {
	log.Println("[log] server ", s, in)
	return &protocolbuffers.HelloReply{
		Message: s.greeting + in.Name,
	}, nil
}

func newServer() *Server {
	s := &Server{
		greeting: "Hello, ",
	}
	return s
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8081))
	if err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	protocolbuffers.RegisterGreeterServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

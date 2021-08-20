package main

// https://grpc.io/docs/languages/go/basics/

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	gw "github.com/renyddd/golang/somePkg/pbApp/server/v1"

	"google.golang.org/grpc"
)

type Server struct {
	gw.UnimplementedGreeterServer
	greetingPrefix string
}

func (s *Server) SayHello(ctx context.Context, in *gw.HelloRequest) (*gw.HelloReply, error) {
	log.Println("[log] server ", s, in)
	return &gw.HelloReply{
		Msg: s.greetingPrefix + in.Name,
	}, nil
}

func newServer() *Server {
	s := &Server{
		greetingPrefix: "Hello, ",
	}
	return s
}

func main() {
	ServerWithoutGateway()
}

func ServerWithoutGateway() {
	lis, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	gw.RegisterGreeterServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

func ServerWithGateway() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterGreeterHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%d", 8081), opts)
	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		panic(err)
	}
}

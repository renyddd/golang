package main

// https://grpc.io/docs/languages/go/basics/

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	gw "github.com/renyddd/golang/somePkg/pbApp/server/v1"

	"google.golang.org/grpc"
)

type Server struct {
	gw.UnimplementedGreeterServer
	greeting string
}

func (s *Server) SayHello(ctx context.Context, in *gw.HelloRequest) (*gw.HelloReply, error) {
	log.Println("[log] server ", s, in)
	return &gw.HelloReply{
		Msg: s.greeting + in.Name,
	}, nil
}

func newServer() *Server {
	s := &Server{
		greeting: "Hello, ",
	}
	return s
}

func main() {
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

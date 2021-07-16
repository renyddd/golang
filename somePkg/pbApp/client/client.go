package main

import (
	"context"
	"log"

	"github.com/renyddd/golang/somePkg/protocolbuffers"
	"google.golang.org/grpc"
)

func Rpc_SayHello(ctx context.Context, client protocolbuffers.GreeterClient, msg string) (string, error) {
	hRequest := &protocolbuffers.HelloRequest{
		Name: msg,
	}
	hReplay, err := client.SayHello(ctx, hRequest)
	if err != nil {
		return "", err
	}

	return hReplay.Msg, nil
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	serverAddr := "localhost:8081"
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := protocolbuffers.NewGreeterClient(conn)

	msg, err := Rpc_SayHello(context.Background(), client, "Alice")
	if err != nil {
		panic(err)
	}
	log.Println("rpc replay: ", msg)

	msg, err = Rpc_SayHello(context.Background(), client, "Bob")
	if err != nil {
		panic(err)
	}
	log.Println("rpc replay: ", msg)
}

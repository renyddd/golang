package main

import (
	"context"
	"log"

	pb "github.com/renyddd/golang/somePkg/pbApp/server/v1"
	"google.golang.org/grpc"
)

func Rpc_SayHello(ctx context.Context, client pb.GreeterClient, msg string) (string, error) {
	hRequest := &pb.HelloRequest{
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

	client := pb.NewGreeterClient(conn)

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

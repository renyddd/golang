package main

// https://kubernetes.io/zh/docs/setup/production-environment/container-runtimes/
import (
	"github.com/containerd/containerd"
	"log"
)

func main() {
	if err := Example(); err != nil {
		log.Fatalln(err)
	}
}

func Example() error {
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		return err
	}
	defer client.Close()
	return nil
}

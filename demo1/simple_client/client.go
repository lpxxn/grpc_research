package main

import (
	"context"
	"log"

	"github.com/lpxxn/grpc_research"
	pb "github.com/lpxxn/grpc_research/helloworld"
	"google.golang.org/grpc"
)

const (
	address = "127.0.0.1:7051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: grpc_research.GetFullName(2220)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	r, err = c.SayHello(context.Background(), &pb.HelloRequest{Name: grpc_research.GetFullName(2220)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

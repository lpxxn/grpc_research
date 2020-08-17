package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/lpxxn/grpc_research"
	pb "github.com/lpxxn/grpc_research/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	address     = "127.0.0.1:7051"
	defaultName = "world"
)

var kacp = keepalive.ClientParameters{
	Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithKeepaliveParams(kacp))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	ticker := time.NewTicker(time.Second / 20)
	go func() {
		for {
			<-ticker.C
			r, err = c.SayHello(context.Background(), &pb.HelloRequest{Name: grpc_research.GetFullName(500)})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			log.Printf("Greeting: %s", r.GetMessage())
		}
	}()

	osCh := make(chan os.Signal)

	signal.Notify(osCh, os.Interrupt)
	<-osCh
}

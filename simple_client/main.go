package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Pallinder/go-randomdata"
	"google.golang.org/grpc"

	"github.com/lpxxn/grpc_research/protos/api"
	"github.com/lpxxn/grpc_research/protos/model"
)

var port int

func init() {
	flag.IntVar(&port, "port", 10001, "grpc port")
}

func main() {
	rand.Seed(time.Now().Unix())
	conn, err := grpc.Dial(fmt.Sprintf(":%d", port), grpc.WithInsecure())
	//conn, err := grpc.Dial(fmt.Sprintf("192.168.10.94:%d", 62937), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := api.NewStudentSrvClient(conn)
	//time.Sleep(2*time.Second)
	for i := 0; i < 1; i++ {
		student := &model.Student{
			Id:   rand.Int63(),
			Name: randomdata.FullName(randomdata.RandomGender) + randomdata.City(),
			Age:  rand.Int31n(30),
		}
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			time.Sleep(time.Second*3)
			cancel()
		}()
		r, err := c.NewStudent(ctx, student)
		if err != nil {
			panic(err)
		}
		//cancel()
		fmt.Println("add student ", r.Code)
		time.Sleep(time.Second * 5)
		fmt.Println("-----------")
	}
	time.Sleep(time.Second)
	//time.Sleep(time.Second * 10)
}

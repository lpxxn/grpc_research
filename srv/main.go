package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	"github.com/lpxxn/grpc_research/common"
	"github.com/lpxxn/grpc_research/protos/api"
	"github.com/lpxxn/grpc_research/protos/model"
)

var port int

func init() {
	flag.IntVar(&port, "port", 10001, "grpc port")
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterStudentSrvServer(grpcServer, &common.StudentSrv{StudentList: []*model.Student{
		&model.Student{Id: 1, Name: "tom", Age: 5},
		&model.Student{Id: 2, Name: "jerry", Age: 6},
	}})
	fmt.Println("serv running...")
	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}

func WithInterceptor() []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

				return handler(ctx, req)
			},
		)),
	}
}

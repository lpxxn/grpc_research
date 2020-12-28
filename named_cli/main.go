package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"

	"github.com/lpxxn/grpc_research/common"
	"github.com/lpxxn/grpc_research/protos/api"
)

func main() {
	r := common.NewResolver([]string{"http://127.0.0.1:2379"}, common.ServName)
	resolver.Register(r)
	addr := fmt.Sprintf("%s:///%s", r.Scheme(), common.ServName)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
		grpc.WithBlock(),
	)
	if err != nil {
		panic(err)
	}
	client := api.NewStudentSrvClient(conn)

	for i := 0; i < 20; i++ {
		rev, err := client.StudentByID(context.Background(), &api.QueryStudent{})
		if err != nil {
			panic(err)
		}
		for _, item := range rev.StudentList {
			fmt.Println(item)
		}
		time.Sleep(time.Second / 2)
	}

}

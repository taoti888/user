package main

import (
	"context"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"github.com/taoti888/user/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	address := "consul://10.102.81.2:8500/user_dev?healthy=true"
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := proto.NewPermissionsClient(conn)

	md := metadata.Pairs("x-api-key", "xxxxxxxx")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, err := client.GetPermissions(ctx, &proto.GetPermissionsRequest{
		Id: 1,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Name)
}

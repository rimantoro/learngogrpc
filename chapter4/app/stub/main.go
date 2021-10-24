package main

import (
	"context"
	"fmt"
	"rimantoro/learngogrpc/chapter4/proto"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	stub := proto.NewStarfriendsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req, err := stub.ListFilms(ctx, &proto.ListFilmsRequest{})
	if err != nil {
		panic(err)
	}

	fmt.Println(req.GetFilms())
}

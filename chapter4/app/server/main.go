package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"rimantoro/learngogrpc/chapter4/proto"
	"rimantoro/learngogrpc/chapter4/service"

	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 8080, "The port on which gRPC server will listen")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fail(err)
	}
	fmt.Printf("Listening on %v\n", lis)
	svr := grpc.NewServer()

	proto.RegisterStarfriendsServer(svr, &service.StarfriendsImpl{})

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("Shutting down...")
		svr.GracefulStop()
	}()

	if err := svr.Serve(lis); err != nil {
		fail(err)
	}
}

func fail(err error) {
	fmt.Println(os.Stderr, err)
	os.Exit(1)
}

package main

import (
	"context"
	"fmt"

	pb "gRPC_demo/helloworld/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9652", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("failed to connect: %v", err)
		return
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "cao"})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}

	fmt.Printf("Greeting: %s !\n", r.Message)
}
package main

import (
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"

	pb "github.com/grpc-example/simple/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50001"
)

func main() {
	conn, err := grpc.Dial(address,
		grpc.WithTransportCredentials(insecure.NewCredentials()), //禁用https
	)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	name := "lin"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println(r.Message)
}

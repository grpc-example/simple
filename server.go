package main

import (
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	pb "github.com/grpc-example/simple/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	PORT = ":50001"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("request: ", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	grpc.EnableTracing = true
	pb.RegisterGreeterServer(s, &server{})
	log.Println("rpc服务已经开启")
	s.Serve(lis)
}

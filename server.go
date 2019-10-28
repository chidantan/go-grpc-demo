package main

import (
	pb "joyrry/grpc/protos/hello"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	// PORT 服务端的
	PORT = ":50001"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Println("request: ", in.Greeting)
	return &pb.HelloResponse{Reply: "Hello " + in.Greeting}, nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	log.Println("rpc服务已经开启")
	s.Serve(lis)
}

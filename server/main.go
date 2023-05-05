package main

import (
	"log"
	"net"

	pb "go-grpc/proto"

	"google.golang.org/grpc"
)

const (
	port = ":10005"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterGreetServiceServer(server, &helloServer{})
	log.Printf("server started at %v", lis.Addr)
	if err = server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

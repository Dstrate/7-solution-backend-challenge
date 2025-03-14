package main

import (
	"fmt"
	"net"

	pb "github.com/Dstrate/7-solution-backend-challenge-3/beef/proto"
	"google.golang.org/grpc"
)

const port = "50051"

type server struct {
	pb.UnimplementedBeefServiceServer
}

func main() {
	listener, _ := net.Listen("tcp", ":"+port)

	grpcServer := grpc.NewServer()
	pb.RegisterBeefServiceServer(grpcServer, &server{})
	fmt.Printf("Start Server Port %s\n", port)
	grpcServer.Serve(listener)
}

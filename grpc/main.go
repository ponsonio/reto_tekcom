package main

import (
	"fmt"
	"github.com/ponsonio/reto_tekcom/multiplication"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := multiplication.Server{}

	grpcServer := grpc.NewServer()

	multiplication.RegisterMultiplicationServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

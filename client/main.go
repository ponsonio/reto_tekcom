package main

import (
	"github.com/ponsonio/reto_tekcom/multiplication"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := multiplication.NewMultiplicationServiceClient(conn)

	response, err := c.Multiply(context.Background(), &multiplication.Factors{
		A: 77,
		B: 11,
	})
	if err != nil {
		log.Fatalf("Error when calling mULTIPLY: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}

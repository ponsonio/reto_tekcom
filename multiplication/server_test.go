package multiplication

import (
	"context"
	"github.com/go-playground/assert/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"

)

//More test are needed, error codes, etc, neglected fue time constraints

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	RegisterMultiplicationServiceServer(server, &Server{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestMultiplicationServer_Multiply(t *testing.T) {

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := NewMultiplicationServiceClient(conn)

	request := &Factors{
		A:             10,
		B:             88,
	}

	response, err := client.Multiply(ctx, request)

	assert.Equal(t, response.Body, "880")
}
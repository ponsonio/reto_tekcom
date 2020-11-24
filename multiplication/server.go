package multiplication

import (
	"github.com/ponsonio/reto_tekcom/services"
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) Multiply(ctx context.Context, in *Factors) (*Result, error) {
	log.Printf("Receive message body from client: %v times %v", in.A, in.B)
	//errors needs to be handled
	mult, _ := services.NewMultiplicator()
	res, _ := mult.Multiply(int64(in.A), int64(in.B))
	return &Result{Body: res.String()}, nil
}

package pkg

import (
	"context"
	"fmt"

	"github.com/wrallen/sampleBazel6/api"
)

type HelloServer struct {
	api.UnimplementedHelloServiceServer
}

func (h *HelloServer) Hello(ctx context.Context, req *api.HelloRequest) (*api.HelloResponse, error) {
	fmt.Println("====")
	return nil, nil
}

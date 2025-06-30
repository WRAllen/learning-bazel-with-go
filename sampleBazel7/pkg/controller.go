package pkg

import (
	"context"
	"fmt"

	api "github.com/wrallen/sampleBazel7/api/hello/v1"
)

type HelloServer struct {
	api.UnimplementedHelloServiceServer
}

func (h *HelloServer) Hello(ctx context.Context, req *api.HelloRequest) (*api.HelloResponse, error) {
	fmt.Println("====")
	return nil, nil
}

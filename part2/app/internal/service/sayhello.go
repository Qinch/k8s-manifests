package service

import (
	"context"
	"log"

	pbapp "github.com/Qinch/k8s-manifests/part2/lib/proto/app"
)

type AppService struct {
	pbapp.UnimplementedAppServiceServer
}

func (c *AppService) SayHello(ctx context.Context, in *pbapp.HelloRequest) (*pbapp.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pbapp.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

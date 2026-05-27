package service

import (
	"context"
	"log"

	"github.com/Qinch/k8s-manifests/part2/gateway/internal/grpcclients"
	pbapp "github.com/Qinch/k8s-manifests/part2/lib/proto/app"
	pbgw "github.com/Qinch/k8s-manifests/part2/lib/proto/gateway"
)

type GwService struct {
	pbgw.UnimplementedGatewayServiceServer
}

func (c *GwService) Hello(ctx context.Context, in *pbgw.HelloRequest) (*pbgw.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	resp, err := grpcclients.AppClient.SayHello(ctx, &pbapp.HelloRequest{Name: in.GetName()})
	if err != nil {
		log.Printf("Error: %v", err)
		return &pbgw.HelloResponse{Message: resp.GetMessage()}, err
	}
	return &pbgw.HelloResponse{Message: resp.GetMessage()}, nil
}

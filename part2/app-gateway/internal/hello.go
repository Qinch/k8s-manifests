package service

import (
	"context"
	"log"

	"github.com/Qinch/k8s-manifests/part2/app-gateway/internal/grpcclients"
	pbapp "github.com/Qinch/k8s-manifests/part2/lib/proto/app"
	pbgw "github.com/Qinch/k8s-manifests/part2/lib/proto/app_gateway"
)

type GwService struct {
	pbgw.UnimplementedAppGatewayServiceServer
}

func (c *GwService) Hello(ctx context.Context, in *pbgw.HelloRequest) (*pbgw.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	resp, err := grpcclients.AppGrpcClient.SayHello(ctx, &pbapp.HelloRequest{Name: in.GetName()})
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	return &pbgw.HelloResponse{Message: resp.GetMessage()}, nil
}

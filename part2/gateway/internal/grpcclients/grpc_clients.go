package grpcclients

import (
	"log"
	"time"

	pbapp "github.com/Qinch/k8s-manifests/part2/lib/proto/app"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

var (
	AppClient  pbapp.AppServiceClient
	grpcTarget = "dns://app.test.svc.cluster.local:50051"
)

func initGrpcCLient() (*pbapp.AppServiceClient, error) {
	conn, err := grpc.NewClient(grpcTarget, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                10 * time.Second,
			Timeout:             1 * time.Second,
			PermitWithoutStream: true,
		}),
	})
	if err != nil {
		log.Printf("grpc client init err: %v", err)
		return nil, err
	}

	client := pbapp.NewAppServiceClient(conn)
	return &client, nil
}

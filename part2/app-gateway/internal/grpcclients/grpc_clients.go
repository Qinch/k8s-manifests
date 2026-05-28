package grpcclients

import (
	"log"
	"os"
	"time"

	pbapp "github.com/Qinch/k8s-manifests/part2/lib/proto/app"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

var (
	grpcConn      *grpc.ClientConn
	AppGrpcClient pbapp.AppServiceClient
	// 简化处理，这里其实可以放到config文件，然后通过test或者prod前缀来获取对应的target
	/*
		test.grpcTarget: "dns:///app.prod.svc.cluster.local:50051"
		prod.grpcTarget: "dns:///app.prod.svc.cluster.local:50051"
	*/
	grpcProdTarget = "dns:///app.prod.svc.cluster.local:50051"
	grpcTestTarget = "dns:///app.test.svc.cluster.local:50051"
)

func init() {
	client, err := initGrpcClient()
	if err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}
	AppGrpcClient = client
}

func initGrpcClient() (pbapp.AppServiceClient, error) {
	env, ok := os.LookupEnv("ENV")
	var grpcTarget string
	if !ok || (env == "test") {
		grpcTarget = grpcTestTarget
	} else {
		grpcTarget = grpcProdTarget
	}
	conn, err := grpc.NewClient(grpcTarget,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                10 * time.Second,
			Timeout:             1 * time.Second,
			PermitWithoutStream: true,
		}),
	)
	if err != nil {
		log.Printf("grpc client init err: %v", err)
		return nil, err
	}

	client := pbapp.NewAppServiceClient(conn)
	return client, nil
}

// CloseGrpcConn 安全关闭gRPC连接
func CloseGrpcConn() error {
	if grpcConn != nil {
		log.Println("closing gRPC connection...")
		err := grpcConn.Close()
		grpcConn = nil
		AppGrpcClient = nil
		if err != nil {
			log.Printf("gRPC connection close failed: %v", err)
			return err
		}
		log.Println("gRPC connection closed successfully")
	}
	return nil
}

package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	service "github.com/Qinch/k8s-manifests/part2/app-gateway/internal"
	pbgw "github.com/Qinch/k8s-manifests/part2/lib/proto/app-gateway"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcPort     = "50051"
	grpcEndpoint = "localhost:" + grpcPort
	httpPort     = "8080"
	httpEndpoint = "localhost:" + httpPort
)

func main() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcSrv := grpc.NewServer()
	pbgw.RegisterGatewayServiceServer(grpcSrv, &service.GwService{})
	go func() {
		if err := grpcSrv.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
		log.Println("grpc server is running")
	}()

	ctx := context.Background()
	gwMux := runtime.NewServeMux()
	err = pbgw.RegisterGatewayServiceHandlerFromEndpoint(ctx, gwMux, grpcEndpoint, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatalf("failed to register: %v", err)
	}

	httpSrv := http.Server{
		Addr:    httpEndpoint,
		Handler: gwMux,
	}
	go func() {
		if err := httpSrv.ListenAndServe(); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
		log.Println("http server is running")
	}()

	// wait signal
	shutdown := make(chan os.Signal)
	signal.Notify(shutdown,
		syscall.SIGKILL,
		syscall.SIGTERM)
	sig := <-shutdown
	log.Printf("received shutdown signal:%v", sig.String())

	// close http server
	err = httpSrv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("failed to shutdown: %v", err)
	}

	// close grpc server
	grpcSrv.GracefulStop()
}

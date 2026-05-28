package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Qinch/k8s-manifests/part2/app/internal/service"
	pbapp "github.com/Qinch/k8s-manifests/part2/lib/proto/app"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcPort     = "50051"
	grpcEndpoint = "0.0.0.0:" + grpcPort
	httpPort     = "8080"
	httpEndpoint = "0.0.0.0:" + httpPort
)

func main() {
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcSrv := grpc.NewServer()
	pbapp.RegisterAppServiceServer(grpcSrv, &service.AppService{})
	go func() {
		log.Println("grpc server is starting")
		if err := grpcSrv.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	ctx := context.Background()
	gwMux := runtime.NewServeMux()
	err = pbapp.RegisterAppServiceHandlerFromEndpoint(ctx, gwMux, grpcEndpoint, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		log.Fatalf("failed to register: %v", err)
	}

	rootMux := http.NewServeMux()
	// 容器健康检查 httpGet: /health
	rootMux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"UP"}`))
	})

	rootMux.Handle("/", gwMux)
	httpSrv := http.Server{
		Addr:    httpEndpoint,
		Handler: rootMux,
	}
	go func() {
		log.Println("http server is starting")
		if err := httpSrv.ListenAndServe(); (err != nil)&&(err != http.ErrServerClosed) {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// wait signal
	shutdown := make(chan os.Signal)
	signal.Notify(shutdown,
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

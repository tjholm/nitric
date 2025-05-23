package server

import (
	"log"
	"net"

	storagepb "github.com/nitrictech/nitric/core/pkg/proto/storage/v1"
	pubsubpb "github.com/nitrictech/nitric/core/pkg/proto/topics/v1"
	"github.com/nitrictech/nitric/server/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type RegisterFunction[T any] func(name string, constructor runtime.PluginConstructor[T])

func RegisterPlugins[T any](register RegisterFunction[T], plugins map[string]runtime.PluginConstructor[T]) {
	// Register the plugins
	for name, constructor := range plugins {
		register(name, constructor)
	}
}

func Start() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	// Register plugin router
	storagepb.RegisterStorageServer(srv, &GrpcServer{})
	pubsubpb.RegisterTopicsServer(srv, &GrpcServer{})

	// Register reflection service on gRPC server
	reflection.Register(srv)

	log.Printf("Starting server on %s", lis.Addr().String())

	srv.Serve(lis)
}

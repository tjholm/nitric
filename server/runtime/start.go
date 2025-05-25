package runtime

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"

	pubsubpb "github.com/nitrictech/nitric/proto/pubsub/v2"
	storagepb "github.com/nitrictech/nitric/proto/storage/v2"
	"github.com/nitrictech/nitric/server/runtime/plugin"
	"github.com/nitrictech/nitric/server/runtime/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RegisterPlugins[T any](register plugin.Register[T], plugins map[string]plugin.Constructor[T]) {
	// Register the plugins
	for name, constructor := range plugins {
		register(name, constructor)
	}
}

func Start(cmd string) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	// Register plugin router
	storagepb.RegisterStorageServer(srv, &GrpcServer{})
	pubsubpb.RegisterPubsubServer(srv, &GrpcServer{})

	// Register reflection service on gRPC server
	reflection.Register(srv)

	log.Printf("Starting server on %s", lis.Addr().String())

	// Start the runtime services in a goroutine
	go func() {
		srv.Serve(lis)
	}()

	// Get the PORT of the local service
	servicePort := os.Getenv("PORT")

	// Start the actual nitric service
	// TODO: Determine how we will provide this command
	cmdParts := strings.Split(cmd, " ")
	runCmd := exec.Command(cmdParts[0], cmdParts[1:]...)
	runCmd.Env = os.Environ()
	runCmd.Stdout = os.Stdout
	runCmd.Stderr = os.Stderr

	if err := runCmd.Start(); err != nil {
		log.Fatalf("failed to start service: %v", err)
	}

	// Start the service gateway and proxy
	err = service.Start(service.NewHttpServerProxy(fmt.Sprintf("localhost:%s", servicePort)))
	if err != nil {
		log.Fatalf("failed to start ingress: %v", err)
	}
}

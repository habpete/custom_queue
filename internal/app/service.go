package app

import (
	"fmt"
	"log"
	"net"

	pb "github.com/custom_queue/pkg/proto"
	"google.golang.org/grpc"
)

func Start() error {
	listener, err := net.Listen("tcp", "localhost:8082")
	if err != nil {
		return fmt.Errorf("failed to listen port: %d", 8082)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMessageQueueServer(grpcServer, &Service{})

	log.Print("service started on port 8082")
	if err := grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("failed to start GRPC service: %v", err)
	}

	return nil
}

type Service struct {
	pb.UnimplementedMessageQueueServer
}

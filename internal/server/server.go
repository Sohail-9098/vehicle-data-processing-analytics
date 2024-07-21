package main

import (
	"context"
	"log"
	"net"

	"github.com/Sohail-9098/vehicle-data-processing-analytics/internal/protobufs/vehicle"
	"google.golang.org/grpc"
)

type server struct {
	vehicle.UnimplementedDataProcessingServiceServer
}

func (s *server) ProcessTelemetryData(ctx context.Context, data *vehicle.Telemetry) (*vehicle.Empty, error) {
	log.Printf("received data: %v", data)
	return &vehicle.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen on port 50052: %v", err)
	}

	grpcServer := grpc.NewServer()
	vehicle.RegisterDataProcessingServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start grpc server: %v", err)
	}
}

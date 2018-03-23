package main

import (
	pb "PlayerTracker/api"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type playerTrackingServer struct {
}

func (playerTrackingServer) CardInEvent(ctx context.Context, in *pb.CardIn) (*pb.Status, error) {
	//s.savedCustomers = append(s.savedCustomers, in)
	return &pb.Status{Message: "OK ", Points: 100, DisplayName: "Test"}, nil
	//return &pb.CustomerResponse{Id: in.Id, Success: true}, nil
	//return &pb.Status{Message: pb}, nil
}

func (playerTrackingServer) MeterUpdateEvent(ctx context.Context, in *pb.MeterUpdate) (*pb.PointUpdate, error) {

	return &pb.PointUpdate{MachNumber: "1234", CardNumber: 1234, NewPointBalance: 321}, nil
}

func (playerTrackingServer) CardOutEvent(ctx context.Context, in *pb.CardOut) (*pb.Status, error) {
	return &pb.Status{Message: "OK", Points: 321, DisplayName: ""}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var players playerTrackingServer
	// Creates a new gRPC server
	s := grpc.NewServer()
	pb.RegisterPlayerTrackingServer(s, players)
	s.Serve(lis)
}

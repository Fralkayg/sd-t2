package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb2 "./Service2"
	"google.golang.org/grpc"
)

type server struct {
	seguimiento int
	lock        bool
}

const (
	port = ":50051"
)

func (s *server) SendDistributionProposal(ctx context.Context, in *pb2.DistributionRequest) (*pb2.DistributionReply, error) {
	fmt.Println("Llego el archivo " + in.FileName)
	return &pb2.DistributionReply{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	s := server{}

	//Inicializacion variables servidor logistica.
	s.seguimiento = 0
	s.lock = false

	pb2.RegisterDataToNameServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

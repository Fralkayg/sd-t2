package main

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	"os"

	pb "./Service"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	seguimiento int
	lock        bool
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetMensaje())
	return &pb.HelloReply{Mensaje: "Hello " + in.GetMensaje()}, nil
}

func (s *server) SendChunk(ctx context.Context, in *pb.ChunkInformation) (*pb.ChunkStatus, error) {
	fileName := in.FileName + "_" + in.ChunkIndex
	_, err := os.Create("Chunks/" + fileName)

	if err != nil {
		os.Exit(1)
	}

	ioutil.WriteFile("Chunks/"+fileName, in.Chunk, os.ModeAppend)
	return &pb.ChunkStatus{Status: fileName + "OK"}, nil
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

	pb.RegisterFileManagementServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

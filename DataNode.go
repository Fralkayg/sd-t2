package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"

	pb "./Service"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	seguimiento int
	lock        bool
	file        File
}

type File struct {
	fileName   string
	totalParts int
	chunks     []Chunk
}

type Chunk struct {
	Chunk      []byte
	ChunkIndex int
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetMensaje())
	return &pb.HelloReply{Mensaje: "Hello " + in.GetMensaje()}, nil
}

func (s *server) SendChunk(ctx context.Context, in *pb.ChunkInformation) (*pb.ChunkStatus, error) {
	if in.LastChunk {
		fmt.Println("Llego el último chunk.")
		var chunk Chunk
		chunk.ChunkIndex = int(in.ChunkIndex)
		chunk.Chunk = in.Chunk
		s.file.chunks = append(s.file.chunks, chunk)

		if in.Option == 1 {
			//Centralizado
			fmt.Println("Generar distribucion centralizada")
			for i := 0; i < s.file.totalParts; i++ {
				fileName := s.file.fileName + "_" + strconv.Itoa(int(s.file.chunks[i].ChunkIndex))
				_, err := os.Create("Chunks/" + fileName)

				if err != nil {
					os.Exit(1)
				}

				ioutil.WriteFile("Chunks/"+fileName, s.file.chunks[i].Chunk, os.ModeAppend)

			}

		} else {
			//Distribuido
			fmt.Println("Generar distribucion distribuida")
		}
	} else {
		s.file.fileName = in.FileName
		s.file.totalParts = int(in.TotalParts)

		var chunk Chunk
		chunk.ChunkIndex = int(in.ChunkIndex)
		chunk.Chunk = in.Chunk
		s.file.chunks = append(s.file.chunks, chunk)
	}

	return &pb.ChunkStatus{Status: "Parte " + strconv.Itoa(int(in.ChunkIndex)) + " OK"}, nil
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

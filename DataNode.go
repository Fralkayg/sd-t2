package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	pb "./Service"
	pb2 "./Service2"
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

const (
	nameNodeAddress = "dist56:50051"
	defaultName     = "world"
)

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
			generateCentralizedDistribution(s)

			// for i := 0; i < s.file.totalParts; i++ {
			// 	fileName := s.file.fileName + "_" + strconv.Itoa(int(s.file.chunks[i].ChunkIndex))
			// 	_, err := os.Create("Chunks/" + fileName)

			// 	if err != nil {
			// 		os.Exit(1)
			// 	}

			// 	ioutil.WriteFile("Chunks/"+fileName, s.file.chunks[i].Chunk, os.ModeAppend)

			// }

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

func generateCentralizedDistribution(s *server) {
	fmt.Println("Generando distribucion centralizada")
	var firstNodeDistribution []int32
	var secondNodeDistribution []int32
	var thirdNodeDistribution []int32

	for i := 0; i < s.file.totalParts; i++ {
		var result int32
		result = int32(i % 3)
		if result == 0 {
			firstNodeDistribution = append(firstNodeDistribution, int32(i))
		} else if result == 1 {
			secondNodeDistribution = append(secondNodeDistribution, int32(i))
		} else {
			thirdNodeDistribution = append(thirdNodeDistribution, int32(i))
		}
	}
	fmt.Println("Paso la generacion de arreglos")

	conn, err := grpc.Dial(nameNodeAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb2.NewDataToNameServiceClient(conn)

	fmt.Println("Se conecto")
	distributionReply, err := c.SendDistributionProposal(context.Background(), &pb2.DistributionRequest{
		FileName:   s.file.fileName,
		TotalParts: int32(s.file.totalParts),
		Machines: []*pb2.DistributionRequest_MachineInformation{
			{Address: "dist53", Distribution: firstNodeDistribution, Status: 1},
			{Address: "dist54", Distribution: secondNodeDistribution, Status: 1},
			{Address: "dist55", Distribution: thirdNodeDistribution, Status: 1},
		},
	})

	fmt.Println("Envio datos")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(distributionReply.FileName)
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

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
	pb2 "./Service2"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	seguimiento    int
	lock           bool
	file           File
	currentAddress string
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

func (s *server) SaveChunk(ctx context.Context, in *pb.StoreChunkRequest) (*pb.StoreChunkReply, error) {
	_, err := os.Create("Chunks/" + in.FileName)

	if err != nil {
		os.Exit(1)
	}

	ioutil.WriteFile("Chunks/"+in.FileName, in.Chunk, os.ModeAppend)
}

func (s *server) CheckNodeStatus(ctx context.Context, in *pb.StatusRequest) (*pb.StatusReply, error) {
	return &pb.StatusReply{Online: true}, nil
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
		s.currentAddress = in.Address

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

	conn, err := grpc.Dial(nameNodeAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb2.NewDataToNameServiceClient(conn)

	distributionReply, err := c.SendDistributionProposal(context.Background(), &pb2.DistributionRequest{
		FileName:   s.file.fileName,
		TotalParts: int32(s.file.totalParts),
		Machines: []*pb2.DistributionRequest_MachineInformation{
			{Address: "dist53:50051", Distribution: firstNodeDistribution, Status: 1},
			{Address: "dist54:50051", Distribution: secondNodeDistribution, Status: 1},
			{Address: "dist55:50051", Distribution: thirdNodeDistribution, Status: 1},
		},
	})

	fmt.Println("Estados")
	fmt.Println(distributionReply.Machines[0].Status)
	fmt.Println(distributionReply.Machines[1].Status)
	fmt.Println(distributionReply.Machines[2].Status)

	if distributionReply.Machines[0].Status == 1 {
		fmt.Println(len(distributionReply.Machines[0].Distribution))
		for j := 0; j < len(distributionReply.Machines[0].Distribution); j++ {
			//Crear chunk
			fmt.Println("Maquina: " + distributionReply.Machines[0].Address)
			// Chunk: []byte = s.file.chunks[distributionReply.Machines[0].Distribution[i]].Chunk)
			fmt.Println("Chunk index: " + strconv.Itoa(int(s.file.chunks[distributionReply.Machines[0].Distribution[j]].ChunkIndex)))
			fileName := distributionReply.FileName + "_" + strconv.Itoa(int(s.file.chunks[distributionReply.Machines[0].Distribution[j]].ChunkIndex))
			chunk := s.file.chunks[distributionReply.Machines[0].Distribution[j]].Chunk
			if distributionReply.Machines[0].Address == s.currentAddress {
				_, err := os.Create("Chunks/" + fileName)

				if err != nil {
					os.Exit(1)
				}

				ioutil.WriteFile("Chunks/"+fileName, chunk, os.ModeAppend)
			} else {
				conn, err := grpc.Dial(distributionReply.Machines[0].Address, grpc.WithInsecure())
				if err != nil {
					log.Fatalf("did not connect: %v", err)
				}
				defer conn.Close()

				c := pb.NewFileManagementServiceClient(conn)

				response, _ := c.SaveChunk(context.Background(), &pb.StoreChunkRequest{
					FileName: fileName,
					Chunk:    chunk,
				})

				fmt.Println("Mensaje: ", response.Status)
			}
		}
	}
	if distributionReply.Machines[1].Status == 1 {
		fmt.Println(len(distributionReply.Machines[1].Distribution))
		for j := 0; j < len(distributionReply.Machines[1].Distribution); j++ {
			//Crear chunk
			fmt.Println("Maquina: " + distributionReply.Machines[1].Address)
			// Chunk: []byte = s.file.chunks[distributionReply.Machines[1].Distribution[i]].Chunk)
			fmt.Println("Chunk index: " + strconv.Itoa(int(s.file.chunks[distributionReply.Machines[1].Distribution[j]].ChunkIndex)))
			fileName := distributionReply.FileName + "_" + strconv.Itoa(int(s.file.chunks[distributionReply.Machines[1].Distribution[j]].ChunkIndex))
			chunk := s.file.chunks[distributionReply.Machines[1].Distribution[j]].Chunk
			if distributionReply.Machines[1].Address == s.currentAddress {

				_, err := os.Create("Chunks/" + fileName)

				if err != nil {
					os.Exit(1)
				}

				ioutil.WriteFile("Chunks/"+fileName, chunk, os.ModeAppend)
			} else {
				conn, err := grpc.Dial(distributionReply.Machines[1].Address, grpc.WithInsecure())
				if err != nil {
					log.Fatalf("did not connect: %v", err)
				}
				defer conn.Close()

				c := pb.NewFileManagementServiceClient(conn)

				response, _ := c.SaveChunk(context.Background(), &pb.StoreChunkRequest{
					FileName: fileName,
					Chunk:    chunk,
				})

				fmt.Println("Mensaje: ", response.Status)
			}
		}
	}
	if distributionReply.Machines[2].Status == 1 {
		fmt.Println(len(distributionReply.Machines[2].Distribution))
		for j := 0; j < len(distributionReply.Machines[2].Distribution); j++ {
			//Crear chunk
			fmt.Println("Maquina: " + distributionReply.Machines[2].Address)
			// Chunk: []byte = s.file.chunks[distributionReply.Machines[2].Distribution[i]].Chunk)
			fmt.Println("Chunk index: " + strconv.Itoa(int(s.file.chunks[distributionReply.Machines[2].Distribution[j]].ChunkIndex)))
			fileName := distributionReply.FileName + "_" + strconv.Itoa(int(s.file.chunks[distributionReply.Machines[2].Distribution[j]].ChunkIndex))
			chunk := s.file.chunks[distributionReply.Machines[2].Distribution[j]].Chunk
			if distributionReply.Machines[2].Address == s.currentAddress {
				_, err := os.Create("Chunks/" + fileName)

				if err != nil {
					os.Exit(1)
				}

				ioutil.WriteFile("Chunks/"+fileName, chunk, os.ModeAppend)
			} else {
				conn, err := grpc.Dial(distributionReply.Machines[2].Address, grpc.WithInsecure())
				if err != nil {
					log.Fatalf("did not connect: %v", err)
				}
				defer conn.Close()

				c := pb.NewFileManagementServiceClient(conn)

				response, _ := c.SaveChunk(context.Background(), &pb.StoreChunkRequest{
					FileName: fileName,
					Chunk:    chunk,
				})

				fmt.Println("Mensaje: ", response.Status)

			}
		}
	}

	// saveChunks(s, distributionReply)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(distributionReply.FileName)
}

func saveChunks(s *server, distributionReply *pb2.DistributionReply) {
	for i := 0; i < s.file.totalParts; i++ {
		fileName := s.file.fileName + "_" + strconv.Itoa(int(s.file.chunks[i].ChunkIndex))
		_, err := os.Create("Chunks/" + fileName)

		if err != nil {
			os.Exit(1)
		}

		ioutil.WriteFile("Chunks/"+fileName, s.file.chunks[i].Chunk, os.ModeAppend)

	}
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

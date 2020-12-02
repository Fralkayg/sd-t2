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

func (s *server) RetrieveChunk(ctx context.Context, in *pb.ChunkRequest) (*pb.ChunkReply, error) {
	file, err := os.Open("./Chunks/" + in.FileName)
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return &pb.ChunkReply{Chunk: content}, nil
}

func (s *server) SaveChunk(ctx context.Context, in *pb.StoreChunkRequest) (*pb.StoreChunkReply, error) {
	_, err := os.Create("Chunks/" + in.FileName)

	if err != nil {
		os.Exit(1)
	}

	ioutil.WriteFile("Chunks/"+in.FileName, in.Chunk, os.ModeAppend)

	return &pb.StoreChunkReply{Status: "OK"}, nil
}

func (s *server) CheckNodeStatus(ctx context.Context, in *pb.StatusRequest) (*pb.StatusReply, error) {
	return &pb.StatusReply{Online: true}, nil
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetMensaje())
	return &pb.HelloReply{Mensaje: "Hello " + in.GetMensaje()}, nil
}

func (s *server) SendChunks(ctx context.Context, in *pb.ChunkInformation) (*pb.ChunkStatus, error) {
	for s.lock {
	}

	s.lock = true

	s.file.fileName = in.FileName
	s.file.totalParts = int(in.TotalParts)
	s.currentAddress = in.Address

	for i := 0; i < len(in.Chunks); i++ {
		var chunk Chunk
		chunk.ChunkIndex = int(in.Chunks[i].ChunkIndex)
		chunk.Chunk = in.Chunks[i].Chunk
		s.file.chunks = append(s.file.chunks, chunk)
	}

	if in.Option == 1 {
		generateCentralizedDistribution(s)
	} else {
		notValid := true
		firstNodeStatus := int32(0)
		secondNodeStatus := int32(0)
		thirdNodeStatus := int32(0)

		//Propuesta inicial asume 3 nodos
		availableNodes := int32(3)

		for notValid {
			fmt.Println("Generando propuesta de distribución")
			first, second, third, nodes := generateDistributedDistribution(s)
			if nodes == availableNodes {
				notValid = false
				firstNodeStatus = 1
				secondNodeStatus = 1
				thirdNodeStatus = 1
			} else {
				firstNodeStatus = first
				secondNodeStatus = second
				thirdNodeStatus = third
				availableNodes = nodes
			}
		}
		if firstNodeStatus == 1 {
			fmt.Println("DataNode 53 incluido en propuesta")
		}
		if secondNodeStatus == 1 {
			fmt.Println("DataNode 54 incluido en propuesta")
		}
		if thirdNodeStatus == 1 {
			fmt.Println("DataNode 55 incluido en propuesta")
		}

		generateDistribution(s, availableNodes, s.file.totalParts, firstNodeStatus, secondNodeStatus, thirdNodeStatus)

	}

	s.file = File{}

	s.lock = false
	return &pb.ChunkStatus{Status: "Se recibio el archivo " + in.FileName}, nil
}

func generateDistribution(s *server, availableNodes int32, totalParts int, firstNodeStatus int32, secondNodeStatus int32, thirdNodeStatus int32) {
	var firstNodeDistribution []int32
	var secondNodeDistribution []int32
	var thirdNodeDistribution []int32
	for i := 0; i < totalParts; i++ {
		var result int32
		if availableNodes == 3 {
			result = int32(i % 3)
			if result == 0 && firstNodeStatus == 1 {
				firstNodeDistribution = append(firstNodeDistribution, int32(i))
			} else if result == 1 && secondNodeStatus == 1 {
				secondNodeDistribution = append(secondNodeDistribution, int32(i))
			} else if thirdNodeStatus == 1 {
				thirdNodeDistribution = append(thirdNodeDistribution, int32(i))
			}
		} else if availableNodes == 2 {
			result = int32(i % 2)
			if firstNodeStatus == 1 && secondNodeStatus == 1 {
				firstDistribution, first, secondDistribution, _ := makeLocalDistribution("dist53:50051", "dist54:50051", result, i)
				if first {
					firstNodeDistribution = append(firstNodeDistribution, firstDistribution)
				} else {
					secondNodeDistribution = append(secondNodeDistribution, secondDistribution)
				}
			} else if firstNodeStatus == 1 && thirdNodeStatus == 1 {
				firstDistribution, first, secondDistribution, _ := makeLocalDistribution("dist53:50051", "dist55:50051", result, i)
				if first {
					firstNodeDistribution = append(firstNodeDistribution, firstDistribution)
				} else {
					thirdNodeDistribution = append(thirdNodeDistribution, secondDistribution)
				}
			} else if secondNodeStatus == 1 && thirdNodeStatus == 1 {
				firstDistribution, first, secondDistribution, _ := makeLocalDistribution("dist54:50051", "dist55:50051", result, i)
				if first {
					secondNodeDistribution = append(secondNodeDistribution, firstDistribution)
				} else {
					thirdNodeDistribution = append(thirdNodeDistribution, secondDistribution)
				}
			}
		} else {
			if firstNodeStatus == 1 {
				firstNodeDistribution = append(firstNodeDistribution, int32(i))
			} else if secondNodeStatus == 1 {
				secondNodeDistribution = append(secondNodeDistribution, int32(i))
			} else {
				thirdNodeDistribution = append(thirdNodeDistribution, int32(i))
			}
		}
	}

	conn, err := grpc.Dial(nameNodeAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb2.NewDataToNameServiceClient(conn)

	fmt.Println("Largo distribucion DataNode 53: ")
	fmt.Println(len(firstNodeDistribution))
	fmt.Println("Largo distribucion DataNode 54: ")
	fmt.Println(len(secondNodeDistribution))
	fmt.Println("Largo distribucion DataNode 55: ")
	fmt.Println(len(thirdNodeDistribution))

	distributionReply, err := c.SendDistribution(context.Background(), &pb2.DistributionRequest2{
		FileName:   s.file.fileName,
		TotalParts: int32(s.file.totalParts),
		Machines: []*pb2.DistributionRequest2_MachineInformation{
			{Address: "dist53:50051", Distribution: firstNodeDistribution, Status: firstNodeStatus},
			{Address: "dist54:50051", Distribution: secondNodeDistribution, Status: secondNodeStatus},
			{Address: "dist55:50051", Distribution: thirdNodeDistribution, Status: thirdNodeStatus},
		},
	})

	fmt.Println("Distribuyendo chunks a nodos incluidos en propuesta.")

	if distributionReply.Machines[0].Status == 1 {
		for j := 0; j < len(distributionReply.Machines[0].Distribution); j++ {
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
		for j := 0; j < len(distributionReply.Machines[1].Distribution); j++ {
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
		for j := 0; j < len(distributionReply.Machines[2].Distribution); j++ {
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

}

func makeLocalDistribution(address1 string, address2 string, result int32, i int) (int32, bool, int32, bool) {
	var firstNodeDistribution int32
	var secondNodeDistribution int32
	first := false
	second := false
	if result == 0 {
		firstNodeDistribution = int32(i)
		first = true
	} else if result == 1 {
		secondNodeDistribution = int32(i)
		second = true
	}
	return firstNodeDistribution, first, secondNodeDistribution, second
}

func generateDistributedDistribution(s *server) (int32, int32, int32, int32) {
	firstNodeStatus := int32(0)
	secondNodeStatus := int32(0)
	thirdNodeStatus := int32(0)
	availableNodes := int32(0)

	//Chequear Propuesta
	fmt.Println("Verificando estado de los DataNode")
	for i := 53; i < 56; i++ {
		address := "dist" + strconv.Itoa(i) + port
		if address != s.currentAddress {
			status := checkNodeStatus(address)
			if i == 53 {
				firstNodeStatus = status
			} else if i == 54 {
				secondNodeStatus = status
			} else {
				thirdNodeStatus = status
			}
			if status == 1 {
				availableNodes++
			}

			fmt.Println("Estado de dist" + strconv.Itoa(i) + ": ")
			fmt.Println(status)
		} else {
			fmt.Println("Estado de dist" + strconv.Itoa(i) + ": ")
			if i == 53 {
				fmt.Println(firstNodeStatus)
				firstNodeStatus = 1
			} else if i == 54 {
				fmt.Println(secondNodeStatus)
				secondNodeStatus = 1
			} else if i == 55 {
				fmt.Println(thirdNodeStatus)
				thirdNodeStatus = 1
			}
			availableNodes++

		}
	}
	fmt.Println("Nodos disponibles: ", strconv.Itoa(int(availableNodes)))
	return firstNodeStatus, secondNodeStatus, thirdNodeStatus, availableNodes

}

func checkNodeStatus(address string) int32 {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	c := pb.NewFileManagementServiceClient(conn)

	_, connectionError := c.CheckNodeStatus(context.Background(), &pb.StatusRequest{Online: true})
	if connectionError != nil {
		return 0
	}
	return 1
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

	fmt.Println("Generando propuesta inicial incluyendo a los 3 nodos.")

	conn, err := grpc.Dial(nameNodeAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb2.NewDataToNameServiceClient(conn)

	// fmt.Println("Datos a enviar a NameNode")
	// fmt.Println(s.file.fileName)
	// fmt.Println(s.file.totalParts)

	distributionReply, err := c.SendDistributionProposal(context.Background(), &pb2.DistributionRequest{
		FileName:   s.file.fileName,
		TotalParts: int32(s.file.totalParts),
		Machines: []*pb2.DistributionRequest_MachineInformation{
			{Address: "dist53:50051", Distribution: firstNodeDistribution, Status: 1},
			{Address: "dist54:50051", Distribution: secondNodeDistribution, Status: 1},
			{Address: "dist55:50051", Distribution: thirdNodeDistribution, Status: 1},
		},
	})

	// fmt.Println("Estados")
	// fmt.Println(distributionReply.Machines[0].Status)
	// fmt.Println(distributionReply.Machines[1].Status)
	// fmt.Println(distributionReply.Machines[2].Status)

	if distributionReply.Machines[0].Status == 1 {
		fmt.Println("Distribuyendo chunks a DataNode 53")
	}

	if distributionReply.Machines[1].Status == 1 {
		fmt.Println("Distribuyendo chunks a DataNode 54")
	}

	if distributionReply.Machines[2].Status == 1 {
		fmt.Println("Distribuyendo chunks a DataNode 55")
	}

	if distributionReply.Machines[0].Status == 1 {
		for j := 0; j < len(distributionReply.Machines[0].Distribution); j++ {
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
		for j := 0; j < len(distributionReply.Machines[1].Distribution); j++ {
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
		// fmt.Println(distributionReply.Machines[2].Distribution)
		fmt.Println(len(distributionReply.Machines[2].Distribution))
		for j := 0; j < len(distributionReply.Machines[2].Distribution); j++ {
			fileName := distributionReply.FileName + "_" + strconv.Itoa(int(s.file.chunks[distributionReply.Machines[2].Distribution[j]].ChunkIndex))
			chunk := s.file.chunks[distributionReply.Machines[2].Distribution[j]].Chunk
			if distributionReply.Machines[2].Address == s.currentAddress {
				fmt.Println("Creando archivos local")
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

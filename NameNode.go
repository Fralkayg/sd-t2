package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	pb "./Service"
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
	var firstNodeDistribution []int32
	var secondNodeDistribution []int32
	var thirdNodeDistribution []int32
	var firstNodeStatus int32
	var secondNodeStatus int32
	var thirdNodeStatus int32
	var availableNodes int32
	availableNodes = 0

	fmt.Println("Llego el archivo " + in.FileName)

	for i := 53; i < 56; i++ {
		address := "dist" + strconv.Itoa(i) + port
		fmt.Println(address)
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
	}
	fmt.Println("Nodos disponibles: " + strconv.Itoa(int(availableNodes)))

	file, err := os.OpenFile("./LOG.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Println(err)
	}
	if _, err := file.WriteString(in.FileName + " " + strconv.Itoa(int(in.TotalParts)) + "\n"); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < int(in.TotalParts); i++ {
		var result int32
		if availableNodes == 3 {
			result = int32(i % 3)
			if result == 0 && firstNodeStatus == 1 {
				firstNodeDistribution = append(firstNodeDistribution, int32(i))
				writeToLogFile(file, in.FileName+"_"+strconv.Itoa(int(i))+" "+"dist53:50051\n")
			} else if result == 1 && secondNodeStatus == 1 {
				secondNodeDistribution = append(secondNodeDistribution, int32(i))
				writeToLogFile(file, in.FileName+"_"+strconv.Itoa(int(i))+" "+"dist54:50051\n")
			} else if thirdNodeStatus == 1 {
				thirdNodeDistribution = append(thirdNodeDistribution, int32(i))
				writeToLogFile(file, in.FileName+"_"+strconv.Itoa(int(i))+" "+"dist55:50051\n")
			}
		} else if availableNodes == 2 {
			result = int32(i % 2)
			if firstNodeStatus == 1 && secondNodeStatus == 1 {
				firstDistribution, secondDistribution := makeDistribution(file, in.FileName, "dist53:50051", "dist54:50051", result, i)
				firstNodeDistribution = firstDistribution
				secondNodeDistribution = secondDistribution
			} else if firstNodeStatus == 1 && thirdNodeStatus == 1 {
				firstDistribution, secondDistribution := makeDistribution(file, in.FileName, "dist53:50051", "dist55:50051", result, i)
				firstNodeDistribution = firstDistribution
				thirdNodeDistribution = secondDistribution
			} else if secondNodeStatus == 1 && thirdNodeStatus == 1 {
				firstDistribution, secondDistribution := makeDistribution(file, in.FileName, "dist54:50051", "dist55:50051", result, i)
				secondNodeDistribution = firstDistribution
				thirdNodeDistribution = secondDistribution
			}
		} else {
			if firstNodeStatus == 1 {
				firstNodeDistribution = append(firstNodeDistribution, int32(i))
				writeToLogFile(file, in.FileName+"_"+strconv.Itoa(int(i))+" "+"dist53:50051\n")

			} else if secondNodeStatus == 1 {
				secondNodeDistribution = append(firstNodeDistribution, int32(i))
				writeToLogFile(file, in.FileName+"_"+strconv.Itoa(int(i))+" "+"dist54:50051\n")
			} else {
				thirdNodeDistribution = append(firstNodeDistribution, int32(i))
				writeToLogFile(file, in.FileName+"_"+strconv.Itoa(int(i))+" "+"dist55:50051\n")
			}
		}
	}

	file.Close()

	return &pb2.DistributionReply{
		FileName:   in.FileName,
		TotalParts: int32(in.TotalParts),
		Machines: []*pb2.DistributionReply_MachineInformation{
			{Address: "dist53:50051", Distribution: firstNodeDistribution, Status: firstNodeStatus},
			{Address: "dist54:50051", Distribution: secondNodeDistribution, Status: secondNodeStatus},
			{Address: "dist55:50051", Distribution: thirdNodeDistribution, Status: thirdNodeStatus},
		},
	}, nil
}

func makeDistribution(file *os.File, fileName string, address1 string, address2 string, result int32, i int) ([]int32, []int32) {
	var firstNodeDistribution []int32
	var secondNodeDistribution []int32
	if result == 0 {
		firstNodeDistribution = append(firstNodeDistribution, int32(i))
		writeToLogFile(file, fileName+"_"+strconv.Itoa(int(i))+" "+address1+"\n")
	} else if result == 1 {
		secondNodeDistribution = append(secondNodeDistribution, int32(i))
		writeToLogFile(file, fileName+"_"+strconv.Itoa(int(i))+" "+address2+"\n")
	}
	return firstNodeDistribution, secondNodeDistribution
}

func writeToLogFile(file *os.File, line string) {
	if _, err := file.WriteString(line); err != nil {
		log.Fatal(err)
	}
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

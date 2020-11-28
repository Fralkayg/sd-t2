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

type server struct {
	seguimiento int
	lock        bool
}

const (
	port = ":50051"
)

func (s *server) SendDistributionProposal(ctx context.Context, in *pb2.DistributionRequest) (*pb2.DistributionReply, error) {
	fmt.Println("Llego el archivo " + in.FileName)
	for i := 53; i < 56; i++ {
		address := "dist" + strconv.Itoa(i) + port
		fmt.Println(address)
		status := checkNodeStatus(address)
		fmt.Println("Estado de dist" + strconv.Itoa(i) + ": ")
		fmt.Println(status)
	}

	return &pb2.DistributionReply{}, nil
}

func checkNodeStatus(address string) bool {
	fmt.Println("Entro a verificar estado")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	fmt.Println("Conexion establecida")
	c := pb.NewFileManagementServiceClient(conn)

	fmt.Println("Confirmando estado")
	_, connectionError := c.CheckNodeStatus(context.Background(), &pb.StatusRequest{Online: true})
	fmt.Println("Estado confirmado")
	if connectionError != nil {
		return false
	}
	return true
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

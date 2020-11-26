package main

import (
	"log"

	pb "./Service"
	"google.golang.org/grpc"
)

const (
	address     = "dist54:50051"
	defaultName = "world"
)

func helloWorld(conn *grpc.ClientConn) {
	c := pb.NewFileManagementServiceClient(conn)
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	helloWorld(conn)

}

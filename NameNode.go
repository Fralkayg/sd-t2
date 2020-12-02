package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

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

//Descripción: Lee el archivo LOG retornando una estructura que contiene los libros disponibles
//             junto a una estructura que contiene la ubicación de cada una de las partes que lo contiene
func (s *server) ReadLogFile(ctx context.Context, in *pb2.LogRequest) (*pb2.LogReply, error) {
	f, err := os.Open("LOG.txt")

	if err != nil {
		fmt.Println("Error al leer el archivo LOG")
		return nil, errors.New("No hay libros disponibles para descargar")
	}

	defer f.Close()

	// var logReply pb2.LogReply

	scanner := bufio.NewScanner(f)

	var files []pb2.LogReply_FileInfo

	j := int32(0)

	for scanner.Scan() {
		if scanner.Text() == "" {
			//donothing
		} else {
			var aux pb2.LogReply_FileInfo
			line := strings.Split(scanner.Text(), " ")

			totalParts, _ := strconv.Atoi(line[1])

			aux.FileName = line[0]
			aux.TotalParts = line[1]

			for i := 0; i < totalParts; i++ {
				var aux2 pb2.LogReply_FileInfo_FileDistribution
				scanner.Scan()
				filePart := strings.Split(scanner.Text(), " ")

				aux2.Part = filePart[0]
				aux2.Address = filePart[1]

				aux.Distribution = append(aux.Distribution, &aux2)
			}

			aux.FileIndex = j

			files = append(files, aux)
			j++
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo LOG")
		return nil, errors.New("No hay libros disponibles para descargar")
	}

	//ESTO FUNCIONA
	var logReply pb2.LogReply

	for i := 0; i < len(files); i++ {
		logReply.Files = append(logReply.Files, &files[i])

	}

	return &logReply, nil
}

//Descripción: Recibe la distribución aprobada por el DataNode registrando la información en el archivo LOG.txt
func (s *server) SendDistribution(ctx context.Context, in *pb2.DistributionRequest2) (*pb2.DistributionReply2, error) {
	for s.lock {

	}

	s.lock = true

	file, err := os.OpenFile("./LOG.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Println(err)
	}
	if _, err := file.WriteString(in.FileName + " " + strconv.Itoa(int(in.TotalParts)) + "\n"); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < int(in.TotalParts); i++ {
		if in.Machines[0].Status == 1 {
			for j := 0; j < len(in.Machines[0].Distribution); j++ {
				if in.Machines[0].Distribution[j] == int32(i) {
					fileName := in.FileName + "_" + strconv.Itoa(int(in.Machines[0].Distribution[j]))
					writeToLogFile(file, fileName+" "+in.Machines[0].Address+"\n")
				}
			}
		}

		if in.Machines[1].Status == 1 {
			for j := 0; j < len(in.Machines[1].Distribution); j++ {
				if in.Machines[1].Distribution[j] == int32(i) {
					fileName := in.FileName + "_" + strconv.Itoa(int(in.Machines[1].Distribution[j]))
					writeToLogFile(file, fileName+" "+in.Machines[1].Address+"\n")
				}

			}
		}

		if in.Machines[2].Status == 1 {
			for j := 0; j < len(in.Machines[2].Distribution); j++ {
				if in.Machines[2].Distribution[j] == int32(i) {
					fileName := in.FileName + "_" + strconv.Itoa(int(in.Machines[2].Distribution[j]))
					writeToLogFile(file, fileName+" "+in.Machines[2].Address+"\n")
				}

			}
		}
	}

	file.Close()

	s.lock = false

	return &pb2.DistributionReply2{
		FileName:   in.FileName,
		TotalParts: int32(in.TotalParts),
		Machines: []*pb2.DistributionReply2_MachineInformation{
			{Address: "dist53:50051", Distribution: in.Machines[0].Distribution, Status: in.Machines[0].Status},
			{Address: "dist54:50051", Distribution: in.Machines[1].Distribution, Status: in.Machines[1].Status},
			{Address: "dist55:50051", Distribution: in.Machines[2].Distribution, Status: in.Machines[2].Status},
		},
	}, nil

}

//Descripción: Recibe la propuesta de distribución inicial realizada por el DataNode.
//             Verifica que la propuesta sea válida, de ser así la reenvia al DataNode.
//			   En caso contrario, genera una nueva propuesta válida enviandola posteriormente al DataNode.
func (s *server) SendDistributionProposal(ctx context.Context, in *pb2.DistributionRequest) (*pb2.DistributionReply, error) {
	for s.lock {

	}

	s.lock = true

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
				firstDistribution, first, secondDistribution, _ := makeDistribution(file, in.FileName, "dist53:50051", "dist54:50051", result, i)
				if first {
					firstNodeDistribution = append(firstNodeDistribution, firstDistribution)
				} else {
					secondNodeDistribution = append(secondNodeDistribution, secondDistribution)
				}
			} else if firstNodeStatus == 1 && thirdNodeStatus == 1 {
				firstDistribution, first, secondDistribution, _ := makeDistribution(file, in.FileName, "dist53:50051", "dist55:50051", result, i)
				if first {
					firstNodeDistribution = append(firstNodeDistribution, firstDistribution)
				} else {
					thirdNodeDistribution = append(thirdNodeDistribution, secondDistribution)
				}
			} else if secondNodeStatus == 1 && thirdNodeStatus == 1 {
				firstDistribution, first, secondDistribution, _ := makeDistribution(file, in.FileName, "dist54:50051", "dist55:50051", result, i)
				if first {
					secondNodeDistribution = append(secondNodeDistribution, firstDistribution)
				} else {
					thirdNodeDistribution = append(thirdNodeDistribution, secondDistribution)
				}
			}
		} else {
			if firstNodeStatus == 1 {
				firstNodeDistribution = append(firstNodeDistribution, int32(i))
				writeToLogFile(file, in.FileName+"_"+strconv.Itoa(int(i))+" "+"dist53:50051\n")

			} else if secondNodeStatus == 1 {
				secondNodeDistribution = append(secondNodeDistribution, int32(i))
				writeToLogFile(file, in.FileName+"_"+strconv.Itoa(int(i))+" "+"dist54:50051\n")
			} else {
				thirdNodeDistribution = append(thirdNodeDistribution, int32(i))
				writeToLogFile(file, in.FileName+"_"+strconv.Itoa(int(i))+" "+"dist55:50051\n")
			}
		}
	}

	file.Close()

	if availableNodes == 3 {
		fmt.Println("Propuesta aceptada.")
	} else {
		fmt.Println("Se rechazo la propuesta enviada. Se reenvio una nueva propuesta válida.")
		fmt.Println("La nueva propuesta considera los siguientes nodos.")
		if firstNodeStatus == 1 {
			fmt.Println("DataNode 53 incluido en propuesta de distribucion.")
		}
		if secondNodeStatus == 1 {
			fmt.Println("DataNode 54 incluido en propuesta de distribucion.")
		}
		if thirdNodeStatus == 1 {
			fmt.Println("DataNode 55 incluido en propuesta de distribucion.")
		}
	}

	s.lock = false

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

//Descripción: Función auxiliar para registrar en LOG.txt en el caso de que se encuentren dos DataNode disponibles.
func makeDistribution(file *os.File, fileName string, address1 string, address2 string, result int32, i int) (int32, bool, int32, bool) {
	var firstNodeDistribution int32
	var secondNodeDistribution int32
	first := false
	second := false
	if result == 0 {
		firstNodeDistribution = int32(i)
		first = true
		writeToLogFile(file, fileName+"_"+strconv.Itoa(int(i))+" "+address1+"\n")
	} else if result == 1 {
		secondNodeDistribution = int32(i)
		second = true
		writeToLogFile(file, fileName+"_"+strconv.Itoa(int(i))+" "+address2+"\n")
	}
	return firstNodeDistribution, first, secondNodeDistribution, second
}

//Descripción: Escribe en el archivo que se envia como parametro la linea correspondiente que se envia como parametro.
func writeToLogFile(file *os.File, line string) {
	if _, err := file.WriteString(line); err != nil {
		log.Fatal(err)
	}
}

//Descripción: Verificar el estado de los DataNode.
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

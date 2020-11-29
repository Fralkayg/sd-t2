package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"

	pb "./Service"
	pb2 "./Service2"
	"google.golang.org/grpc"
)

const (
	nameNodeAddress = "dist56:50051"
	defaultName     = "world"
	port            = ":50051"
)

type ChunkedFile struct {
	FileName   string
	TotalParts int
	ChunkName  []string
}

func helloWorld(conn *grpc.ClientConn) {
	c := pb.NewFileManagementServiceClient(conn)

	response, _ := c.SayHello(context.Background(), &pb.HelloRequest{
		Mensaje: "Christian"})

	fmt.Println("Mensaje: ", response.Mensaje)
}

func splitFile(targetFile string) ChunkedFile {
	fileToBeChunked := "./Books/" + targetFile // change here!

	file, err := os.Open(fileToBeChunked)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	var fileSize int64 = fileInfo.Size()

	const fileChunk = 250000 // 1 MB, change this to your requirement

	// calculate total number of parts the file will be chunked into

	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	var chunkedFile ChunkedFile
	chunkedFile.TotalParts = int(totalPartsNum)
	chunkedFile.FileName = targetFile

	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

	for i := uint64(0); i < totalPartsNum; i++ {

		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)

		// write to disk
		fileName := "./Client/Chunks/" + targetFile + "_" + strconv.FormatUint(i+1, 10)
		_, err := os.Create(fileName)

		chunkedFile.ChunkName = append(chunkedFile.ChunkName, fileName)

		if err != nil {
			fmt.Println("cago aca")
			fmt.Println(err)
			os.Exit(1)
		}

		// write/save buffer to disk
		ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)
		fmt.Println("Split to : ", fileName)
	}
	return chunkedFile
}

func obtainChunk(filePath string) []byte {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

//https://flaviocopes.com/go-list-files/
func displayLibrary() []string {
	var files []string

	root := "./Books/"

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files = append(files, info.Name())
		return nil
	})

	if err != nil {
		panic(err)
	}

	var i int
	i = 1

	for _, file := range files {
		fmt.Println(i, ":", file)
		i++
	}
	return files
}

func uploadBook(option int) {
	var validOption bool
	var bookIndex int
	validOption = true

	for validOption {
		fmt.Println("Escoja libro a subir")
		files := displayLibrary()
		fmt.Scanln(&bookIndex)
		if bookIndex > len(files) || bookIndex < 0 {
			fmt.Println("Opción invalida. Reingresar")
		} else {
			chunkedFile := splitFile(files[bookIndex-1])
			address := connectToDataNode()

			conn, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()

			fmt.Println("paso")
			for i := 0; i < chunkedFile.TotalParts; i++ {
				var lastChunk bool

				if i == chunkedFile.TotalParts-1 {
					lastChunk = true
				} else {
					lastChunk = false
				}

				chunk := obtainChunk(chunkedFile.ChunkName[i])

				c := pb.NewFileManagementServiceClient(conn)

				status, _ := c.SendChunk(context.Background(), &pb.ChunkInformation{
					Chunk:      chunk,
					ChunkIndex: int32(i),
					FileName:   chunkedFile.FileName,
					LastChunk:  lastChunk,
					Option:     int32(option),
					TotalParts: int32(chunkedFile.TotalParts),
					Address:    address,
				})
				fmt.Println(status)
			}

			validOption = false
		}

	}
}

func downloadBookMenu() {
	conn, err := grpc.Dial(nameNodeAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb2.NewDataToNameServiceClient(conn)

	logReply, er := c.ReadLogFile(context.Background(), &pb2.LogRequest{File: "LOG.txt"})
	if er != nil {
		fmt.Println("Error al leer el archivo LOG")
	}

	validOption := true
	var option int
	for validOption {
		fmt.Println("Escoja un archivo a descargar:")
		for i := 0; i < len(logReply.Files); i++ {
			fmt.Println(strconv.Itoa(int(i+1)) + ":" + logReply.Files[i].FileName)
		}
		fmt.Scanln(&option)
		if option > 0 && option <= len(logReply.Files) {
			validOption = false
		}
	}
	downloadBook(logReply.Files, option-1)
}

func downloadBook(files []*pb2.LogReply_FileInfo, option int) bool {
	for i := 0; i < len(files); i++ {
		if files[i].FileIndex == int32(option) {
			newFileName := files[i].FileName
			_, err := os.Create("./Downloads/" + newFileName)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			file, openFileError := os.OpenFile("./Downloads/"+newFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

			if openFileError != nil {
				fmt.Println("Error al abrir el archivo")
				fmt.Println(err)
				os.Exit(1)
			}
			for j := 0; j < len(files[i].Distribution); j++ {
				conn, err := grpc.Dial(files[i].Distribution[j].Address, grpc.WithInsecure())
				if err != nil {
					log.Fatalf("did not connect: %v", err)
				}
				defer conn.Close()

				c := pb.NewFileManagementServiceClient(conn)
				chunk, connectionError := c.RetrieveChunk(context.Background(), &pb.ChunkRequest{
					FileName: files[i].Distribution[j].Part,
				})

				if connectionError != nil {
					fmt.Println("No se puede descargar el archivo.")
					os.Remove("./Downloads/" + newFileName)
					return false
				}
				_, writeError := file.Write(chunk.Chunk)

				if writeError != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				file.Sync()

			}
			file.Close()
		}
	}
	fmt.Println(option)
	return true
}

func centralizedOrDistributed() {
	var validOption bool
	var option int
	validOption = true

	for validOption {
		fmt.Println("Ingrese una opción")
		fmt.Println("1: Centralizado")
		fmt.Println("2: Distribuido")
		fmt.Println("3: Volver")
		fmt.Scanln(&option)

		switch option {
		case 1:
			uploadBook(option)
			validOption = false
		case 2:
			uploadBook(option)
			validOption = false
		case 3:
			validOption = false
		default:
			fmt.Println("Opción inválida. Reingrese.")
		}
	}
}

func connectToDataNode() string {
	for i := 53; i < 56; i++ {
		address := "dist" + strconv.Itoa(i) + port
		fmt.Println(address)
		_, status := checkStatus(address)

		if i == 53 && status == 1 {
			fmt.Println("DataNode 53 en línea")
			return address
		} else if i == 54 && status == 1 {
			fmt.Println("DataNode 54 en línea")
			return address
		} else if i == 55 && status == 1 {
			fmt.Println("DataNode 55 en línea")
			return address
		} else {
			return ""
		}
	}
	return ""
}

func checkStatus(address string) (*grpc.ClientConn, int32) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	c := pb.NewFileManagementServiceClient(conn)

	_, connectionError := c.CheckNodeStatus(context.Background(), &pb.StatusRequest{Online: true})
	if connectionError != nil {
		return conn, 0
	}
	return conn, 1
}

func main() {
	var validOption bool
	var option int
	// conn, err := grpc.Dial(address, grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }
	// defer conn.Close()

	validOption = true

	for validOption {
		fmt.Println("Ingrese una opción válida:")
		fmt.Println("1: Subir libro")
		fmt.Println("2: Descargar libro")
		fmt.Println("3: Salir")
		fmt.Scanln(&option)
		switch option {
		case 1:
			centralizedOrDistributed()
		case 2:
			downloadBookMenu()
		case 3:
			fmt.Println("Adiós!")
			validOption = false
		default:
			//
		}

	}

	// helloWorld(conn)

}

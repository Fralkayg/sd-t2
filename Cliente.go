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
	"google.golang.org/grpc"
)

const (
	address     = "dist54:50051"
	defaultName = "world"
)

type ChunkedFile struct {
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

//https://flaviocopes.com/go-list-files/
func displayLibrary() []string {
	var files []string

	root := "./Books/"

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	var i int
	i = 1

	for _, file := range files {
		fmt.Println(i, ":", file)
	}
	return files
}

func uploadBook(conn *grpc.ClientConn, option int) {
	var validOption bool
	var bookIndex int
	validOption = true

	for validOption {
		fmt.Println("Escoja libro a subir")
		files := displayLibrary()
		fmt.Scanln(&option)
		if bookIndex > len(files) || bookIndex <= 0 {
			fmt.Println("Opción invalida. Reingresar")
		} else {
			chunkedFile := splitFile(files[bookIndex])

			fmt.Println("Cantidad de partes: ", chunkedFile.TotalParts)

			for i := 0; i < chunkedFile.TotalParts; i++ {
				fmt.Println("Parte: ", chunkedFile.ChunkName[i])
			}
		}

	}
}

func centralizedOrDistributed(conn *grpc.ClientConn) {
	var validOption bool
	var option int
	validOption = true

	for validOption {
		fmt.Println("Ingrese una opción")
		fmt.Println("1: Centralizado")
		fmt.Println("2: Distribuido")
		fmt.Println("3: Volver")

		switch option {
		case 1:
			uploadBook(conn, option)
		case 2:
			uploadBook(conn, option)
		case 3:
			validOption = false
		default:
			fmt.Println("Opción inválida. Reingrese.")
		}
	}
}

func main() {
	var validOption bool
	var option int
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	chunkedFile := splitFile("El_maravilloso_Mago_de_Oz-L._Frank_Baum.pdf")

	validOption = true
	fmt.Println("Ingrese una opción válida:")
	for validOption {
		fmt.Println("1: Subir libro")
		fmt.Println("2: Descargar libro")
		fmt.Println("3: Salir")
		fmt.Scanln(&option)
		switch option {
		case 1:
			//Subir libro
			fmt.Println("Subir libro")
			centralizedOrDistributed(conn)
		case 2:
			//Bajar libro
			fmt.Println("Bajar libro")
		case 3:
			fmt.Println("Adiós!")
		default:
			fmt.Println("Ingrese una opción válida:")
		}

	}

	helloWorld(conn)

}

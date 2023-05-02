package main

import (
	"context"
	"fmt"
	"grpc-lesson/pb"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

func callListFiles(client pb.FilesServiceClient) {
	res, err := client.ListFiles(context.Background(), &pb.ListFilesRequest{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.GetFilenames())
}

func callDownload(client pb.FilesServiceClient) {
	req := &pb.DownloadRequest{Filename: "name.txt"}
	stream, err := client.Download(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		// always accept it every time server returns response
		res, err := stream.Recv()
		// when server reaching EOF
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("Response from Download(Bytes): %v", res.GetData())
		log.Printf("Response from Download(string): %v", string(res.GetData()))
	}

}

func callUpload(client pb.FilesServiceClient) {
	filename := "sports.txt"
	path := "/Users/user/workspace/grpc-lesson/storage/" + filename

	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	stream, err := client.Upload(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	buf := make([]byte, 5)
	for {
		n, err := file.Read(buf)
		if n == 0 || err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln(err)
		}

		// put read content into buffer
		req := &pb.UploadRequest{Data: buf[:n]}

		// send request to server
		sendErr := stream.Send(req)

		if sendErr != nil {
			log.Fatalln(sendErr)
		}
		time.Sleep(1 * time.Second)
	}

	// notify the end of request to server and get response from server
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("received data size: %v", res.GetSize())

}

func main() {
	// establish connection with localhost:50051
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	// always finish connection
	defer conn.Close()

	// get file service client
	// this method is created automatically by protocol buffer's compile
	// this func name includes service in proto file

	// func NewFilesServiceClient(cc grpc.ClientConnInterface) FilesServiceClient {
	// 	return &filesServiceClient{cc}
	// }
	// this returns interface FilesServiceClient
	// this client is FilesServiceClient type
	client := pb.NewFilesServiceClient(conn)
	// callListFiles(client)
	// callDownload(client)

	callUpload(client)

}
